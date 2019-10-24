package core

import (
	"fmt"
	"testing"
)

func TestNewGridManager(t *testing.T) {
	gm:=NewGridManager(0,50,5,0,10,10)
	//fmt.Println("TestNewGridManager gm :",gm)
	for _,value:=range gm.GetSurroundingGridsByGid(4){
		fmt.Print(value.gid," ")
	}

}