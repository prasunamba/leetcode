package module1

import (
	"fmt"
	"sort"
)

func NumberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i <= len(nums)-2; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	fmt.Println("", nums)
	return nums
}
