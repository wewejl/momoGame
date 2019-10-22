package MMoGame

import (
	"fmt"
	"sync"
)

type WordManange struct {
	//map要上锁
	lock sync.RWMutex
	//玩家的集合
	Players map[int]*Player
	//格子管理器
	gridMar *gridManange
}

func NewWorkManange() *WordManange  {
	return &WordManange{
		Players:make(map[int]*Player),
		gridMar:NewGridManange(),
	}
}


//方法：
//
//1. 玩家上线时==》添加玩家
//1. 将玩家添加到集合中
//2. 将玩家添加到画布中

func (w *WordManange)AddPlayer(player *Player)  {
	w.lock.Lock()
	//1.将玩家从集合中添加
	w.Players[player.playerId]=player
	//2.将玩家添加到画布中
	gid:=w.gridMar.GetGidBuPos(player.x,player.z)
	w.gridMar.AddPlayer(player.playerId,gid,nil)
	//w.gridMar.AddPlayerByPos(player.playerId,player.x,player.z)
	w.lock.Unlock()
}

//2. 玩家下线时==》删除玩家
//1. 将玩家从集合中删除
//2. 从画布中将玩家删除

func (w *WordManange)RemovePlayer(player *Player)  {
	//将玩家从集合中删除
	delete(w.Players,player.playerId)
	//从画布中将玩家删除
	gid:=w.gridMar.GetGidBuPos(player.x,player.z)

	w.gridMar.ReadPlayer(player.playerId,gid)
}


//3.根据pid获取到对应的玩家
func (w *WordManange)GetPlayer(pid int) *Player {
	//读锁
	w.lock.RLock()
	value,ok:=w.Players[pid]
	if !ok {
		fmt.Println("[WorkManange]:WorkManange GetPlayer pid无效")
		return nil
	}

	w.lock.RUnlock()
	return value
}

//4.获取全部玩家GetAllPlayer
func (w *WordManange)GetAllPlayer() (Players []*Player)  {
	w.lock.RLock()
	for _,value:=range w.Players{
		Players=append(Players,value)
	}
	w.lock.RUnlock()
	return
}