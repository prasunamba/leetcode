/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import "fmt"

func push(nums []int, i int) []int {
	nums = append(nums, i)
	return nums
}

func pop(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nums, fmt.Errorf("cannot pop from an empty slice")
	}
	nums = nums[:len(nums)-1]
	return nums, nil
}

func seek(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("cannot seek from an empty slice")
	}
	return nums[len(nums)-1], nil
}

func stackops() {
	var nums []int
	newnums := push(nums, 1)
	newnums = push(newnums, 2)
	fmt.Println("after push", newnums)

	// Seek the last element
	val, err := seek(newnums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("seek", val)
	}

	// Pop elements and handle errors
	newnums, err = pop(newnums)
	if err != nil {
		fmt.Println(err)
	}
	newnums, err = pop(newnums)
	if err != nil {
		fmt.Println(err)
	}
	newnums, err = pop(newnums) // Trying to pop from an empty slice
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after pop", newnums)
}
