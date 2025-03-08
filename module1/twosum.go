package module1

func Twosum() []int {
	nums := []int{2, 0, 11, 15}
	target := 9
	map1 := make(map[int]int)
	// fmt.Println("Indices:", twoSum(nums, target)) // Output: [0, 1]
	for idx, num := range nums {
		if i, n := map1[target-idx]; n {
			return []int{idx, i}
		}
		map1[num] = idx
	}
	return nil
}
