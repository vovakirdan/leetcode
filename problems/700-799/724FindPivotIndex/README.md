# Find Pivot Index

Level: Easy

[Ссылка на задачу](https://leetcode.com/problems/find-pivot-index/)

## 🧠 Задача:

> Найти **первый индекс** в массиве `nums`, такой что **сумма всех чисел слева от него равна сумме всех чисел справа от него**.  
> Если индекс на краю массива — соответствующая сумма считается равной `0`.  
> Вернуть индекс, если он найден, иначе вернуть `-1`.

---

## 📌 Идея:

- Подсчитать **общую сумму** массива
- Использовать переменную `leftSum` для слежения за текущей суммой слева
- Вычислять сумму справа как `rightSum = total - leftSum - nums[i]`
- Если `leftSum == rightSum`, значит `i` — искомый индекс

---

## 📏 Структура:

- `total int` — сумма всех элементов массива
- `leftSum int` — сумма слева от текущего индекса

---

## 🔁 Шаги алгоритма:

1. Вычислить `total := sum(nums)`
2. Инициализировать `leftSum := 0`
3. Пройти по всем `i` от 0 до len(nums)-1:
   - `rightSum := total - leftSum - nums[i]`
   - Если `leftSum == rightSum` → вернуть `i`
   - Иначе `leftSum += nums[i]`
4. Если индекс не найден — вернуть `-1`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(1)

---

## 📄 Пример:

Вход:
```go
nums := []int{1, 7, 3, 6, 5, 6}
```

Выход:
```go
3
```

Пояснение:  
`1 + 7 + 3 = 11` слева от `6`,  
`5 + 6 = 11` справа от `6`.

---

## 📝 Решение:

```go
func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum := total - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
```

## 🏷 Теги:
- Array
- Prefix Sum

---