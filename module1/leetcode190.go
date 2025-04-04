package module1

func reverseBits(num uint32) uint32 {
	var result uint32 = 0

	for i := 0; i < 32; i++ {
		bit := num & 1               // Extract the last bit
		result = (result << 1) | bit // Shift left and add the extracted bit
		num >>= 1                    // Right shift num to process the next bit
	}
	return result
}
