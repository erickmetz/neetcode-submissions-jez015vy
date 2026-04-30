func isPalindrome(s string) bool {
	// normalize - convert to lowercase
	// strip whitespace and symbols and lowercase conversion
	s = strings.ToLower(s)
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = reg.ReplaceAllString(s, "")

	result := true

	sLen := len(s)
	for i := range(s) {
		j := sLen-1-i
		if i >= j || j <= i {
			break
		}

		if s[i] != s[j] {
			result = false
			break
		}
	}

	return result
}
