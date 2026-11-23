package main

import (
	"fmt"
	"main/utils"
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
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

func main() {
	fmt.Println("", largestSumAfterKNegations([]int{-8, 3, -5, -3, -5, -2}, 6))
}
