package MMoGame

import (
	"fmt"
	"sync"
)

//格子管理器
type gridManange struct {
	//格子管理器的最大x,最小x
	GM_minX, GM_maxX float64
	cotX             int

	//格子管理器的最大y,最小y
	GM_minY, GM_maxY float64
	cotY             int

	//多个格子
	Grids map[int]*Grid

	lock sync.RWMutex
}

const (
	GM_minX = 0
	GM_maxX = 100
	cotX    = 10

	GM_minY = 0
	GM_maxY = 100
	cotY    = 10
)

//创建格子管理器
func NewGridManange() *gridManange {
	gm := &gridManange{
		GM_minX: GM_minX,
		GM_maxX: GM_maxX,
		cotX:    cotX,
		GM_minY: GM_minY,
		GM_maxY: GM_maxY,
		cotY:    cotY,
		Grids:   make(map[int]*Grid),
	}
	//把Grids铺满
	for x := 0; x < cotX; x++ {
		for y := 0; y < cotY; y++ {
			grid := y*cotX + x
			grid_minX := gm.GM_minX+float64((grid % x) * gm.GetWitdh())
			grid_maxX := gm.GM_minX+float64(((grid % x)+1) * gm.GetWitdh())
			grid_minY := gm.GM_minY+float64(y*gm.GetHeight())
			grid_maxY := gm.GM_minY+float64((y+1)*gm.GetHeight())
			//func NewGrid(grid int,minX,maxX,minY,maxY float64) *Grid {
			ManangeGrid:=NewGrid(grid,grid_minX,grid_maxX,grid_minY,grid_maxY )
			gm.Grids[grid]=ManangeGrid
		}
	}
	return gm
}

//获取长度
func (gm *gridManange) GetWitdh() int {
	return int(gm.GM_maxX-gm.GM_minX) / gm.cotX
}
//获取高度
func (gm *gridManange) GetHeight() int {
	return int(gm.GM_maxY-gm.GM_minY) / gm.cotY
}

//重写打印的
func (gm *gridManange) String() string {
	res := fmt.Sprintf("gridmanager: minX:%f, maxX:%f, cntX, %d, minY:%f, maxY:%f, cntY:%d",
		gm.GM_minX, gm.GM_maxX, gm.cotX, gm.GM_minY, gm.GM_maxY, gm.cotY)

	//对每个格子进行打印，总的输出返回
	for _, grid := range gm.Grids {
		res += fmt.Sprint(grid, "\n")
	}
	return res
}

//********把gird的方法重写一遍

func (gm *gridManange)AddPlayer(pid ,gid int,player interface{})  {
	fmt.Println("GridMananger===>AddPlayer,gid:",gid)
	gm.lock.RLock()
	defer gm.lock.RUnlock()
	_,ok:=gm.Grids[gid]
	if !ok {
		fmt.Println("无效的gid")
		return
	}
	gm.Grids[pid].AddPlayer(pid,player)
}

func (gm *gridManange)ReadPlayer(pid int,gid int)  {
	fmt.Println("GridMananger===>ReadPlayer,gid:",gid)
	gm.lock.RLock()
	defer gm.lock.RUnlock()
	_,ok:=gm.Grids[gid]
	if !ok {
		fmt.Println("无效的gid")
		return
	}
	delete(gm.Grids,pid)
}

func (gm *gridManange)GetAllPlayerid(gid int) []int {
	fmt.Println("GridMananger===>GetAllPlayIds,gid:",gid)
	_,ok:=gm.Grids[gid]
	if !ok {
		fmt.Println("无效的gid")
		return nil
	}
	return gm.Grids[gid].GetAllPlayerid()
}
//******************

//要获取这个格子的周边的格子集合
func (gm *gridManange)GetSurroundingGridsByGid(gid int) (Grids []*Grid) {
	_, ok := gm.Grids[gid]
	if !ok {
		fmt.Println("无效的gid", gid)
		return
	}
	Grids=append(Grids,gm.Grids[gid])
	if (gid % gm.cotX) > 0 {
		Grids = append(Grids, gm.Grids[gid-1])
	}

	if (gid % gm.cotX) < gm.cotX-1 {
		Grids = append(Grids, gm.Grids[gid+1])
	}
	for _,value :=range Grids{
		if (value.grid / gm.cotX) >0 {
			Grids=append(Grids,gm.Grids[value.grid-gm.cotX])
		}
		if (value.grid / gm.cotX) < gm.cotY-1 {
			Grids=append(Grids,gm.Grids[value.grid+gm.cotX])
		}
	}
	return
}

//通过具体坐标获取gid
func (gm *gridManange)GetGidBuPos(x,y float64) int {
		idx:=int(x-gm.GM_minX)/gm.GetWitdh()
		idy:=int(y-gm.GM_minY)/gm.GetHeight()

		gid:=idx+idy*gm.cotX
	return gid
}

//postion 位置

//格子管理器添加玩家
func (gm *gridManange)AddPlayerByPos(playerId int,x,y float64)  {
	gid:=gm.GetGidBuPos(x,y)
	gm.AddPlayer(playerId,gid,nil)
}

//删除玩家
func (gm *gridManange)RemovePlayerByPos(playerId int,x,y float64)  {
	gid:=gm.GetGidBuPos(x,y)
	gm.ReadPlayer(playerId,gid)
}

//获取gid的格子的周围的格子集合
func (gm *gridManange)GetSurroundingGridsByPos(x,y float64) []*Grid {
	gid:=gm.GetGidBuPos(x,y)
	return gm.GetSurroundingGridsByGid(gid)
}