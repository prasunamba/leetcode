package module1

func IsValid(s string) bool {
	mymap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	var stack []string

	for _, lit := range s {
		char := string(lit)

		// If the character is an opening bracket, push it onto the stack
		if _, exists := mymap[char]; exists {
			stack = append(stack, char)
		} else {
			// If stack is empty, it means there's an unmatched closing bracket
			if len(stack) == 0 {
				return false
			}

			// Get the last opening bracket
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop last element

			// Check if it matches the expected closing bracket
			if mymap[last] != char {
				return false
			}
		}
	}

	// If stack is empty, all brackets are matched correctly
	return len(stack) == 0
}
