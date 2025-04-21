package validpalindromeii

func validPalindrome(s string) bool {
	isPalin := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalin(left+1, right) || isPalin(left, right-1)
		}
		left++
		right--
	}

	return true
}
