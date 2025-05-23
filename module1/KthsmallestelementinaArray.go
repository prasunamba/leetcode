package module1

import (
	"sort"
)

func KthsmallestelementinaArray(nums []int, k int) int {
	sort.Ints(nums)
	return nums[k-1]
}
