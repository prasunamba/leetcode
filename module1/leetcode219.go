package module1

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && int(math.Abs(float64(i-j))) == k {
				return true
			}
		}
	}
	return false
}
func optimizedcontainsNearbyDuplicate(nums []int, k int) bool {
	map1 := make(map[int]int) // Stores the last index of each number

	for i, num := range nums {
		if prevIndex, exists := map1[num]; exists && i-prevIndex <= k {
			return true
		}
		map1[num] = i // Update the index of the number
	}
	return false
}
