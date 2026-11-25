package module1

func CreateTargetArray(nums []int, index []int) []int {

	var target []int

	for i := 0; i < len(nums); i++ {

		idx := index[i]
		value := nums[i]

		target = append(target[:idx],
			append([]int{value}, target[idx:]...)...)
	}

	return target
}
