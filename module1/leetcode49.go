package module1

import (
	"fmt"
	"sort"
)

func sortstr(value string) string {
	valuerune := []rune(value)
	sort.Slice(valuerune, func(i, j int) bool {
		return valuerune[i] > valuerune[j]
	})
	fmt.Println("", valuerune)
	return string(valuerune)
}
func groupAnagrams(strs []string) [][]string {
	map1 := make(map[string][]string)
	for idx, value := range strs {
		fmt.Println("", idx, value)
		sorted := sortstr(value)
		map1[sorted] = append(map1[sorted], value)
	}
	fmt.Println("", map1)
	result := [][]string{}
	for _, val := range map1 {
		result = append(result, val)
	}
	return result
}
