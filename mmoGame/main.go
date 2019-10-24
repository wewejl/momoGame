package main

import (
	"fmt"
	"github.com/itcast/zinx/iface"
	"github.com/itcast/zinx/net"
	"mmoGame/core"
	"mmoGame/router"
)

//钩子函数
func TooManange(conn iface.Iconnection) {
	//在世界的软件
	//workMange := core.NewWorkManange()
	//创建一个玩家
	player := core.NewPlayer(100, 0, 100, 0, conn)
	//fmt.Println(workMange,player)

	//同步pid
	player.SyncPid()
	//要客户端进行同步位置
	player.SyncPosition()

	conn.SetProperty("pid",player.GetPlayerId())
	core.GworkMange.AddPlayer(player)
	player.SyncSurround()
}
func onEndFunc(conn iface.Iconnection)  {
	fmt.Println("onEndFunc called!")
	pidInterface:=conn.GetProperty("pid")
	pid,ok:=pidInterface.(int)
	if !ok {
		fmt.Println("pid断言失败！")
		return
	}
	//根据pid获取玩家
	player := core.GworkMange.GetPlayer(pid)
	player.OffLine()
}

func main() {

	server := net.NewServer("mmoGame")
	//注册世界聊天路由
	server.AddRouter(2,&router.Wordrouter{}) //注册聊天处理路由
	server.AddRouter(3,&router.MaveRouter{})
	//注册钩子函数
	server.AddStartHookFunc(TooManange)
	server.AddStopHookFunc(onEndFunc)
	server.Serve()
}
