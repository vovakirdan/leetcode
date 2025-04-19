# Kth Largest Element in an Array

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/kth-largest-element-in-an-array/)

## 🧠 Задача:

> Дан массив `nums` и число `k`.  
> Нужно вернуть **k-й по величине** элемент в массиве (не обязательно уникальный).  
> То есть, после сортировки по убыванию это будет `nums[k-1]`.

> Задача должна быть решена **без сортировки всего массива**.

---

## 📌 Идея:

### Вариант 1: Min-Heap (приоритетная очередь)
- Поддерживаем кучу из `k` наибольших элементов
- Если в куче больше `k` элементов → удаляем наименьший
- На вершине останется `k`-й по величине элемент

### Вариант 2: QuickSelect (по принципу быстрой сортировки)
- Используем модифицированный QuickSort для поиска нужной позиции
- Разделяем массив вокруг опорного элемента
- Время в среднем — O(n), в худшем — O(n²)

---

## 📏 Структура:

### Для Min-Heap:
- Кастомная реализация `MinHeap` (через `container/heap`)

### Для QuickSelect:
- Рекурсивная функция `quickSelect(nums, left, right, k)`
- Функция `partition(...)` — разделение элементов относительно опорного

---

## 🔁 Шаги алгоритма (QuickSelect):

1. Преобразуем `k` в индекс: `len(nums) - k`
2. Вызываем `quickSelect` для поиска этого индекса
3. Базовый случай: если `left == right`, возвращаем значение
4. В `partition(...)` выбираем опорный и размещаем его на правильное место
5. Рекурсивно ищем нужную часть массива

---

## ⏱️ Сложность:

### Min-Heap:
- Время: O(n log k)
- Память: O(k)

### QuickSelect:
- Время: O(n) в среднем, O(n²) в худшем
- Память: O(1) — in-place

---

## 📄 Пример:

Вход:
```go
nums := []int{3, 2, 1, 5, 6, 4}
k := 2
```

Выход:
```go
5
```

---

## 📝 Решение 1 (Min-Heap):

```go
import (
    "container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    heap.Init(h)

    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h)
        }
    }

    return (*h)[0]
}
```

---

## 📝 Решение 2 (QuickSelect):

```go
func findKthLargest(nums []int, k int) int {
    target := len(nums) - k
    return quickSelect(nums, 0, len(nums)-1, target)
}

func quickSelect(nums []int, left, right, k int) int {
    if left == right {
        return nums[left]
    }

    pivotIndex := partition(nums, left, right)

    if pivotIndex == k {
        return nums[pivotIndex]
    } else if pivotIndex < k {
        return quickSelect(nums, pivotIndex+1, right, k)
    } else {
        return quickSelect(nums, left, pivotIndex-1, k)
    }
}

func partition(nums []int, left, right int) int {
    pivot := nums[right]
    i := left

    for j := left; j < right; j++ {
        if nums[j] < pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }

    nums[i], nums[right] = nums[right], nums[i]
    return i
}
```

---

## 🏷 Теги:
- Array
- Divide and Conquer
- Sorting
- Heap (Priority Queue)
- Quickselect

---
