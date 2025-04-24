package module1

func hash(nums []int) int {
	map1 := make(map[int]struct{})
	for _, value := range nums {
		map1[value] = struct{}{}
	}
	return len(map1)
}
func checkValid(matrix [][]int) bool {
	l := len(matrix)
	var rowcount int
	for _, row := range matrix {
		distinct := hash(row)
		if l != distinct {
			return false
		}
		rowcount++
	}

	var colcount int
	for i := 0; i < l; i++ {
		col := []int{}
		for j := 0; j < l; j++ {
			col = append(col, matrix[j][i])
		}
		distinct := hash(col)
		if distinct != l {
			return false
		}
		colcount++
	}
	if rowcount == l && colcount == l {
		return true
	}
	return false
}
