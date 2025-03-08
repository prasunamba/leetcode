package module1

import "fmt"

func FindMissingANDRepeatedValues() {
	matrix := [][]int{
		{1, 3},
		{2, 2},
	}
	l := len(matrix) * len(matrix)
	map1 := make(map[int]int)

	for _, mat := range matrix {
		for _, key := range mat {
			map1[key]++
		}
	}
	value, double := -1, -1
	for i := 1; i <= l; i++ {
		fmt.Println("i", i)
		if _, exists := map1[i]; !exists {
			value = i
		}
		for k := range map1 {
			if map1[k] == 2 {
				double = k
			}
		}
	}
	// result := []int{double, value}
	fmt.Println("", double, value)
}
