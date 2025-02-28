package module1

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	list3 := append(nums1, nums2...)
	sort.Ints(list3)
	fmt.Println("list3", list3)
	fmt.Println("len(list3)", len(list3)/2)
	if len(list3)%2 == 1 {
		index := float64(list3[len(list3)/2])
		fmt.Println("-", index)
		return index
	} else {
		index := float64(list3[len(list3)/2]+list3[len(list3)/2-1]) / 2
		fmt.Println("+-", index)
		return index
	}

}
