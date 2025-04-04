package module1

import (
	"fmt"
	"slices"
)

func intersection(nums1 []int, nums2 []int) []int {
	resultslice := []int{}
	for _, val := range nums1 {
		if slices.Contains(nums2, val) {
			if !slices.Contains(resultslice, val) {
				resultslice = append(resultslice, val)
				fmt.Println("result", resultslice)
			}
		}
	}
	return resultslice
}
