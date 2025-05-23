package module1

func lengthAfterTransformations(s string, t int) int {
	const mod = 1_000_000_007

	// Initialize a frequency array for each character 'a' to 'z'
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

	for i := 0; i < t; i++ {
		newCount := make([]int, 26)
		for j := 0; j < 26; j++ {
			if j == 25 {
				// 'z' transforms to 'a' and 'b'
				newCount[0] = (newCount[0] + count[25]) % mod
				newCount[1] = (newCount[1] + count[25]) % mod
			} else {
				// Other characters shift to the next character
				newCount[j+1] = (newCount[j+1] + count[j]) % mod
			}
		}
		count = newCount
	}

	// Sum the counts to get the total length
	result := 0
	for _, val := range count {
		result = (result + val) % mod
	}
	return result
}
