package module1

func minimumIndex(nums []int) int {
	// Step 1: Find the dominant element using Boyer-Moore Majority Vote Algorithm
	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// Step 2: Count the total occurrences of the dominant element
	totalCount := 0
	for _, num := range nums {
		if num == candidate {
			totalCount++
		}
	}

	// Step 3: Iterate to find the minimum valid index
	leftCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			leftCount++
		}
		rightCount := totalCount - leftCount

		if leftCount > (i+1)/2 && rightCount > (len(nums)-i-1)/2 {
			return i
		}
	}

	return -1
}
