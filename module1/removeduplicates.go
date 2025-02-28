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

func Mapremoveduplicates() {
	list1 := []int{2, 4, 1, 5, 2, 3, 2, 4, 1}
	map1 := make(map[int]bool)
	var result []int
	for _, num := range list1 {
		if !map1[num] {
			map1[num] = true
			result = append(result, num)
		}
	}
	fmt.Println("", result)
}
