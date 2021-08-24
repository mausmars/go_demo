package main

import (
	"fmt"
)

type ISceneEvent interface {
	EventType() int
}

type IScene interface {
	Size() int
}

type Scene struct {
	size  int
	level int

	scenes    [][]int
	chanEvent chan ISceneEvent
}

func NewScene(size int, level int) {

}

func main() {
	length := 1000 //地图边长
	totalCell := length * length
	userTotalCell := 10 * 10              //玩家占用格子数
	userSize := totalCell / userTotalCell //玩家数量

	fmt.Println("地图边长 length=", length)
	fmt.Println("总格子数 totalCell=", totalCell)
	fmt.Println("玩家占用格子数 userTotalCell=", userTotalCell)
	fmt.Println("玩家数量 userSize=", userSize)
}
