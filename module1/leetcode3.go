package module1

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	longest := 0

	for i := 0; i < n; i++ {
		charMap := make(map[byte]bool)
		for j := i; j < n; j++ {
			fmt.Printf("%c \n ", s[j])
			if charMap[s[j]] {
				break // Stop if a duplicate is found
			}
			charMap[s[j]] = true
			longest = max(longest, j-i+1)
		}
	}

	return longest
}

func optimizedlengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // Stores last index of characters
	left, maxLength := 0, 0

	for right := 0; right < len(s); right++ {
		if lastIndex, found := charIndex[s[right]]; found && lastIndex >= left {
			left = lastIndex + 1 // Move left pointer after last occurrence
		}

		charIndex[s[right]] = right // Update last seen index
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
