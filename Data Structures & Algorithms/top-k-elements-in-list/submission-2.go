func topKFrequent(nums []int, k int) []int {
	frequencies := make(map[int]int, 0)

	// collect frequencies
	for _, num := range(nums) {
		if _, ok := frequencies[num]; ok {
			frequencies[num] += 1
		} else {
			frequencies[num] = 1
		}
	}

	// create and sort return slice
	keys := make([]int, 0)
	for k := range frequencies {	
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return frequencies[keys[i]] > frequencies[keys[j]]
	})

	topNums := keys[0:k]

	return topNums
}
