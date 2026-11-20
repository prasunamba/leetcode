package module1

func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	numsmap := make(map[int]int)
	var result []int
	setmap1 := make(map[int]bool)
	for _, value := range nums1 {
		setmap1[value] = true
	}
	for value, _ := range setmap1 {
		numsmap[value]++
	}
	setmap2 := make(map[int]bool)
	for _, value := range nums2 {
		setmap2[value] = true
	}
	for value, _ := range setmap2 {
		numsmap[value]++
	}
	setmap3 := make(map[int]bool)
	for _, value := range nums3 {
		setmap3[value] = true
	}
	for value, _ := range setmap3 {
		numsmap[value]++
	}

	for key, value := range numsmap {
		if value >= 2 {
			result = append(result, key)
		}
	}
	return result
}
