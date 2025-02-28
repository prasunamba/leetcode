package module1

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) {
	n := len(nums)
	var result [][]int
	seen := make(map[string]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					slice := []int{nums[i], nums[j], nums[k]}
					sort.Ints(slice)
					key := strconv.Itoa(slice[0]) + "," + strconv.Itoa(slice[1]) + "," + strconv.Itoa(slice[2])

					// Add to result only if not already present
					if !seen[key] {
						result = append(result, slice)
						seen[key] = true
					}
				}
			}
		}
	}
	fmt.Println("", result)
}
