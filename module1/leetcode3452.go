package module1

func SumOfGoodNumbers(nums []int, k int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		good := true

		left := i - k
		right := i + k

		if left >= 0 && nums[i] <= nums[left] {
			good = false
		}

		if right < len(nums) && nums[i] <= nums[right] {
			good = false
		}

		if good {
			sum += nums[i]
		}
	}

	return sum
}
