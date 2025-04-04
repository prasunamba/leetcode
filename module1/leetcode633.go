package module1

import "math"

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c))) // Two pointers: 0 and sqrt(c)
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++ // Increase sum
		} else {
			right-- // Decrease sum
		}
	}
	return false
}
