# Delete Node in a BST

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/delete-node-in-a-bst/)

## 🧠 Задача:

> Дано корневое значение `root` бинарного дерева поиска (BST) и целое число `key`.  
> Удалить узел со значением `key`, если он существует, и вернуть **обновлённый корень** дерева.

> Удаление делится на два этапа:
> - Найти узел с `key`
> - Удалить его корректно, сохраняя свойства BST

---

## 📌 Идея:

- Используем **рекурсивный обход** дерева, сравнивая `key` с текущим значением
- При удалении:
  - Если у узла **0 или 1 потомок** — просто заменяем
  - Если **2 потомка** — заменяем значение на **минимальный элемент из правого поддерева**, затем удаляем этот элемент

---

## 📏 Структура:

- Функция `deleteNode(root *TreeNode, key int) *TreeNode`
- Вспомогательная функция `findMin(node *TreeNode) *TreeNode` — находит самый левый узел (наименьший)

---

## 🔁 Шаги алгоритма:

1. Если `root == nil` → вернуть `nil`
2. Если `key < root.Val` → идём влево
3. Если `key > root.Val` → идём вправо
4. Если `key == root.Val`:
   - Если `Left == nil` → вернуть `Right`
   - Если `Right == nil` → вернуть `Left`
   - Иначе:
     - Найти минимальный узел в `Right`
     - Присвоить его значение текущему
     - Рекурсивно удалить этот минимальный узел в `Right`
5. Вернуть `root`

---

## ⏱️ Сложность:

- Время: O(h), где h — высота дерева (в среднем O(log n), в худшем O(n))
- Память: O(h) стек вызова

---

## 📄 Пример:

Вход:
```go
root = [5,3,6,2,4,null,7], key = 3
```

Выход:
```go
[5,4,6,2,null,null,7] // или другой корректный вариант
```

---

## 📝 Решение:

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// Нашли удаляемый узел
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// У узла два потомка — находим минимум справа
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}

	return root
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
```

---

## 🏷 Теги:
- Tree
- Binary Search Tree
- Binary Tree

---
