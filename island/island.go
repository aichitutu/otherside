package island

import (
	"encoding/json"
	"os"
)

var Graph = make(map[int][]int)

var maxVariance = 10e-7

func init() {
	bs,err := os.ReadFile("./data.json")
	if err != nil{
		bs,err = os.ReadFile("../data.json")
		if err != nil{
			panic(err)
		}
	}
	var is struct {
		Data [][3]float64
	}
	err = json.Unmarshal(bs,&is)
	if err != nil{
		panic(err)
	}
	data := is.Data
	for i,arr := range data{
		for j := i+1; j < len(data); j++{
			if distance(arr[1],arr[2],data[j][1],data[j][2])<maxVariance{
				Graph[i] = append(Graph[i],j)
				Graph[j] = append(Graph[j],i)
			}
		}
	}
}

func distance(x1,y1,x2,y2 float64) float64 {
	return (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2)
}

func bfs(start int, nodes map[int][]int, fn func(int)) {
	frontier := []int{start}
	visited := map[int]bool{}
	var next []int

	for 0 < len(frontier) {
		next = []int{}
		for _, node := range frontier {
			if visited[node] {
				continue
			}
			visited[node] = true
			fn(node)
			for _, n := range bfsFrontier(node, nodes, visited) {
				next = append(next, n)
			}
		}
		frontier = next
	}
}

func bfsFrontier(node int, nodes map[int][]int, visited map[int]bool) []int {
	var next []int
	iter := func(n int) bool { _, ok := visited[n]; return !ok }
	for _, n := range nodes[node] {
		if iter(n) {
			next = append(next, n)
		}
	}
	return next
}

func GetIsland(num int) (v []int){
	bfs(num, Graph, func(i int) {
		v = append(v, i)
	})
	return v
}