package module1

func optimizedcontainsDuplicate(nums []int) bool {
	map1 := make(map[int]bool)
	for _, n := range nums {
		if map1[n] {
			return true // Found a duplicate, return early
		}
		map1[n] = true
	}
	return false
}
func containsDuplicate(nums []int) bool {
	map1 := make(map[int]int)
	for _, n := range nums {
		map1[n]++
	}
	for _, val := range map1 {
		if val > 1 {
			return true
		}
	}
	return false
}
