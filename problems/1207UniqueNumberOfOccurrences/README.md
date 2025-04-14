# Unique Number Of Occurrences

Level: Easy

[Ссылка на задачу](https://leetcode.com/problems/unique-number-of-occurrences/)

## 🧠 Задача:

Дан массив целых чисел `arr`. Нужно вернуть `true`, если **количество вхождений каждого числа уникально**, иначе — `false`.

---

## 📌 Идея:

1. Посчитать, сколько раз встречается каждое число (`map[int]int`)
2. Проверить, что **все количества уникальны**, т.е. не повторяются (используем `map[int]bool`)

---

## 📏 Структура:

- `freq map[int]int` — число → количество вхождений
- `seen map[int]bool` — было ли уже такое количество

---

## 🔁 Шаги алгоритма:

1. Инициализировать `freq` и `seen`
2. Пройти по массиву `arr`, посчитать частоты
3. Пройти по значениям `freq`:
   - если `seen[count] == true` → уже встречалось → `return false`
   - иначе `seen[count] = true`
4. Вернуть `true`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(n)

---

## 📄 Пример:

Вход:
```
arr = [1,2,2,1,1,3]
```

Промежуточные данные:
```
freq = {1:3, 2:2, 3:1}
seen = {3:true, 2:true, 1:true}
```

Выход:
```
true
```

---

## 📝 Решение:

```go
func uniqueOccurrences(arr []int) bool {
    freq := make(map[int]int)
    seen := make(map[int]bool)

    for _, num := range arr {
        freq[num]++
    }

    for _, count := range freq {
        if seen[count] {
            return false
        }
        seen[count] = true
    }

    return true
}
```

