package router

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/itcast/zinx/iface"
	"github.com/itcast/zinx/net"
	"mmoGame/core"
	mmomsg "mmoGame/pb"
)

type MaveRouter struct {
	net.Router
}

//移动更新
func (m *MaveRouter)Handle(rep iface.IRequest)  {
	//获取数据
	protoData:=rep.GetMessage().GetMsgData()
	//创建对象来接受
	var position mmomsg.Position
	//解析
	err:=proto.Unmarshal(protoData,&position)
	if err != nil {
		fmt.Println("MaveRouter Handle  proto.Unmarshal err:",err)
		return
	}
	//得到解析后的数据了  要进行广播   得到玩家
	pidInterface:=rep.GetConn().GetProperty("pid")
	playerid,ok:=pidInterface.(int)
	if !ok {
		fmt.Println("pid 断言失败!")
		return
	}
	Player:=core.GworkMange.GetPlayer(playerid)
	fmt.Println("*****************玩家更新位置了")
	Player.UpdatePosition(position.X,position.Y,position.Z,position.V)
}