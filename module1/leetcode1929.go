package module1

import "fmt"

func GetConcatenation(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		nums = append(nums, nums[i])
	}
	fmt.Println("", nums)
	return nums
}
