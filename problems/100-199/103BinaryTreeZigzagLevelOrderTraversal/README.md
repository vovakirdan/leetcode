Вот `README.md` для задачи **103. Binary Tree Zigzag Level Order Traversal**, оформленный по твоему шаблону:

---

# Binary Tree Zigzag Level Order Traversal

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)

## 🧠 Задача:

> Дано бинарное дерево `root`.  
> Вернуть **зигзагообразный обход по уровням** — то есть:
> - первый уровень — слева направо
> - второй уровень — справа налево
> - третий — снова слева направо и т.д.

---

## 📌 Идея:

- Классический **обход в ширину (BFS)** с очередью
- Сохраняем текущий уровень и чередуем порядок:
  - если уровень чётный → слева направо
  - если нечётный → справа налево

---

## 📏 Структура:

- Очередь: `queue []*TreeNode`
- Результат: `res [][]int`
- Счётчик уровня: `level int`

---

## 🔁 Шаги алгоритма:

1. Если `root == nil` → вернуть `[]`
2. Инициализировать `queue := []*TreeNode{root}`, `level := 0`
3. Пока очередь не пуста:
   - `size := len(queue)`
   - Создать слайс `row := make([]int, size)`
   - Пройти `size` элементов:
     - Извлечь `node := queue[0]`
     - Добавить значение `node.Val`:
       - в конец `row[i]`, если чётный уровень
       - в начало `row[size-1-i]`, если нечётный
     - Добавить в очередь `node.Left` и `node.Right` при наличии
   - Добавить `row` в `res`
   - Увеличить `level++`
4. Вернуть `res`

---

## ⏱️ Сложность:

- Время: O(n)
- Память: O(n)

---

## 📄 Пример:

Вход:
```go
root = [3,9,20,null,null,15,7]
```

Выход:
```go
[[3], [20, 9], [15, 7]]
```

---

## 📝 Решение:

```go
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		row := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				row[i] = node.Val
			} else {
				row[size-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, row)
		level++
	}

	return res
}
```

---

## 🏷 Теги:
- Tree
- Breadth-First Search
- Binary Tree

---
