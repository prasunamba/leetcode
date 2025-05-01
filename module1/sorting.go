package module1

import (
	"fmt"
	"sort"
)

func Sorting() {
	nums := []int{2, 1, 4, 3, 11, 45}
	// sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	/* sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	}) */
	sort.Ints(nums)
	fmt.Println("", nums)
}
