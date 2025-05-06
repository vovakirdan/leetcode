# Evaluate Division

Level: Medium

[Ссылка на задачу](https://leetcode.com/problems/evaluate-division/)

## 🧠 Задача:

> Даны уравнения вида `A / B = value` и массив запросов `C / D = ?`.
> Для каждого запроса нужно найти результат деления, либо вернуть `-1.0`, если вычислить невозможно.

---

## 📌 Идея:

* Построим **ориентированный граф**, где переменные — это узлы, а `A / B = x` означает:

  * ребро `A → B` с весом `x`
  * ребро `B → A` с весом `1/x`
* Для ответа на запрос запускаем **поиск в ширину (BFS)** от `C` к `D` с накоплением произведения весов

---

## 📏 Структура:

* `map[string][]Edge` — граф: `Edge` хранит `to` (куда) и `weight` (коэффициент)
* `Edge` — структура ребра
* `bfs(...)` — поиск пути и накопление произведения весов

---

## 🔁 Шаги алгоритма:

1. Построить граф:

   * Для каждого `A / B = value`:

     * добавить `A → B (value)` и `B → A (1 / value)`
2. Для каждого запроса:

   * Если хотя бы одна переменная отсутствует — вернуть `-1.0`
   * Иначе запустить `bfs` от `C` к `D` и записать результат

---

## ⏱️ Сложность:

* Время:

  * Построение графа: O(e)
  * Обработка запросов: O(q \* (v + e)) в худшем случае
* Память: O(v + e)

Где:

* `e` — количество уравнений
* `v` — количество уникальных переменных
* `q` — количество запросов

---

## 📄 Пример:

Вход:

```go
equations := [][]string{{"a", "b"}, {"b", "c"}}
values := []float64{2.0, 3.0}
queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
```

Выход:

```go
[]float64{6.0, 0.5, -1.0, 1.0, -1.0}
```

---

## 📝 Решение:

```go
type Edge struct {
	to     string
	weight float64
}

func bfs(graph map[string][]Edge, start, end string) float64 {
	visited := make(map[string]bool)
	queue := []Edge{{to: start, weight: 1.0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.to == end {
			return current.weight
		}

		visited[current.to] = true

		for _, neighbor := range graph[current.to] {
			if !visited[neighbor.to] {
				queue = append(queue, Edge{
					to:     neighbor.to,
					weight: current.weight * neighbor.weight,
				})
			}
		}
	}

	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Edge)

	for i, eq := range equations {
		a, b := eq[0], eq[1]
		graph[a] = append(graph[a], Edge{to: b, weight: values[i]})
		graph[b] = append(graph[b], Edge{to: a, weight: 1 / values[i]})
	}

	results := make([]float64, len(queries))
	for i, query := range queries {
		a, b := query[0], query[1]
		if _, ok := graph[a]; !ok || _, ok := graph[b]; !ok {
			results[i] = -1.0
		} else {
			results[i] = bfs(graph, a, b)
		}
	}

	return results
}
```

---

## 🏷 Теги:
- Array
- String
- Depth-First Search
- Breadth-First Search
- Union Find
- Graph
- Shortest Path

---
