func twoSum(nums []int, target int) []int {
	solveCache := make(map[int]int)

	for idx, val := range(nums) {
		remainder := target-val
		if _, ok := solveCache[remainder]; ok {
			return []int{
				solveCache[remainder], 
				idx,
			}
		}

		solveCache[val] = idx
	}

	return []int{}
}
