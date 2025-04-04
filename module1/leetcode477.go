package module1

import "fmt"

func totalHammingDistance(nums []int) int {
	sum, result := 0, 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("pair", nums[i], nums[j])
			result = HammingDistance(nums[i], nums[j])
			fmt.Println("result", result)
			sum += result
		}
		fmt.Println("sum", sum)
	}
	return sum
}
func totalHammingDistance_optimized(nums []int) int {
	total := 0
	n := len(nums)

	// Iterate through all 32 bits (for 32-bit integers)
	for bit := 0; bit < 32; bit++ {
		countOnes := 0
		for _, num := range nums {
			if (num>>bit)&1 == 1 { // Check if the bit is set
				countOnes++
			}
		}
		countZeros := n - countOnes
		total += countOnes * countZeros // Pairs contributing to Hamming distance
	}

	return total
}
