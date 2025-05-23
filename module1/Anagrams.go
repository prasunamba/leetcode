package module1

import (
	"slices"
	"sort"
	"strings"
)

func Anagrams(name1, name2 string) bool {
	name1 = strings.ToLower(name1)
	name2 = strings.ToLower(name2)
	name1slice := []rune(name1)
	name2slice := []rune(name2)
	slices.Sort(name1slice)
	sort.Slice(name2slice, func(i, j int) bool { return name2slice[i] < name2slice[j] })

	return string(name1slice) == string(name2slice)
}

func Anagramsoptimized(name1, name2 string) bool {

	name1 = strings.ToLower(name1)
	name1 = strings.TrimSpace(name1)
	name2 = strings.ToLower(name2)
	name2 = strings.TrimSpace(name2)

	if len(name1) != len(name2) {
		return false
	}
	counts := make(map[rune]int)
	for _, ch := range name1 {
		counts[ch]++
	}
	for _, ch := range name2 {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}
