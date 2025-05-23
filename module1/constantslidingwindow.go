package module1

import (
	"fmt"
	"golang/utils"
)

// k- window size
// max sum from a constant windowsize
func Sumconstant(nums []int, k int) int {
	sum := 0
	maxsum := 0
	l := 0
	r := k - 1
	for i := l; i <= r; i++ {
		sum += nums[i]
	}
	fmt.Println("sum", sum)
	for r < len(nums)-1 {
		sum -= nums[l]
		l++
		r++
		sum += nums[r]
		fmt.Println("nums[r],sum,maxsum", nums[r], sum, maxsum)
		maxsum = utils.Maxm(sum, maxsum)
	}
	return maxsum
}
