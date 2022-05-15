package island

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var webData [][][2]float64

func InitRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLFiles("island/island.html")
	r.GET("/island", func(c *gin.Context) {
		c.HTML(http.StatusOK, "island.html", gin.H{
			"data":          webData,
		})
	})
	return r
}

func initWebData(data [][2]float64)[][][2]float64{
	webData = make([][][2]float64,4)
	i1 := GetIsland(0)
	i2 := GetIsland(686)
	i3 := GetIsland(3931)
	fmt.Println("Center Island Points: ",len(i1))
	fmt.Println("Island 2 Points",len(i2))
	fmt.Println("Island 3 Points",len(i3))
	visited := make(map[int]bool)
	for i,arr := range [][]int{i1,i2,i3}{
		var tmp [][2]float64
		for _,v := range arr{
			visited[v] = true
			tmp = append(tmp, data[v])
		}
		webData[i] = tmp
	}
	for k,v := range data{
		if !visited[k]{
			webData[3] = append(webData[3],v)
		}
	}
	return webData
}