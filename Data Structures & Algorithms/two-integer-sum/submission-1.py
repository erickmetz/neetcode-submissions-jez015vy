class Solution:
	def twoSum(self, nums: List[int], target: int) -> List[int]:
		length = len(nums)
		solveCache = {}
		
		for idx, val in enumerate(nums):
			# calc index from other end
			ridx = length - 1 - idx
			rval = nums[ridx]

			# return immediately if alreadyu correct
			tCheck = val + rval
			if tCheck == target and idx != ridx:
				return sorted([idx,ridx])

			solved = False
			solution = []
			remainder = target-val
			rremainder = target-rval
			if remainder in solveCache:
				solved = True
				solution = [idx, solveCache[remainder]]
			elif rremainder in solveCache:
				solved = True
				solution = [ridx, solveCache[rremainder]]

			if solved:
				if solution[0] > solution[1]:
					return [solution[1], solution[0]]
				else:
					return [solution[0], solution[1]]
			
			solveCache[val] = idx
			solveCache[rval] = ridx

			if idx >= ridx:
				break
