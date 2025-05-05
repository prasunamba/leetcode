package module1

import "fmt"

func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]%2 == 0 {
			left++
		} else if nums[right]%2 != 0 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	fmt.Println("result", nums)
	return nums
}
