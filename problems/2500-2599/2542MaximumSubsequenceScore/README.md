# Maximum Subsequence Score

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/maximum-subsequence-score/)

## 🧠 Задача:

> Даны два массива `nums1` и `nums2` одинаковой длины `n`, и целое число `k`.
> Нужно выбрать **подпоследовательность длины `k`** и вычислить её **оценку (score)** как:
> `score = (сумма выбранных nums1[i]) * (минимум из соответствующих nums2[i])`
> Вернуть максимальное возможное значение `score`.

---

## 📌 Идея:

* Сортируем пары `(nums2[i], nums1[i])` **по убыванию nums2\[i]**
* Проходим по ним, поддерживая сумму `k` наибольших `nums1[i]` в **min-heap**
* Для каждой итерации:

  * `nums2[i]` гарантированно — **текущее минимальное значение** среди `k`
  * Считаем `score = sum * nums2[i]`, обновляем максимум

---

## 📏 Структура:

* `pairs [][2]int` — пары `(nums2[i], nums1[i])`, отсортированные по `nums2`
* `heap MinHeap` — куча размера `k` для хранения наименьших `nums1[i]`
* `sum int64` — текущая сумма элементов в куче
* `result int64` — текущий максимум `score`

---

## 🔁 Шаги алгоритма:

1. Объединить `nums1` и `nums2` в пары `(nums2[i], nums1[i])`
2. Отсортировать по убыванию `nums2[i]`
3. Инициализировать min-heap и сумму
4. Для каждой пары:

   * Добавить `nums1[i]` в кучу и сумму
   * Если куча превышает `k` — удалить минимум и обновить сумму
   * Если длина кучи равна `k`, вычислить `score = sum * nums2[i]`, обновить результат
5. Вернуть `result`

---

## ⏱️ Сложность:

* Время: O(n log k)
* Память: O(k)

---

## 📄 Пример:

Вход:

```go
nums1 := []int{1,3,3,2}
nums2 := []int{2,1,3,4}
k := 3
```

Выход:

```go
12
```

Пояснение: выбираем индексы 0, 2, 3 → сумма = 1+3+2 = 6, минимум из nums2 = 2 → 6\*2 = 12

---

## 📝 Решение:

```go
import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums2[i], nums1[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)
	var sum int64 = 0
	var result int64 = 0

	for _, p := range pairs {
		heap.Push(h, p[1])
		sum += int64(p[1])

		if h.Len() > k {
			removed := heap.Pop(h).(int)
			sum -= int64(removed)
		}

		if h.Len() == k {
			score := sum * int64(p[0])
			if score > result {
				result = score
			}
		}
	}

	return result
}
```

---

## 🏷 Теги:
- Array
- Greedy
- Sorting
- Heap (Priority Queue)

---
