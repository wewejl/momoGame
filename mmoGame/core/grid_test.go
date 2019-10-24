package core

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	grid := NewGrid(10, 1, 10, 1, 10)
	players1 := "猪心也"
	players2 := "传奇任务"
	grid.AddPlayerToGrid(1, players1)
	grid.AddPlayerToGrid(2, players2)
	AllPlayers := grid.GetAllPlayerIds()

	if len(AllPlayers) != 2 {
		t.Error("没有到达预期的长度2 ",len(AllPlayers))
	}

	grid.RemovePlayerFromGrid(1)
	AllPlayers = grid.GetAllPlayerIds()
	if len(AllPlayers) !=1 {
		t.Error("没有达到预期的长度1",len(AllPlayers))
	}

}
