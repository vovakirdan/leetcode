package issubsequence

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}

	return i == len(s)
}
