# Find Peak Element

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/find-peak-element)

## 🧠 Задача:
> Дан массив целых чисел `nums` (индексация с 0). Нужно найти **любой пик** — элемент, который **строго больше** своих соседей, и вернуть его индекс.

> Можно считать, что `nums[-1] = nums[n] = -∞`.  
> Алгоритм должен работать за **O(log n)**.

---

## 📌 Идея:
- Используем **бинарный поиск**
- Если `nums[mid] > nums[mid+1]` — пик в **левой** части (включая `mid`)
- Иначе — пик в **правой** части

---

## 📏 Структура:
- Два указателя: `left int`, `right int`
- Переменная `mid int`

---

## 🔁 Шаги алгоритма:

1. Инициализируем: `left := 0`, `right := len(nums) - 1`
2. Пока `left < right`:
   - `mid := (left + right) / 2`
   - Если `nums[mid] > nums[mid + 1]`, значит пик **слева**:
     - `right = mid`
   - Иначе пик **справа**:
     - `left = mid + 1`
3. Вернуть `left` (или `right`) — индекс пика

---

## ⏱️ Сложность:
- Время: O(log n)
- Память: O(1)

---

## 📄 Пример:

Вход:
```go
nums := []int{1, 2, 1, 3, 5, 6, 4}
```

Выход:
```go
5 // (или 1 — оба индекса корректны, т.к. оба соответствуют пикам)
```

---

## 📝 Решение:

```go
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid+1] {
			right = mid // пик слева
		} else {
			left = mid + 1 // пик справа
		}
	}

	return left
}
```

## 🏷 Теги:
- Array
- Binary Search

---