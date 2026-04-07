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
