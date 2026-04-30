func reverseBits(n int) int {
	reversed := 0

	for c := 0; c < 32; c++ {
		if (n & 1) == 1 {
			andBit := (1 << (31-c))
			reversed = reversed ^ andBit
		}
		n >>= 1
	}

	return reversed
}
