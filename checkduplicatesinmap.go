package main

import "fmt"

// check for duplicates from map1 and map2 and add to map3
func Map3() {
	var map1 map[int]string
	map1 = make(map[int]string)
	map3 := make(map[int]string)
	map1[1] = "ok"
	map1[2] = "op"
	fmt.Println("map1", map1)
	map2 := map[int]string{
		1: "ok",
		3: "op",
	}
	for k, v := range map1 {
		for i, j := range map2 {
			if v == j && k == i {
				map3[k] = map1[k]
			}
		}
	}
	fmt.Println("map3", map3)
}
