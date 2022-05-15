package main

import "land/island"

func main() {
	island.InitRouter().Run()
	//island.GetIsland(686)
	//fmt.Println(time.Now().UnixNano() - island.T,"s")
}
