package module1

import "fmt"

func RemoveDuplicates() {
	list := []int{1, 2, 3, 2, 1, 3, 4, 5, 77, 33}
	var result []int
	map1 := make(map[int]int)
	for _, num := range list {
		map1[num]++

	}
	for key, value := range map1 {
		if value == 1 {
			result = append(result, key)
		}
	}
	fmt.Println("map", map1)
	fmt.Println("result", result)
}
