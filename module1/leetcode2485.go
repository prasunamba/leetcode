package module1

import (
	"fmt"
	"math"
)

// 40ms
func ssum(n1, n2 int) int {
	sum := 0
	for i := n1; i <= n2; i++ {
		sum += i
	}
	return sum
}
func PivotInteger(n int) int {
	pivot := -1
	if n == 1 {
		return 1
	}
	j := 1
	for j < n {
		sum, sum1 := 0, 0
		sum = ssum(1, j)
		/* for i := 1; i <= j; i++ {
			sum += i
		} */
		sum1 = ssum(j, n)
		/* for i := j; i <= n; i++ {
			sum1 += i
		} */
		fmt.Println("", sum, sum1)
		if sum == sum1 {
			pivot = j
		}
		j++
	}

	return pivot
}

// -------------------------------------------------
// time 40ms
func PivotInteger2(n int) int {
	pivot := -1
	if n == 1 {
		return 1
	}
	j := 1
	for j < n {
		sum, sum1 := 0, 0
		for i := 1; i <= j; i++ {
			sum += i
		}
		for i := j; i <= n; i++ {
			sum1 += i
		}
		fmt.Println("", sum, sum1)
		if sum == sum1 {
			pivot = j
		}
		j++
	}

	return pivot
}

// -----------------------------------------------------

// time =0 ms
// formula pivot2=(n(n+1))/2
func pivotInteger3(n int) int {
	total := n * (n + 1) / 2

	pivot := int(math.Sqrt(float64(total)))

	if pivot*pivot == total {
		return pivot
	}

	return -1
}
