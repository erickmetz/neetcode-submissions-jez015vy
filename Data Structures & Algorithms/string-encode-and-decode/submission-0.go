type Solution struct{
	startIdx []int
	lens []int
}

func (s *Solution) Encode(strs []string) string {
	encoded := ""

	offset := 0
	for _, str := range(strs) {
		s.startIdx = append(s.startIdx, offset)
		s.lens = append(s.lens, len(str))
		encoded = encoded + str
		offset += len(str)
	}

	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	decoded := make([]string, len(s.startIdx))

	for i, _ := range(s.startIdx) {
		decoded[i] = encoded[s.startIdx[i]:s.startIdx[i]+s.lens[i]]
	}

	return decoded
}
