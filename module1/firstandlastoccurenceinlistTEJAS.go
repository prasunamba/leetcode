package module1

import (
	"fmt"
	"golang/utils"
)

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
	list := []int{5, 7, 7, 8, 8, 10}
	target := 8

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

func Interviewsolution2() []int {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	firstindex := -1
	lastindex := -1
	// lastflag := false
	// firstflag := false
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if target == nums[i] {
			if firstindex == -1 {
				firstindex = i
				// firstflag = true
				fmt.Println("i", firstindex)
			} else {
				firstindex = utils.Minm(i, firstindex)
			}
			if lastindex == -1 {
				lastindex = i
			} else if i > lastindex {
				lastindex = utils.Maxm(i, lastindex)
				fmt.Println("i last", lastindex)
			}

		}
		if target == nums[j] {
			if lastindex == -1 {
				lastindex = j
				// lastflag = true
				fmt.Println("j", lastindex)
			} else {
				lastindex = utils.Maxm(j, lastindex)
			}
			if firstindex == -1 {
				fmt.Println("j -1", j)
				firstindex = j
			} else if j < firstindex {
				firstindex = utils.Minm(j, firstindex)
				fmt.Println("j first", firstindex)
			}
		}
	}
	if firstindex == -1 || lastindex == -1 {
		if firstindex == -1 {
			firstindex = utils.Maxm(firstindex, lastindex)
		} else {
			lastindex = utils.Maxm(firstindex, lastindex)
		}
	}
	return []int{firstindex, lastindex}

}
