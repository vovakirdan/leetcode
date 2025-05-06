# Guess Number Higher or Lower

Level: Easy

[Ссылка на задачу](https://leetcode.com/problems/guess-number-higher-or-lower/)

## 🧠 Задача:

> Угадай число, задуманное противником в диапазоне от `1` до `n`.
> Ты можешь вызывать API `guess(num int) int`, которое возвращает:
>
> * `0` если `num == pick`
> * `-1` если `num > pick`
> * `1` если `num < pick`
>
> Нужно вернуть число `pick`.

---

## 📌 Идея:

* Используем **бинарный поиск** от `1` до `n`, сравнивая середину с `pick` через `guess()`
* Сдвигаем границы поиска в зависимости от результата

---

## 📏 Структура:

* Указатели: `left`, `right`, `mid`
* Вызов `guess(mid)` — основное условие

---

## 🔁 Шаги алгоритма:

1. Инициализировать `left = 1`, `right = n`
2. Пока `left <= right`:

   * Вычислить `mid := left + (right - left) / 2`
   * Вызвать `guess(mid)`

     * Если `0` → вернуть `mid`
     * Если `-1` → число меньше → `right = mid - 1`
     * Если `1` → число больше → `left = mid + 1`
3. Возврат `-1` невозможен по условию

---

## ⏱️ Сложность:

* Время: O(log n)
* Память: O(1)

---

## 📄 Пример:

Вход:

```go
n := 10 // pick = 6
```

Выход:

```go
6
```

---

## 📝 Решение:

```go
func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1 // недостижимый случай
}
```

---

## 🏷 Теги:
- Binary Search
- Interactive

---
