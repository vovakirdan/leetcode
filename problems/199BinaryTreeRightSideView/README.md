# Binary Tree Right Side View

Level: Medium

## 🧠 Задача:
> Дан корень бинарного дерева. Нужно вернуть значения узлов, которые видны при смотре на дерево справа (по уровням).

---

## 📌 Идея:
- Обход в ширину (**BFS**) по уровням
- На каждом уровне сохраняем последний узел

---

## 📏 Структура:
- Очередь: `queue []*TreeNode`
- Результат: `result []int`

---

## 🔁 Шаги алгоритма:

1. Если `root == nil`, вернуть `[]`
2. Инициализировать: `queue := []*TreeNode{root}`
3. Пока `queue` не пуста:
   1. `levelSize := len(queue)`
   2. Цикл `i := 0 с проходом всех узлов уровня`
      - `node := queue[0]`, `queue = queue[1:]`
      - Если `i == levelSize - 1` → `result = append(result, node.Val)`
      - Добавить `node.Left` и `node.Right`, если не `nil`
4. Вернуть `result`

---

## ⏱️ Сложность:
- Время: O(n)
- Память: O(n)

---

## 📄 Пример:

Вход:
```
    1
   / \
  2   3
   \   \
    5   4
```

Выход: `[1, 3, 4]`

---

## 📝 Решение:

```go
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	queue := []*TreeNode{root}
	result := []int{}

	for len(queue) > 0 {
        level_size := len(queue)

        for i := 0; i < level_size; i++ {
            node := queue[0]
            queue = queue[1:]

            if i == level_size - 1 {
                result = append(result, node.Val)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    return result
}
```