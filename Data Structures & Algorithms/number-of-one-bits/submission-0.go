func hammingWeight(n int) int {
	counter := 0
	ptr := n

	for c := 0; c < 32; c++ {
		if (ptr & 1) == 1 {
			counter++
		}
		ptr >>= 1
	}

	return counter
}
