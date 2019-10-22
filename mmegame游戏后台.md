# 创建基础grid.go 格子

代码思维：我们要写一个游戏就要创建一个画布，我们先创建一个大小固定的格子，然后我们在创建格子管理器，格子管理器是要把格子进行管理拼接成一个画布、格子管理器是实现多的一个功能就是（我这个玩家，我只要看到我所在格子周围所连接的格子是否有玩家）。

### 定义一个格子

```go
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
```

然后格子的功能

```go

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


```

## 2.创建格子管理器就是(画布)

```go
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
//定义画布的  这个可以在配置文件里面获取
const (
	GM_minX = 0
	GM_maxX = 100
	cotX    = 10

	GM_minY = 0
	GM_maxY = 100
	cotY    = 10
)

```

创建格子管理器

```go

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
//基础方法
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
```

把gird的方法重写一遍

```go
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
```

要获取这个格子的周边的格子集合

```go
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

```

新添加方法

```go
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
```



## 添加Player玩家的集合Player.go

```go

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
	conn 	iface.Iconnection
}

//playerid生成器
var pidGenerator int = 1

//锁，防止两个用户同时登陆
var pidLock sync.Mutex
func NewPlayer(x,y,z,v float64,conn iface.Iconnection) *Player {
	pidLock.Lock()
	defer pidLock.Unlock()

	player:=&Player{
		playerId: pidGenerator,
		x:        x,
		y:        y,
		z:        z,
		v:        v,
		conn:     conn,
	}
	return player
}


```

## 添加WorkManager.go

```go


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
```

