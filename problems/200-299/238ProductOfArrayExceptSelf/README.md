# Product of Array Except Self

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/product-of-array-except-self/)

## 🧠 Задача:

> Дан массив целых чисел `nums`, вернуть массив `answer`, где `answer[i]` — это произведение всех элементов `nums`, **кроме `nums[i]`**.  
> **Нельзя использовать деление**, и алгоритм должен работать за **O(n)**.

---

## 📌 Идея:

- Сначала проходим массив слева направо, считая **префиксное произведение** до каждого элемента
- Затем проходим справа налево, умножая на **постфиксное произведение** справа от каждого элемента
- Используем только один массив `answer` и переменную `right` (O(1) доп. памяти, не считая результата)

---

## 📏 Структура:

- `answer []int` — результат
- `right int` — постфиксное произведение справа от текущего индекса

---

## 🔁 Шаги алгоритма:

1. Инициализировать `answer[0] = 1`
2. Пройти слева направо:
   - `answer[i] = answer[i-1] * nums[i-1]`
3. Инициализировать `right = 1`
4. Пройти справа налево:
   - `answer[i] *= right`
   - `right *= nums[i]`
5. Вернуть `answer`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(1) дополнительной (если не считать `answer` как доп. память)

---

## 📄 Пример:

Вход:
```go
nums := []int{1, 2, 3, 4}
```

Выход:
```go
[24, 12, 8, 6]
```

---

## 📝 Решение:

```go
func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}
```

---

## 🏷 Теги:
- Array
- Prefix Sum

---