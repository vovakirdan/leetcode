# Greatest Common Divisor of Strings

Level: Easy

[Ссылка на задачу](https://leetcode.com/problems/greatest-common-divisor-of-strings/)

## 🧠 Задача:
> Для двух строк `s` и `t` говорят, что "`t` делит `s`", если `s = t + t + ... + t`.  
> Даны две строки `str1` и `str2`.  
> Нужно вернуть **наибольшую строку `x`**, которая делит **обе строки**.

---

## 📌 Идея:
- Если `str1 + str2 != str2 + str1` → нет общей строки-делителя
- Иначе — ответом будет `str1[:gcd(len(str1), len(str2))]`

---

## 📏 Структура:
- Функция `gcd(a, b int) int` — наибольший общий делитель длин
- Строка-результат `x`

---

## 🔁 Шаги алгоритма:

1. Проверить: если `str1 + str2 != str2 + str1` → вернуть `""`
2. Вычислить `g := gcd(len(str1), len(str2))`
3. Вернуть `str1[:g]`

---

## ⏱️ Сложность:
- Время: O(n + m), где n и m — длины строк
- Память: O(1)

---

## 📄 Пример:

Вход:
```go
str1 := "ABABAB"
str2 := "ABAB"
```

Выход:
```go
"AB"
```

---

## 📝 Решение:

```go
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	g := gcd(len(str1), len(str2))
	return str1[:g]
}
```

---

## 🏷 Теги:
- Math
- String

---
