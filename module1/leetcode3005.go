package module1

func MaxFrequencyElements(nums []int) int {
	numsmap := make(map[int]int)
	for _, value := range nums {
		numsmap[value]++
	}
	max := 0
	for _, value := range numsmap {
		if value > max {
			max = value
		}
	}

	result := 0
	for _, value := range numsmap {
		if value == max {
			result += value
		}
	}
	return result
}
