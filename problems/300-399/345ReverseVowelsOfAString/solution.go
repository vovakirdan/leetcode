package reversevowelsofastring

func reverseVowels(s string) string {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	slice := []rune(s)
	i, j := 0, len(slice)-1
	for i < j {
		if _, ok := vowels[slice[i]]; !ok {
			i++
		} else if _, ok := vowels[slice[j]]; !ok {
			j--
		} else {
			slice[i], slice[j] = slice[j], slice[i]
			i++
			j--
		}
	}
	return string(slice)
}
