package core

import (
	"fmt"
	"sync"
)

//定义一个格子
type Grid struct {
	//格子id  每个格子的唯一标识
	gid int

	//最小和最大横坐标
	minX,maxX float64
	//最小和最大纵坐标
	minY,maxY float64

	//管理当前格子里面的玩家的集合
	//key：玩家id,value：玩家本身			//现在是interface  因为没有定义
	Players map[int]*Player

	//map要上锁
	lock sync.RWMutex
}

func NewGrid(gid int,minX,maxX,minY,maxY float64) *Grid {
	return &Grid{
		gid:gid,
		minX:minX,
		maxX:maxX,
		minY:minY,
		maxY:maxY,
		Players:make(map[int]*Player),
	}
}

//添加玩家到格子 AddPlayerToGrid
func (g *Grid)AddPlayerToGrid(playerid int,player *Player )  {
	//上锁
	fmt.Println("g.Players[palyerid]",g.Players[playerid])
	g.lock.Lock()
	g.Players[playerid]=player
	g.lock.Unlock()
}
//从格子中删除玩家 RemovePlayerFromGrid
func (g *Grid)RemovePlayerFromGrid(playerid int)  {
	//上锁
	g.lock.Lock()
	delete(g.Players,playerid)
	g.lock.Unlock()

}

//获取当前格子内所有的玩家的id集合 GetAllPlayerIds
func (g *Grid)GetAllPlayerIds() []int {
	Players:=[]int{}
	g.lock.RLock()
	for key,_:=range g.Players{
		Players = append(Players, key)
	}
	g.lock.RUnlock()
	return Players
}

//func NewGrid(gid int,minX,maxX,minY,maxY int) *Grid {
//重写 String
func (g *Grid)String()string  {
	return fmt.Sprintf("git:%d,minX:%f,maxX:%f,minY:%f,maxY:%f",g.gid,g.minX,g.maxX,g.minY,g.maxY)
}


