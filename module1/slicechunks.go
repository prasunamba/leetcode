package module1

import "fmt"

func SliceChunks() {
	slice := []int{2, 1, 3, 55, 22, 2, 1, 3, 55, 22}
	size := 3
	var result [][]int
	for size < len(slice) {
		slice, result = slice[size:], append(result, slice[:size])
	}
	result = append(result, slice)
	fmt.Println("", result)
}

func slicestring(s string, k int) string {
	for i := 0; i < len(s); i = i + 2*k {
		end := i + 4
		if end > len(s) {
			end = len(s)
		}
		fmt.Println("", s[i:end])
	}
	return ""
}
