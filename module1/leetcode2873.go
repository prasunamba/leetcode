package module1

func maximumTripletValue(nums []int) int64 {
	maxresult := 0
	for i := 0; i <= len(nums)-2; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				result := (nums[i] - nums[j]) * nums[k]
				if result > maxresult {
					maxresult = result
				}
			}

		}
	}
	if maxresult < 0 {

		return 0
	}
	return int64(maxresult)
}
