# Longest Subarray of 1's After Deleting One Element

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/)

## 🧠 Задача:

> Дан бинарный массив `nums`. Нужно **удалить один элемент** (обязательно).  
> Вернуть длину самой длинной непустой подстроки, содержащей **только единицы**, после удаления одного элемента.  
> Если такой подстроки не существует — вернуть `0`.

---

## 📌 Идея:

- Используем **двухуказательную технику** (скользящее окно)
- Допускаем **не более одного нуля** в текущем окне
- Размер окна `right - left`, но **учитываем, что один элемент должен быть удалён**

---

## 📏 Структура:

- `left, right int` — границы окна
- `zeroCount int` — количество нулей в текущем окне
- `maxLen int` — результат

---

## 🔁 Шаги алгоритма:

1. Инициализируем `left = 0`, `zeroCount = 0`, `maxLen = 0`
2. Проходим `right` от 0 до конца:
   - Если `nums[right] == 0`, увеличиваем `zeroCount`
   - Если `zeroCount > 1` — сдвигаем `left` и уменьшаем `zeroCount`, если слева был 0
3. Обновляем `maxLen = max(maxLen, right - left)`
4. Возвращаем `maxLen`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(1)

---

## 📄 Пример:

Вход:
```go
nums := []int{1, 1, 0, 1}
```

Выход:
```go
3 // удаляем 0, получаем [1,1,1]
```

---

## 📝 Решение:

```go
func longestSubarray(nums []int) int {
	left := 0
	zeroCount := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxLen = max(maxLen, right-left)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

---

## 🏷 Теги:
- Array
- Dynamic Programming
- Sliding Window

---
