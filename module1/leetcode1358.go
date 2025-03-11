package module1

import (
	"fmt"
	"strings"
)

// not a optimized one TLE
func NumberOfSubstrings() {
	s := "abcabc"
	count := 0
	containsA := 0
	map1 := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			word := s[i:j]
			// fmt.Print("  ", word)

			count++
			map1[word]++
		}
	}
	for key, val := range map1 {
		if strings.Contains(key, "a") && strings.Contains(key, "b") && strings.Contains(key, "c") {
			fmt.Print(" ", key)
			containsA += val
		}
	}
	fmt.Println("\n map1", map1)
	fmt.Println("total substrings", count) // n*(n+1) / 2
	fmt.Println("distinct substrings", len(map1))
	fmt.Println("contains A", containsA)
}
func numberOfSubstringsoptimized(s string) int {
	count := 0
	lastSeen := map[byte]int{'a': -1, 'b': -1, 'c': -1} // Track last index of 'a', 'b', 'c'

	for i := 0; i < len(s); i++ {
		lastSeen[s[i]] = i // Update the last seen index of the character

		// The number of substrings that end at i and contain 'a', 'b', 'c' at least once
		count += 1 + min(lastSeen['a'], lastSeen['b'], lastSeen['c'])
	}

	return count
}

// Helper function to find the minimum of three values
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
