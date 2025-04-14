package mergestringsalternately

func mergeAlternately(word1 string, word2 string) string {
	r1 := []rune(word1)
	r2 := []rune(word2)

	res := []rune{}
	i := 0

	for i < len(r1) || i < len(r2) {
		if i < len(r1) {
			res = append(res, r1[i])
		}
		if i < len(r2) {
			res = append(res, r2[i])
		}
		i++
	}

	return string(res)
}
