package module1

import (
	"fmt"
	"sort"
)

func OptimizeddivideArray(nums []int) bool {
	sort.Ints(nums) // Sort the array so equal numbers are adjacent

	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] { // If a pair is not equal, return false
			return false
		}
	}
	return true // All pairs are valid
}
func divideArray(nums []int) bool {
	sort.Ints(nums)
	fmt.Println("", nums)
	pairs := len(nums) / 2
	var resultslice [][]int
	count := 0
	for i := 0; i < len(nums); i = i + 2 {
		s := nums[i : i+2]
		fmt.Println("", nums[i:i+2])

		for j := 0; j <= len(s)-2; j++ {
			if s[j] == s[j+1] {
				count++
				fmt.Println("count", count)
				resultslice = append(resultslice, s)
			} else {
				return false
			}
			fmt.Println("re,s", s, resultslice)
		}
	}
	return count == pairs
}
