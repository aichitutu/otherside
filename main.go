package main

import (
	"awesomeProject/island"
	"fmt"
)

func main() {
	i1 := island.GetIsland(0)
	i2 := island.GetIsland(686)
	i3 := island.GetIsland(3931)
	fmt.Println("中心岛点数：",len(i1))
	fmt.Println("2号岛点数：",len(i2))
	fmt.Println("3号岛点数：",len(i3))
	fmt.Println("")
	fmt.Println("")
	fmt.Println("中心岛:",i1)
	fmt.Println("2号岛：",i2)
	fmt.Println("3号岛：",i3)
}

