package module1

func AlternatingSum(nums []int) int {
	sum := 0
	for i, value := range nums {
		if i%2 == 0 {
			sum += value
		} else {
			sum -= value
		}
	}
	return sum

}
