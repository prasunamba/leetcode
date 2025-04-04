package module1

import "fmt"

func countBits(n int) []int {
	resultslice := []int{}
	for i := 0; i <= n; i++ {
		result := countones(i)
		resultslice = append(resultslice, result)
	}
	return resultslice
}
func countones(num int) int {
	binarynum := fmt.Sprintf("%b", num)
	slicenum := []rune(binarynum)
	onecount := 0
	for _, val := range slicenum {
		if val == '1' {
			onecount++
		}
	}
	return onecount
}
