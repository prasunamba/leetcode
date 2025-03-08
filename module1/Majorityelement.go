package module1

import (
	"fmt"
	"math"
)

// Boyer-Moore's Voting Algorithm[for occurences more then n/2], which runs in O(n) time and O(1) space.
func majorityElement(nums []int) int {
	count, candidate := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// more space complexity O(n) , time complexity O(n)
func MyapproachforMajorityElement() {
	nums := []int{3, 2, 3}
	map1 := make(map[int]int)
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		map1[nums[i]]++
		if i != j {
			map1[nums[j]]++
		}
	}
	var keys []int
	for key := range map1 {
		keys = append(keys, key)
	}
	for _, num := range keys {
		if map1[num] >= int(math.Ceil(float64(len(nums))/2.0)) {
			fmt.Println("", num)
		}
	}
}
