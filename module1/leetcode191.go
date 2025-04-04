package module1

import "fmt"

func hammingWeight(n int) int {
	binaryn := fmt.Sprintf("%b", n)
	fmt.Println("", binaryn)
	slicebinn := []rune(binaryn)
	hammweight := 0
	for _, val := range slicebinn {
		if val == '1' {
			hammweight++
		}
	}
	return hammweight
}
