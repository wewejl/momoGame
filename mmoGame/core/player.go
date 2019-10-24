package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/itcast/zinx/iface"
	"math/rand"
	mmomsg "mmoGame/pb"
	"sync"
)

//属性：
type Player struct {
	//玩家Id
	playerId int

	//1.x ====>横向
	x float64

	//2.y=====>高度（没有使用）
	y float64

	//3.z=====> 纵向(传统意义的y)
	z float64

	//4.v ===>面部朝向
	v float64

	//5.Conn iface.Iconnection ===>连接
	conn iface.Iconnection
}

//playerid生成器
var pidGenerator int = 1

//锁，防止两个用户同时登陆
var pidLock sync.Mutex

func NewPlayer(x, y, z, v float64, conn iface.Iconnection) *Player {
	pidLock.Lock()
	defer pidLock.Unlock()

	player := &Player{
		playerId: pidGenerator,
		x:        100 + float64(rand.Intn(10)),
		y:        y,
		z:        100 + float64(rand.Intn(10)),
		v:        v,
		conn:     conn,
	}
	pidGenerator++
	return player
}

//同步pid

//func (p *Player)SyncPid()  {
//	fmt.Println("同步pid给客户段：",p.playerId)
//	//创建一个SyncPid协议结构
//	protoData :=mmomsg.Syncid{
//		Pid:                  int32(p.playerId),
//	}
//	//2.使用Proto进行编码
//	binaryData,err:=proto.Marshal(&protoData)
//	if err != nil {
//		fmt.Println("SyncPid proto.Marshal err:",err)
//		return
//	}
//
//	//3.通过Connection传递消息
//	p.conn.Send(binaryData,1)
//}

//获取playerid
func (p *Player) GetPlayerId() int {
	return p.playerId
}

//同步pid
func (p *Player) SyncPid() {
	fmt.Println("同步pid给客户端：", p.playerId)
	//创建一个SyncPid协议结构
	ProtoData := mmomsg.Syncid{
		Pid: int32(p.playerId),
	}
	//2.是用proto进行编码
	//	bytesbufferdata,err:=proto.Marshal(&ProtoData)
	//	if err != nil {
	//		fmt.Println("SyncPid proto.Marshal err",err)
	//		return
	//	}
	//	//3.通过Conection传递消息
	//	p.conn.Send(bytesbufferdata,1)
	p.Sendmsg(1, &ProtoData)
}

//玩家上线后,需要同步一下位置
func (p *Player) SyncPosition() {
	fmt.Println("SyncPosition:", p.x, p.y, p.z, p.v)

	//1.创建一个SyncPid协议结构
	protoData := mmomsg.BroadCast{
		Pid: int32(p.playerId),
		Tp:  2,
		Data: &mmomsg.BroadCast_P{P: &mmomsg.Position{
			X: float32(p.x),
			Y: float32(p.y),
			Z: float32(p.z),
			V: float32(p.v),
		}},
	}

	p.Sendmsg(200, &protoData)
}

//我想把send在封装下
func (p *Player) Sendmsg(msgId uint32, pd proto.Message) {
	//进行发送 先进行包装
	bytusprotoData, err := proto.Marshal(pd)
	if err != nil {
		fmt.Println("proto.Marshal err:", err)
		return
	}
	p.conn.Send(bytusprotoData, msgId)
}

func (p *Player) SendContentToAllPlayers(content string) {
	//先组织文件配置
	//组织发送数据
	contentData := mmomsg.BroadCast{
		Pid:  int32(p.playerId),
		Tp:   1,
		Data: &mmomsg.BroadCast_Content{Content: content},
	}

	//获取全部玩家
	players := GworkMange.GetAllPlayer()
	for _, player := range players {
		player.Sendmsg(200, &contentData)
	}
}

//让其他的人可以看到自己，和我看到别人
func (p *Player) SyncSurround() {
	//把我上线的位置打包
	protodata := &mmomsg.BroadCast{
		Pid: int32(p.playerId),
		Tp:  2,
		Data: &mmomsg.BroadCast_P{P: &mmomsg.Position{
			X: float32(p.x),
			Y: float32(p.y),
			Z: float32(p.z),
			V: float32(p.v),
		}},
	}
	//把玩家给别人看
	Grids := GworkMange.gridMar.GetSurroundingGridsByPos(p.x, p.z)
	//把附近的玩家地址集合

	var protoPlayerCon [] *mmomsg.SyncPlayers_Player
	//fmt.Println("player :Grids",Grids)
	for _, grid := range Grids {
		players := GworkMange.gridMar.GetAllPlayIds(grid.gid)
		for _, playerid := range players {
			fmt.Println("********Player", playerid)
			Player := GworkMange.GetPlayer(playerid)
			Player.Sendmsg(200, protodata)

			Positioncon := &mmomsg.SyncPlayers_Player{
				Pid: int32(Player.playerId),
				P: &mmomsg.Position{
					X: float32(Player.x),
					Y: float32(Player.y),
					Z: float32(Player.z),
					V: float32(Player.v),
				},
			}
			protoPlayerCon = append(protoPlayerCon, Positioncon)
		}
	}

	protoPlayers := &mmomsg.SyncPlayers{
		Ps: protoPlayerCon,
	}
	fmt.Println("*****把自己的周围的旁边的人都传过来", protoPlayers)
	//把自己的周围的旁边的人都传过来
	p.Sendmsg(202, protoPlayers)
}

//广播玩家的位置   位置更新
func (p *Player) UpdatePosition(x, y, z, v float32) {
	//更新玩家的位置
	p.x = float64(x)
	p.y = float64(y)
	p.z = float64(z)
	p.v = float64(v)
	//把消息进行封装
	UpdataPosition := &mmomsg.BroadCast{
		Pid: int32(p.playerId),
		Tp:  4,
		Data: &mmomsg.BroadCast_P{
			P: &mmomsg.Position{
				X: x,
				Y: y,
				Z: z,
				V: v,
			}},
	}

	players := p.GerSurroundPlayes()

	for _, player := range players {
		player.Sendmsg(200, UpdataPosition)
	}

}

func (p *Player) GerSurroundPlayes() (Players []*Player) {
	Grids := GworkMange.gridMar.GetSurroundingGridsByPos(p.x, p.z)
	fmt.Println("Grids ************",Grids)
	for _, grid := range Grids {
		fmt.Println("GerSurroundPlayes grid.gid:",grid.gid)
		players := GworkMange.gridMar.GetAllPlayIds(grid.gid)
		for _, playerid := range players {
			Player := GworkMange.GetPlayer(playerid)
			Players = append(Players, Player)
		}
	}
	return
}

//玩家下线逻辑
func (p *Player) OffLine() {
	//玩家下线通知
	//先通知全部玩家
	protocData:=&mmomsg.Syncid{
		Pid: int32(p.playerId),
	}
	//1.拿到集合玩家
	Players := p.GerSurroundPlayes()
	for _, player := range Players {
		player.Sendmsg(202,protocData)
	}
	//在将玩家在世界管理器的player删除
	GworkMange.RemovePlayer(p)
	//在将玩家在格子管理器的palyer删除
	GworkMange.gridMar.RemovePlayerByPos(p.playerId,p.x,p.z)
}
