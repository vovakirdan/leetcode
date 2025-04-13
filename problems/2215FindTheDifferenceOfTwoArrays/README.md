# Find The Difference Of Two Arrays

Level: Easy

[Ссылка на задачу](https://leetcode.com/problems/find-the-difference-of-two-arrays/)

## 🧠 Задача:

Даны два массива целых чисел nums1 и nums2.
Вернуть список из двух списков:

- answer[0] — все уникальные числа из nums1, которых нет в nums2
- answer[1] — все уникальные числа из nums2, которых нет в nums1

Порядок чисел в каждом подсписке — произвольный.

---

## 📌 Идея:

- Используем множества (map[int]bool) для хранения уникальных элементов
- Проходим по каждому множеству и сравниваем с другим

---

## 📏 Структура:

- set1 map[int]bool — уникальные числа из nums1
- set2 map[int]bool — уникальные числа из nums2
- res1 []int, res2 []int — результат

---

## 🔁 Шаги алгоритма:

1. Создать set1 и set2, заполнить числами из nums1 и nums2
2. Пройти по set1:
   - Если число отсутствует в set2 → добавить в res1
3. Пройти по set2:
   - Если число отсутствует в set1 → добавить в res2
4. Вернуть [][]int{res1, res2}

---

## ⏱️ Сложность:

- Время: O(n + m), где n = len(nums1), m = len(nums2)
- Память: O(n + m)

## 📄 Пример:

Вход:

```go
nums1 := []int{1, 2, 3}
nums2 := []int{2, 4, 6}
```

Выход:

```go
[][]int{{1, 3}, {4, 6}}
```

---

## 📝 Решение:

```go
func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]bool{}
	set2 := map[int]bool{}

	for _, elem := range nums1 {
		set1[elem] = true
	}

	for _, elem := range nums2 {
		set2[elem] = true
	}

	res1 := []int{}
	res2 := []int{}

	for num := range set1 {
		if !set2[num] {
			res1 = append(res1, num)
		}
	}

	for num := range set2 {
		if !set1[num] {
			res2 = append(res2, num)
		}
	}

	return [][]int{res1, res2}
}
```
