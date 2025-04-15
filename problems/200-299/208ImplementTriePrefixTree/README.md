# Implement Trie (Prefix Tree)

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/implement-trie-prefix-tree)

## 🧠 Задача:
> Реализовать структуру данных **Trie (префиксное дерево)** с поддержкой следующих операций:
> - `Insert(word string)` — вставка слова
> - `Search(word string) bool` — проверка наличия слова
> - `StartsWith(prefix string) bool` — проверка наличия слов с данным префиксом

---

## 📌 Идея:
- Используем дерево, где каждый узел — это символ слова
- Каждый узел содержит `map[rune]*Trie` для переходов к дочерним узлам
- Флаг `IsTerminal` показывает, заканчивается ли слово на этом узле

---

## 📏 Структура:
- Узел: `type Trie struct`
  - `Children map[rune]*Trie`
  - `IsTerminal bool`
- Методы:
  - `Constructor() Trie`
  - `Insert(word string)`
  - `Search(word string) bool`
  - `StartsWith(prefix string) bool`

---

## 🔁 Шаги алгоритма:

### Insert:
1. Стартуем с корня дерева
2. Для каждого символа:
   - Если его нет в `Children`, добавляем новый узел
   - Переходим в следующий узел
3. По завершению — отмечаем `IsTerminal = true`

### Search:
1. Стартуем с корня
2. Для каждого символа:
   - Если его нет в `Children`, возвращаем `false`
   - Иначе переходим к следующему узлу
3. Возвращаем `true`, если `IsTerminal == true`

### StartsWith:
1. Как и `Search`, но не проверяем `IsTerminal` в конце
2. Если все символы присутствуют — префикс существует

---

## ⏱️ Сложность:
- Время:
  - `Insert`, `Search`, `StartsWith`: O(m), где m — длина слова или префикса
- Память: O(n·m), где n — количество слов, m — средняя длина слова

---

## 📄 Пример:

```go
trie := Constructor()
trie.Insert("apple")
trie.Search("apple")   // true
trie.Search("app")     // false
trie.StartsWith("app") // true
trie.Insert("app")
trie.Search("app")     // true
```

---

## 📝 Решение:

```go
package trieprefixtree

// Trie represents a trie data structure
type Trie struct {
	Children   map[rune]*Trie
	IsTerminal bool
}

// Constructor creates a new Trie
func Constructor() Trie {
	return Trie{
		Children: make(map[rune]*Trie),
	}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if node.Children[c] == nil {
			node.Children[c] = &Trie{
				Children: make(map[rune]*Trie),
			}
		}
		node = node.Children[c]
	}
	node.IsTerminal = true
}

// Search returns true if the word is in the trie
func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		exists, ok := node.Children[c]
		if !ok {
			return false
		}
		node = exists
	}
	return node.IsTerminal
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		exists, ok := node.Children[c]
		if !ok {
			return false
		}
		node = exists
	}
	return true
}
```

---

## 🏷 Теги:
- Hash Table
- String
- Design
- Trie

---