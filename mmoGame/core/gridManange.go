package core

import (
	"fmt"
)

type GridMananger struct {
	//1. 画布横向边界值:
	minX, maxX float64
	//2. 横向的格子数量:
	cntX int
	//3. 画布纵向边界值:
	minY, maxY float64
	//4. 纵向的格子数量:
	cntY int
	//5. 管理所有格子的集合: //a.Key:格子id //b.value:格子本身
	grids map[int]*Grid
}

func NewGridManager(minX, maxX float64, cntX int, minY, maxY float64, cntY int) *GridMananger {
	//1. 将边界值赋值给对应的字段
	gm := GridMananger{
		minX:  minX,
		maxX:  maxX,
		cntX:  cntX,
		minY:  minY,
		maxY:  maxY,
		cntY:  cntY,
		grids: make(map[int]*Grid),
	}
	//2. 对每一个格子进行初始化
	for y := 0; y < gm.cntY; y++ {
		for x := 0; x < gm.cntX; x++ {
			grid := y*gm.cntX + x
			minX := gm.minX + float64(grid%gm.cntX)*gm.GetWidth()
			maxX := gm.minX + float64(grid%gm.cntX+1)*gm.GetWidth()
			minY := gm.minY + float64(y)*gm.GetHegiht()
			maxY := gm.minY + float64(y+1)*gm.GetHegiht()

			gm.grids[grid] = &Grid{
				gid:     grid,
				minX:    minX,
				maxX:    maxX,
				minY:    minY,
				maxY:    maxY,
				Players: make(map[int]*Player),
			}
		}
	}
	//TODO
	return &gm
}

func (gm *GridMananger) GetWidth() float64 {
	return (gm.maxX - gm.minX) / float64(gm.cntX)
}

func (gm *GridMananger) GetHegiht() float64 {
	return (gm.maxY - gm.minY) / float64(gm.cntY)
}

func (gm *GridMananger) String() string {
	res := fmt.Sprintf("gridmanager: minX:%f, maxX:%f, cntX, %d, minY:%f, maxY:%f, cntY:%d",
		gm.minX, gm.maxX, gm.cntX, gm.minY, gm.maxY, gm.cntY)

	//对每个格子进行打印，总的输出返回
	for _, grid := range gm.grids {
		res += fmt.Sprint(grid, "\n")
	}

	return res
}

//操作具体格子的方法
func (gm *GridMananger) AddPlayer(pid, gid int) {
	fmt.Println("GridManager =>AddPlayer, pid :", pid, ", gid:", gid)
	grid, ok := gm.grids[gid]
	if !ok {
		fmt.Println("gid无效,gid:", gid)
		return
	}

	//真正将pid添加到格子中
	grid.AddPlayerToGrid(pid, nil)
}

//a.会调用Grid的AddPlayerToGrid方法
func (gm *GridMananger) RemovePlayer(pid, gid int) {
	fmt.Println("GridManager ====>RemovePlayer ,pid:", pid, ",gid:", gid)
	grid, ok := gm.grids[gid]
	if !ok {
		fmt.Println("gid无效，gid", gid)
		return
	}
	//将玩家从格子里面退出
	grid.RemovePlayerFromGrid(pid)
}

//a.会调用Grid的RemovePlayerFromGrid方法
func (gm *GridMananger) GetAllPlayIds(gid int) []int {
	fmt.Println("GridMananger===>GetAllPlayIds,gid:",gid)
	grid,ok:=gm.grids[gid]
	if !ok {
		fmt.Println("gid无效，gid", gid)
		return nil
	}
	return grid.GetAllPlayerIds()
}




func (gm *GridMananger) GetSurroundingGridsByGid(gid int) (grids []*Grid) {
	_, ok := gm.grids[gid]
	if !ok {
		fmt.Println("无效的gid", gid)
		return
	}
	grids = append(grids, gm.grids[gid])
	if (gid % gm.cntX) > 0 {
		grids = append(grids, gm.grids[gid-1])
	}

	if (gid % gm.cntX) < gm.cntX-1 {
		grids = append(grids, gm.grids[gid+1])
	}

	for _, value := range grids {
		if (value.gid / gm.cntX) > 0 {
			grids = append(grids, gm.grids[value.gid-gm.cntX])
		}
		if (value.gid / gm.cntX) < gm.cntY-1 {
			grids = append(grids, gm.grids[value.gid+gm.cntX])
		}
	}
	return
}

//通过具体坐标获取gid
//x :18.5, y:20.3
//通过x，y获取到格子gid  ==》 GetGidByPos(x, y)  ==> gid
func (gm *GridMananger)GetGidBuPos(x ,y float64) int {
		idx :=int((x - gm.minX) /gm.GetWidth())
		idy :=int((y - gm.minY) /gm.GetHegiht())
		//现在我们拿到用户的x,y 根据公式得到gid
		gid:=idx+idy*gm.cntX
		return int(gid)
}

//postion 位置
//添加玩家
func (gm *GridMananger)AddPlayerByPos(playerId int,x,y float64)  {
	gid:=gm.GetGidBuPos(x,y)
	gm.AddPlayer(playerId,gid)
}

//删除玩家
func (gm *GridMananger)RemovePlayerByPos(playerId int,x,y float64)  {
	gid:=gm.GetGidBuPos(x,y)
	gm.RemovePlayer(playerId,gid)
}

//获取gid的格子的周围的格子集合
func (gm *GridMananger)GetSurroundingGridsByPos(x,y float64) []*Grid  {
	//fmt.Println("**********x",x,"***y",y)
	gid := gm.GetGidBuPos(x,y)
	//fmt.Println(gid)
	return gm.GetSurroundingGridsByGid(gid)

}

