package module1

import "fmt"

func ApplicationPairs(deviceCapacity int32, foregroundAppList [][]int32, backgroundAppList [][]int32) [][]int32 {
	// Write your code here

	foremap := make(map[int32]int32)

	for _, pair := range foregroundAppList {
		if len(pair) == 2 { // Ensure valid key-value pair
			foremap[pair[0]] = pair[1]
		}
	}
	backmap := make(map[int32]int32)

	for _, pair := range backgroundAppList {
		if len(pair) == 2 { // Ensure valid key-value pair
			backmap[pair[0]] = pair[1]
		}
	}
	op := [][]int32{}
	pair := []int32{}
	for id, mem := range foremap {
		for id2, mem2 := range backmap {
			fmt.Println(id, id2, mem, mem2)
			if deviceCapacity > mem+mem2 {
				pair = append(pair, id, id2)

				op = append(op, pair)
			}
		}
	}
	return op

}
