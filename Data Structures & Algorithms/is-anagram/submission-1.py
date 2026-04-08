class Solution:
	def getFrequencies(self, fs: str) -> dict:
		frequencies = {}

		for char in fs:
			if char in frequencies:
				frequencies[char] = frequencies[char] + 1
			else:
				frequencies[char] = 1

		return frequencies

	def isAnagram(self, s: str, t: str) -> bool:
		# differing lengths cannot be anagrams
		if len(s) != len(t):
			return False

		sFreq = self.getFrequencies(s)
		tFreq = self.getFrequencies(t)
		
		return (sFreq == tFreq)
		

		