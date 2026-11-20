package module1

import (
	"fmt"
	"main/utils"
	"sort"
)

func Mergeintervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	result = append(result, intervals[0])
	for _, num := range intervals[1:] {
		last := result[len(result)-1]
		if last[1] >= num[0] {
			last[1] = utils.Maxm(num[1], last[1])

		} else {
			result = append(result, num)
		}
	}

	fmt.Println("", intervals, "\n", result)
	return intervals

}
