package router

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/itcast/zinx/iface"
	"github.com/itcast/zinx/net"
	"mmoGame/core"
	mmomsg "mmoGame/pb"
)

type Wordrouter struct {
	net.Router
}

func (w *Wordrouter)Handle(rep iface.IRequest)  {
	fmt.Println("我是世界聊天功能")
	//要获取传的数据
	protoData:=rep.GetMessage().GetMsgData()
	//进行解码
	var talkData mmomsg.Talk
	err:=proto.Unmarshal(protoData,&talkData)
	if err!=nil {
		fmt.Println("Talk proto.Unmarshal err:", err)
		return
	}
	fmt.Println("拿到数据:",talkData.Content)

	pidInterface:=rep.GetConn().GetProperty("pid")
	playerid,ok:=pidInterface.(int)
	if !ok {
		fmt.Println("pid 断言失败!")
		return
	}
	//Player:=core.WorkManange.GetPlayer(playerid)playerid
	Player:=core.GworkMange.GetPlayer(playerid)

	Player.SendContentToAllPlayers(talkData.Content)
}
