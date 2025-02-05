package main

import "fmt"

func enqueue(nums []int, i int) []int {
	return append(nums, i)
}
func dequeue(nums []int) []int {
	return nums[1:]
}
func queue() {
	var queue []int
	for i := 0; i < 5; i++ {
		queue = enqueue(queue, i)
	}
	queue = enqueue(queue, 11)
	fmt.Println("after enqueue", queue)
	for i := 0; i < 4; i++ {
		queue = dequeue(queue)
	}
	fmt.Println("after dequeue", queue)

}
