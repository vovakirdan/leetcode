# Max Number of K-Sum Pairs

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/max-number-of-k-sum-pairs/)

## 🧠 Задача:

> Дан массив `nums` и целое число `k`.
> За одну операцию можно выбрать **две разные позиции** в массиве, сумма которых равна `k`, и удалить эти два числа.
> Вернуть максимальное количество таких операций.

---

## 📌 Идея:

* Сортируем массив
* Используем два указателя:

  * `left` — с начала
  * `right` — с конца
* Двигаем указатели, подбирая пары с суммой `k`

---

## 📏 Структура:

* `left, right int` — границы массива
* `count int` — счётчик успешных операций

---

## 🔁 Шаги алгоритма:

1. Отсортировать `nums`
2. Инициализировать `left = 0`, `right = len(nums) - 1`
3. Пока `left < right`:

   * Вычислить `sum = nums[left] + nums[right]`
   * Если `sum == k` → увеличить `count`, сдвинуть оба указателя
   * Если `sum < k` → сдвинуть `left`
   * Если `sum > k` → сдвинуть `right`
4. Вернуть `count`

---

## ⏱️ Сложность:

* Время: O(n log n) (на сортировку)
* Память: O(1)

---

## 📄 Пример:

Вход:

```go
nums := []int{1, 2, 3, 4}
k := 5
```

Выход:

```go
2 // пары: (1,4) и (2,3)
```

---

## 📝 Решение:

```go
import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	count := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			count++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}
	return count
}
```

---

## 💡 Альтернатива:

> Можно также решить задачу с помощью `map[int]int` (счётчик чисел), что даст:

* Время: O(n)
* Память: O(n)
* Алгоритм: для каждого `num`, проверять есть ли `k - num` в map

---

## 🏷 Теги:
- Array
- Hash Table
- Two Pointers
- Sorting

---
