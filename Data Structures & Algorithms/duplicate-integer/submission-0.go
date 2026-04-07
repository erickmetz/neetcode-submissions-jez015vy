func hasDuplicate(nums []int) bool {
    record := make(map[int]bool, 0)

	for _, i := range(nums) {
		if _, ok := record[i]; ok {
			return true
		}

		record[i] = true
	}

	return false
}