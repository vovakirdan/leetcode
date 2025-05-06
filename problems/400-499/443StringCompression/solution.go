package stringcompression

import "strconv"

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	write := 0 // куда записывать
	read := 0  // где читать

	for read < len(chars) {
		ch := chars[read]
		count := 0

		for read < len(chars) && ch == chars[read] {
			count++
			read++
		}

		chars[write] = ch
		write++

		if count > 1 {
			t := strconv.Itoa(count)
			for _, symb := range t {
				chars[write] = byte(symb)
				write++
			}
		}
	}
	return write
}
