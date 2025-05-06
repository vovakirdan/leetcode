# Odd Even Linked List

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/odd-even-linked-list/)

## 🧠 Задача:

> Дан односвязный список `head`.
> Нужно перегруппировать узлы так, чтобы сначала шли все **узлы с нечётными индексами**, затем **все с чётными** (индексация начинается с 1).
> При этом **относительный порядок в каждой группе должен сохраняться**.

---

## 📌 Идея:

* Используем два указателя: `odd` и `even`, а также `evenHead`, чтобы потом прикрепить чётные к хвосту нечётных
* Проходим список, пересоединяя ссылки:

  * `odd.Next` указывает на следующий нечётный
  * `even.Next` указывает на следующий чётный
* В конце соединяем `odd.Next = evenHead`

---

## 📏 Структура:

* `odd *ListNode` — текущий нечётный узел
* `even *ListNode` — текущий чётный узел
* `evenHead *ListNode` — начало списка чётных (для склейки в конце)

---

## 🔁 Шаги алгоритма:

1. Если `head == nil` или `head.Next == nil` → вернуть `head`
2. Инициализировать:

   * `odd = head`
   * `even = head.Next`
   * `evenHead = even`
3. Пока `even != nil && even.Next != nil`:

   * `odd.Next = even.Next`
   * `odd = odd.Next`
   * `even.Next = odd.Next`
   * `even = even.Next`
4. `odd.Next = evenHead`
5. Вернуть `head`

---

## ⏱️ Сложность:

* Время: O(n)
* Память: O(1)

---

## 📄 Пример:

Вход:

```go
head = [1,2,3,4,5]
```

Выход:

```go
[1,3,5,2,4]
```

---

## 📝 Решение:

```go
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
```

---

## 🏷 Теги:
- Linked List

---
