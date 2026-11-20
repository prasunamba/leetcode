package module1

import (
	"fmt"
	"sort"
)

func FindKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	return nums[len(nums)-k]
}
