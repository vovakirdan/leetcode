# Kids With the Greatest Number of Candies

Level: Easy

[Ссылка на задачу](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/)

## 🧠 Задача:

> Дан массив `candies`, где `candies[i]` — количество конфет у i-го ребёнка.  
> Также дано число `extraCandies` — дополнительное количество конфет.  
> Нужно вернуть булев массив, где `result[i] == true`, если, дав i-му ребёнку все дополнительные конфеты, он будет иметь **максимальное** количество конфет среди всех детей.

---

## 📌 Идея:

- Найти максимальное число конфет среди всех детей
- Проверить для каждого ребёнка: станет ли он лидером, если получит `extraCandies`

---

## 📏 Структура:

- `maxNum int` — максимум по массиву `candies`
- `res []bool` — результирующий булев массив

---

## 🔁 Шаги алгоритма:

1. Найти `maxNum := max(candies)`
2. Инициализировать `res := make([]bool, len(candies))`
3. Для каждого i:
   - Если `candies[i] + extraCandies >= maxNum` → `res[i] = true`
   - Иначе → `res[i] = false`
4. Вернуть `res`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(n)

---

## 📄 Пример:

Вход:
```go
candies := []int{2, 3, 5, 1, 3}
extraCandies := 3
```

Выход:
```go
[true, true, true, false, true]
```

---

## 📝 Решение:

```go
import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxNum := slices.Max(candies)
	res := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		res[i] = candies[i]+extraCandies >= maxNum
	}

	return res
}
```

---

## 🏷 Теги:
- Array

---
