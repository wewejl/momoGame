package MMoGame

import (
	"github.com/itcast/zinx/iface"
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

