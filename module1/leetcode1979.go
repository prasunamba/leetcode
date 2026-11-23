package module1

import "sort"

func FindGCD(nums []int) int {
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	for {
		temp := max % min
		if temp != 0 {
			max = min
			min = temp
		} else {
			return min
		}
	}
}

// ------------------------------------------------------------
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func FindGCD2(nums []int) int {
	sort.Ints(nums)

	return gcd(nums[0], nums[len(nums)-1])
}
