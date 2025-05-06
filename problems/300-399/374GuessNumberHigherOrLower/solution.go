package guessnumberhigherorlower

// глобально эмулируем "загаданное" число
var picked int

// эмуляция API
func guess(num int) int {
	if num > picked {
		return -1
	} else if num < picked {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1 // по условию задачи до сюда не дойдёт
}
