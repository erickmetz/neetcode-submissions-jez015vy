func countBits(n int) []int {
	results := make([]int, n+1)

	idx := 0
	for seq := range(n+1) {
		counter := 0
		for c := 0; c < 32; c++ {
			if (seq & 1) == 1 {
				counter++
			}
			seq >>= 1
		}
		results[idx] = counter
		idx += 1
	}

	return results

}
