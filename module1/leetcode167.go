package module1

import (
	"slices"
	"sort"
)

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		num := target - numbers[i]
		index := slices.Index(numbers, num)
		/* if !slices.Contains(numbers, num) {
			fmt.Println("found", num)
			break
		} */
		if index != i && slices.Contains(numbers, num) {
			result := []int{index + 1, i + 1}
			sort.Ints(result)
			return result
		}
	}
	return []int{}
}
