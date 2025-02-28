package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{2, 1, 3, 4, 6, 5}
	/* sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	}) */
	sort.Sort(sort.Reverse(sort.IntSlice(slice1)))
	fmt.Println("", slice1)
}
