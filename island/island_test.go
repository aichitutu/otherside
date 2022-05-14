package island

import (
	"fmt"
	"testing"
)

func TestIsland1(t *testing.T) {
	i1 := GetIsland(0)
	if !has(i1,[]int{329,295,88,219}) {
		t.Error("TestIsland1 failed")
	}
	i12 := GetIsland(271)
	if !diff(i1,i12) {
		t.Error("t1 failed")
	}
}

func TestIsland2(t *testing.T) {
	i2 := GetIsland(686)
	if !has(i2,[]int{320,7744,260,244,4521,7977}) {
		t.Error("TestIsland2 failed")
	}
	i22 := GetIsland(6564)
	if !diff(i2,i22) {
		t.Error("t2 failed")
	}
}

func TestIsland3(t *testing.T) {
	i3 := GetIsland(4556)
	if !has(i3,[]int{1634,6008,9158,8056,2696,2655,8806,9025,8130,9381}) {
		t.Error("TestIsland2 failed")
	}
	i32 := GetIsland(5674)
	if !diff(i3,i32) {
		t.Error("t3 failed")
	}
}

func diff(a1,a2 []int) bool {
	if len(a1)<len(a2){
		a1,a2 = a2,a1
	}
	var m = make(map[int]bool)
	for _,v := range a2{
		m[v] = true
	}
	var re []int
	for _,v := range a1{
		if !m[v]{
			re = append(re, v)
		}
	}

	if len(re) == 0{
		return true
	}
	fmt.Println("dont have",re)
	return false
}

func has(i1,a2 []int) bool {
	var m = make(map[int]bool)
	for _,v := range i1{
		m[v] = true
	}
	var tmp []int
	for _,v := range a2{
		if !m[v]{
			tmp = append(tmp,v)
		}
	}
	if len(tmp) == 0{
		return true
	}
	fmt.Println("dont have",tmp)
	return false
}