package module1

func LongestSubArraySlidingWindowBRUTEFORCE(nums []int, k int) int {
	sum := 0
	maxlen := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum += nums[j]

			if sum <= k {
				maxlen = max(maxlen, j-i+1)
			} else {
				break
			}
		}
	}
	return maxlen
}
