package module1

import "fmt"

func Tejasquestion() {
	list := []int{8, 5, 7, 7, 8, 8, 10, 7, 7, 8, 8, 10}
	num := 8
	var count []int
	for index, n := range list {
		if num == n {
			count = append(count, index)
		}
	}
	// Print results safely
	if len(count) > 0 {
		fmt.Println("Indices:", count)
		fmt.Println("First occurrence:", count[0])
		fmt.Println("Last occurrence:", count[len(count)-1])
	} else {
		fmt.Println(num, "not found in list")
	}
}

func Interviewsolution1() {
	list := []int{3, 1, 2, 5, 12, 9, 45, 2, 3, 1, 2, 4, 5, 2}
	target := 2

	first, last := -1, -1
	for index, value := range list {
		if value == target {
			if first == -1 {
				first = index
			}
			last = index

		}
	}
	if first == -1 {
		fmt.Println("Target not found")
	} else {
		fmt.Printf("First occurrence: %d, Last occurrence: %d\n", first, last)
	}

}
func Interviewsolution2() {
	list := []int{3, 1, 2, 5, 12, 9, 45, 2, 3, 1, 2, 4, 5, 2}
	target := 2
	setfirst, setlast := false, false
	first, last := -1, -1
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		if list[i] == target || list[j] == target {
			if first == -1 && !setfirst {
				first = i
			}
			if last == -1 && !setlast {
				last = j
			}
		}
	}
	fmt.Println("", first, last)

}
