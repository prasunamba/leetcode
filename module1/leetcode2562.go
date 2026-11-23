package module1

import "strconv"

func FindTheArrayConcVal(nums []int) int64 {
	var result int64
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			result += int64(nums[i])
		} else {

			temp := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
			tem, _ := strconv.ParseInt(temp, 10, 64)
			result += tem
		}

	}
	return result
}
