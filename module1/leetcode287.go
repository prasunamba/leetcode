package module1

import "fmt"

func findDuplicate(nums []int) int {
	numbers := len(nums) - 1
	sum := (numbers * (numbers + 1)) / 2
	total := 0
	for _, num := range nums {
		total += num
	}
	return total - sum
}
func findDuplicate_v2(nums []int) int {
	map1 := make(map[int]int)
	for _, val := range nums {
		map1[val]++
	}
	fmt.Println("map", map1)
	for idx, val := range map1 {
		fmt.Println("idx,val", idx, val)
		if val > 1 {
			return idx
		}
	}
	return 0
}
