package island

import (
	"encoding/json"
	"os"
)

const maxVariance = 10e-7

var Graph = make(map[int][]int)

func init() {
	bs,err := os.ReadFile("./data.json")
	if err != nil{
		bs,err = os.ReadFile("../data.json")
		if err != nil{
			panic(err)
		}
	}
	var data [][2]float64
	err = json.Unmarshal(bs,&data)
	if err != nil{
		panic(err)
	}
	for i,arr := range data{
		for j := i+1; j < len(data); j++{
			if distance(arr[0],arr[1],data[j][0],data[j][1])<maxVariance{
				Graph[i] = append(Graph[i],j)
				Graph[j] = append(Graph[j],i)
			}
		}
	}
	go initWebData(data[:20000])
}

func distance(x1,y1,x2,y2 float64) float64 {
	return (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2)
}

func bfs(start int, nodes map[int][]int, fn func(int)) {
	frontier := []int{start}
	visited := map[int]bool{
		start: true,
	}
	var next []int

	for 0 < len(frontier) {
		next = []int{}
		for _, node := range frontier {
			fn(node)
			for _,n := range nodes[node]{
				if !visited[n]{
					visited[n] = true
					next = append(next, n)
				}
			}
		}
		frontier = next
	}
}

func GetIsland(num int) (v []int){
	bfs(num, Graph, func(i int) {
		v = append(v, i)
	})
	return v
}