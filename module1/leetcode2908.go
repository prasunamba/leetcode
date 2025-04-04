package module1

import (
	"fmt"
	"math"
)

func minimumSum(nums []int) int {
	minsum := math.MaxInt
	set := false
	for i := 0; i <= len(nums)-2; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[k] < nums[j] {
					set = true
					sum := nums[i] + nums[j] + nums[k]
					fmt.Println("sum", sum)
					if sum < minsum {
						minsum = sum
					}
				}

			}

		}
	}
	if set {
		return minsum
	}
	return -1
}
