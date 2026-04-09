import "maps"

func mapFrequencies(str string) map[rune]int {
	// track character frequencies
	frequencies := make(map[rune]int, 0)
	for _, char := range(str) {
		if _, ok := frequencies[char]; ok {
			frequencies[char] = frequencies[char]+1
		} else {
			frequencies[char] = 1
		}
	}

	return frequencies
}

func isAnagram(s string, t string) bool {
	// unequal length are not anagrams
	if len(s) != len(t) {
		return false
	}

	// track character frequencies
	sFreq := mapFrequencies(s)
	tFreq := mapFrequencies(t)

	return (maps.Equal(sFreq, tFreq))
}

func groupAnagrams(strs []string) [][]string {
	// function result return var
	resultGroups := [][]string{}

	// only check same-length strings for being anagrams
	lengthGroups := make(map[int][]string, 0)
	// group by length
	for _, str := range(strs) {
		sLen := len(str)
		if _, ok := lengthGroups[sLen]; ok {
			lengthGroups[sLen] = append(lengthGroups[sLen], str)
		} else {
			lengthGroups[sLen] = []string{str}
		}
	}

	// checked cache avoids anagram checking words we've already seen once
	checkedCache := make(map[string]bool)

	// process by length
	for _, lengthStrs := range(lengthGroups) {	
		/* 
		  I dont care for the loop within a loop 
		  for this. Like some "two sum" solution 
		  where this is being processed from the "left"
		  and "right" positional indexes might 
		  follow this attempt
		*/
		for _, s := range lengthStrs {
			if _, ok := checkedCache[s]; ok {
				continue
			}
			resultGroup := make(map[string][]string, 0)
			for _, t := range lengthStrs {
				if isAnagram(s,t) {
					if resultGroup[s] == nil {
						resultGroup[s] = []string{t}
					} else {
						resultGroup[s] = append(resultGroup[s], t)
					}
					checkedCache[s] = true
					checkedCache[t] = true
				}
			}
			resultGroups = append(resultGroups, resultGroup[s])
		}
	}

	return resultGroups
}
