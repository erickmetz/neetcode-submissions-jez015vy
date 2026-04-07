class Solution:
	def hasDuplicate(self, nums: List[int]) -> bool:
		record = {}
		for num in nums:
			if num in record:
				return True
			
			record[num] = True

		return False