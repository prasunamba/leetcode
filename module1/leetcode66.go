package module1

import "fmt"

// optimized one
func PlusOne() []int {
	digits := []int{7, 9, 9}

	n := len(digits)
	carry := 1 // We need to add 1

	for i := n - 1; i >= 0; i-- {
		digits[i] += carry
		fmt.Println("", digits)
		fmt.Println("", digits[i])
		if digits[i] == 10 { // If a digit becomes 10, set it to 0 and carry forward
			digits[i] = 0
			carry = 1
		} else {
			carry = 0 // No more carry needed, stop processing
			break
		}
	}

	// If there's still a carry after the loop (e.g., 999 -> 1000), add a new digit
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
