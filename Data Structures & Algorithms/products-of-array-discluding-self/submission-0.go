func productExceptSelf(nums []int) []int {
	nLen := len(nums)
	products := make([]int, nLen)

	for n := range(nLen) {
		product := 1
		for m := range(nLen) {
			if n == m {
				continue
			}
			product = product * nums[m]
		}
		products[n] = product
	}

	return products
}
