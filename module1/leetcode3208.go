package module1

import "fmt"

// alternating groups
func Leetcode3208() {
	colors := []int{0, 1, 0, 1, 0}
	k := 3
	lencolors := len(colors)
	result := 0

	// Count transitions in the first window
	// by default count the first window [0,1,0]
	count := 0
	for i := 1; i < k; i++ {
		if colors[i-1] != colors[i] {
			count++
		}
	}
	if count == k-1 {
		result++
	}

	// Slide the window
	for i := 1; i < lencolors; i++ {
		// Remove effect of outgoing element
		// in the sliding window the first pair is already counted so decrement the pair
		if colors[(i-1)%lencolors] != colors[(i)%lencolors] {
			count--
		}
		// Add effect of incoming element
		// in the sliding window the second pair is new so compare the elements of pair
		if colors[(i+k-2)%lencolors] != colors[(i+k-1)%lencolors] {
			count++
		}
		// Check if the window meets criteria
		if count == k-1 {
			result++
		}
	}

	fmt.Println("Result:", result)
}
