# Increasing Triplet Subsequence

Level: Medium

## 🧠 Задача:
> Дан массив целых чисел `nums`.  
> Нужно определить, существует ли **увеличивающаяся подпоследовательность из трёх элементов** — такие индексы `i < j < k`, что `nums[i] < nums[j] < nums[k]`.  
> Вернуть `true`, если такая тройка найдена, иначе — `false`.

---

## 📌 Идея:
- Ищем два наименьших числа `first` и `second`
- Если найдём `num > second` — последовательность из трёх элементов найдена

---

## 📏 Структура:
- `first int` — минимальное значение
- `second int` — второе минимальное значение

---

## 🔁 Шаги алгоритма:

1. Инициализируем:
   - `first = MaxInt`
   - `second = MaxInt`
2. Проходим по `nums`:
   - Если `num <= first` → обновить `first`
   - Иначе если `num <= second` → обновить `second`
   - Иначе → `num > second` → возвращаем `true`
3. Если цикл завершён без нахождения тройки — возвращаем `false`

---

## ⏱️ Сложность:
- Время: O(n)
- Память: O(1)

---

## 📄 Пример:

Вход:
```go
nums := []int{2, 1, 5, 0, 4, 6}
```

Выход:
```go
true // Тройка: 0 < 4 < 6
```

---

## 📝 Решение:

```go
func increasingTriplet(nums []int) bool {
	first := int(^uint(0) >> 1)  // MaxInt
	second := int(^uint(0) >> 1) // MaxInt

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			// Найдено: num > second → тройка существует
			return true
		}
	}

	return false
}
```
