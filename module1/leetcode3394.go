package module1

import (
	"fmt"
	"sort"
)

func checkValidCuts(n int, rectangles [][]int) bool {
	xslice := make([][]int, len(rectangles))
	yslice := make([][]int, len(rectangles))
	for idx, s := range rectangles {
		x := []int{s[0], s[2]}
		xslice[idx] = x
	}
	for idx, s := range rectangles {
		y := []int{s[1], s[3]}
		yslice[idx] = y
	}
	sort.Slice(xslice, func(i, j int) bool {
		return xslice[i][0] < xslice[j][0]
	})
	sort.Slice(yslice, func(i, j int) bool {
		return yslice[i][0] < yslice[j][0]
	})
	fmt.Println("xslice", xslice)
	fmt.Println("yslice", yslice)
	// fmt.Println("", count_non_overlap(xslice))
	// fmt.Println("", count_non_overlap(yslice))
	result := max(count_non_overlap(xslice), count_non_overlap(yslice))
	fmt.Println("finalresult", result)
	return result >= 3
}
func count_non_overlap(x [][]int) int {
	result := [][]int{x[0]} // Start with the first interval
	count := 1              // At least one section exists

	for _, value := range x[1:] {
		last := result[len(result)-1]

		if value[0] < last[1] { // Overlapping interval
			result[len(result)-1][1] = max(last[1], value[1]) // Merge
		} else { // Non-overlapping interval
			result = append(result, value)
			count++ // Increase section count
		}
	}

	fmt.Println("Merged Intervals:", result)
	return count
}
