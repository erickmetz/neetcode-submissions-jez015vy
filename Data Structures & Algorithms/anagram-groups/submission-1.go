func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	resultGroup := make(map[string][]string)

	for _, s := range(strs) {
		sortedS := sortString(s)
		resultGroup[sortedS] = append(resultGroup[sortedS], s)
	}

	var result [][]string
	for _, group := range resultGroup {
		result = append(result, group)
	}

	return result

}