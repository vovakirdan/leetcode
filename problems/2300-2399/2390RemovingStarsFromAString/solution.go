package removingstarsfromastring

func removeStars(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if ch != '*' {
			stack = append(stack, ch)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(stack)
}
