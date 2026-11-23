package module1

import (
	"main/utils"
	"math"
	"sort"
)

// so tricky
func LargestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	minus := 0
	for _, value := range nums {
		if value < 0 {
			minus++
		}
	}

	if k < minus {
		for i := 0; i < k; i++ {
			nums[i] *= -1
		}
	} else {
		for i := 0; i < minus; i++ {
			nums[i] *= -1
			k--
		}

		if k%2 != 0 {
			absmin := math.MaxInt
			absindex := 0

			for i, value := range nums {
				if utils.Abs(value) < absmin {
					absmin = utils.Abs(value)
					absindex = i
				}
			}

			nums[absindex] *= -1
		}
	}

	sum := 0
	for _, value := range nums {
		sum += value
	}

	return sum
}

func LargestSumAfterKNegations2(nums []int, k int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] *= -1
			k--
		}
	}

	if k%2 != 0 {
		minIndex := 0

		for i := 1; i < len(nums); i++ {
			if utils.Abs(nums[i]) < utils.Abs((nums[minIndex])) {
				minIndex = i
			}
		}

		nums[minIndex] *= -1
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}
