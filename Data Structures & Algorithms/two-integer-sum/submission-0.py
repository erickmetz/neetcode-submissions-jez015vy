class Solution:
	def twoSum(self, nums: List[int], target: int) -> List[int]:
		length = len(nums)
		solveCache = {}
		
		for idx, val in enumerate(nums):
			remainder = target-val

			if remainder in solveCache:
				if solveCache[remainder] < idx:
					return [solveCache[remainder], idx]
				else:
					return [idx, solveCache[remainder]]
			
			solveCache[val] = idx
		

        