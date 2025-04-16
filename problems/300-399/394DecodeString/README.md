# Decode String

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/decode-string/)

## 🧠 Задача:

> Дана закодированная строка `s`, нужно вернуть её **декодированную** версию.  
> Правило кодирования: `k[encoded_string]`, где `encoded_string` повторяется `k` раз.  
> Вложенность и повторения гарантированно корректны. Все цифры — только для повтора (`k`).

---

## 📌 Идея:

- Используем **два стека**:
  - Один для чисел (`numStack`)
  - Второй для промежуточных строк (`strStack`)
- По ходу прохода:
  - Цифры аккумулируются в `currNum`
  - Строки накапливаются в `currStr`
  - При `[` запоминаем контекст
  - При `]` — собираем строку и склеиваем

---

## 📏 Структура:

- `numStack []int` — стек чисел-повторов
- `strStack []string` — стек строк-контекстов
- `currNum int` — текущее число
- `currStr string` — текущая строящаяся строка

---

## 🔁 Шаги алгоритма:

1. Инициализировать `currNum = 0`, `currStr = ""`
2. Пройти по символам строки `s`:
   - Если цифра: обновить `currNum = currNum*10 + digit`
   - Если `[`: сохранить `currNum` и `currStr` в соответствующие стеки, сбросить
   - Если `]`: достать из стеков `n` и `prevStr`, собрать `currStr = prevStr + currStr*n`
   - Иначе — добавить символ к `currStr`
3. Вернуть `currStr`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(n)

---

## 📄 Пример:

Вход:
```go
s := "3[a2[c]]"
```

Выход:
```go
"accaccacc"
```

---

## 📝 Решение:

```go
func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	currNum := 0
	currStr := ""

	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '9':
			currNum = currNum*10 + int(ch-'0')
		case ch == '[':
			numStack = append(numStack, currNum)
			strStack = append(strStack, currStr)
			currNum = 0
			currStr = ""
		case ch == ']':
			n := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currStr = prevStr + strings.Repeat(currStr, n)
		default:
			currStr += string(ch)
		}
	}

	return currStr
}
```

---

## 🏷 Теги:
- String
- Stack
- Recursion

---
