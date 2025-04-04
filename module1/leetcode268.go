package module1

func missingNumber(nums []int) int {
	numbers := len(nums)
	sum := (numbers * (numbers + 1)) / 2
	total := 0
	for _, num := range nums {
		total += num
	}
	return sum - total
}
