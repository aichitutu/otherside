package island

import (
	"encoding/json"
	"os"
)

var Graph = make(map[int][]int)

var maxVariance = 10e-7

func parseData() {
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
	//data := is.Data	too slow
	data := is.Data[:20000]
	for index,arr := range data{
		for i,a := range data {
			if index == i {
				continue
			}
			if distance(arr[1],arr[2],a[1],a[2])<maxVariance{
				key := int(arr[0])
				Graph[key] = append(Graph[key],i)
			}
		}
	}
}

func distance(x1,y1,x2,y2 float64) float64 {
	return (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2)
}

func distance2(d1,d2 [3]float64) float64{
	return (d1[1]-d2[1])*(d1[1]-d2[1])+(d1[2]-d2[2])*(d1[2]-d2[2])
}

func bfs(start int, nodes map[int][]int, fn func(int)) {
	frontier := []int{start}
	visited := map[int]bool{}
	next := []int{}

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
	next := []int{}
	iter := func(n int) bool { _, ok := visited[n]; return !ok }
	for _, n := range nodes[node] {
		if iter(n) {
			next = append(next, n)
		}
	}
	return next
}

func GetIsland(num int) (v []int){
	Graph = make(map[int][]int)
	parseData()
	bfs(num, Graph, func(i int) {
		v = append(v, i)
	})
	return v
}

func GetIsland2(num int,maxD float64) (v []int){
	Graph = make(map[int][]int)
	maxVariance = maxD
	parseData()
	bfs(num, Graph, func(i int) {
		v = append(v, i)
	})
	return v
}