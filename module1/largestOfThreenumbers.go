package module1

import "fmt"

func LargestOfThreeNum(a, b, c int) {
	if a > b && a > c {
		fmt.Printf("%d is largest \n", a)
	} else if b > c {
		fmt.Printf("%d is largest \n ", b)
	} else {
		fmt.Printf("%d is largest \n", c)
	}
}

func OptimizedLargestOfThreeNum(a, b, c int) {
	largest := a
	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}
	fmt.Printf("%d is largest \n", largest)
}
