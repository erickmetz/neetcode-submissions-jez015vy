func twoSum(nums []int, target int) []int {
	length := len(nums)
	solveCache := make(map[int]int)

	for idx, val := range(nums) {
		// calc index from other end
		rIdx := length - 1 - idx
		rVal := nums[rIdx]


		// return immediately if already correct
		tCheck := val + rVal
		if tCheck == target {
			return []int{
				idx,
				rIdx,
			}
		}

		solved := false
		solution := []int{}
		remainder := target-val
		rRemainder := target-rVal

		if _, ok := solveCache[remainder]; ok {
			solved = true
			solution = []int{idx, solveCache[remainder]}
		} else if _, ok := solveCache[rRemainder]; ok {
			solved = true
			solution = []int{rIdx, solveCache[rRemainder]}
		}

		if solved {
			if solution[0] > solution[1] {
				return []int{solution[1], solution[0]}
			} else {
				return []int{solution[0], solution[1]}
			}
		}	
			
		solveCache[val] = idx
		solveCache[rVal] = rIdx

		if rIdx <= idx {
			break
		}

	}

	return []int{}
}