package MMoGame

import "sync"

//定义一个格子
type Grid struct {
	//格子id
	grid int
	//最小x坐标和最大x坐标
	minX,maxX float64
	//最小Y坐标和最大Y坐标
	minY,maxY float64
	//格子的玩家
	Players map[int]interface{}
	//上锁
	Lock  sync.RWMutex
}

//创建新的格子
func NewGrid(grid int,minX,maxX,minY,maxY float64) *Grid {
	return &Grid{
		grid:grid,
		minX:minX,
		maxX:maxX,
		minY:minY,
		maxY:maxY,
		Players:make(map[int]interface{}),
	}
}

//要有一个添加玩家
func (g *Grid)AddPlayer(pid int,player interface{})  {

	//进行map读写操作是有线程安全的   要加锁
	g.Lock.Lock()
	g.Players[pid]=player		//现在把是没有player的类型   先让player类型是interface
	g.Lock.Unlock()
}

//要有一个删除玩家
func (g *Grid)ReadPlayer(pid int)  {
	//删除也是写 要进行加锁
	g.Lock.Lock()
	delete(g.Players,pid)
	g.Lock.Unlock()
}
//要一个获取全部玩家
func (g *Grid)GetAllPlayerid() (Playerid []int)  {
	//这里面是读锁   要加读锁
	g.Lock.RLock()
	for key,_:=range g.Players{
		Playerid = append(Playerid, key)
	}
	g.Lock.RUnlock()
	return
}










