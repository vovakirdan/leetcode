# Valid Palindrome II

Level: Easy

[Ссылка на задачу](https://leetcode.com/problems/valid-palindrome-ii/)

## 🧠 Задача:

> Дана строка `s`, вернуть `true`, если её можно сделать палиндромом, удалив **не более одного символа**.  
> Все символы — строчные латинские буквы.

---

## 📌 Идея:

- Используем два указателя: `left` и `right`
- При несовпадении символов допускаем удаление **либо левого, либо правого**
- Проверяем обе ситуации — если хотя бы одна даст палиндром, возвращаем `true`

---

## 📏 Структура:

- `left, right int` — границы проверки
- Вложенная функция `isPalin(i, j int) bool` — проверка, является ли подстрока палиндромом

---

## 🔁 Шаги алгоритма:

1. Инициализировать `left = 0`, `right = len(s) - 1`
2. Пока `left < right`:
   - Если `s[left] == s[right]` → двигаться к центру
   - Иначе:
     - Проверить: `isPalin(left+1, right)` **или** `isPalin(left, right-1)`
     - Если одно из них `true` → вернуть `true`, иначе `false`
3. Если вся строка палиндромна — вернуть `true`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(1)

---

## 📄 Пример:

Вход:
```go
s := "abca"
```

Выход:
```go
true // можно удалить 'c' → "aba"
```

---

## 📝 Решение:

```go
func validPalindrome(s string) bool {
	isPalin := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalin(left+1, right) || isPalin(left, right-1)
		}
		left++
		right--
	}

	return true
}
```

---

## 🏷 Теги:
- Two Pointers
- String
- Greedy

---
