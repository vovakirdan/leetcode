package evaluatedivision

type Edge struct {
	to     string
	weight float64
}

func bfs(graph map[string][]Edge, start, end string) float64 {
	visited := make(map[string]bool)
	queue := []Edge{{start, 1.0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.to == end {
			return current.weight
		}

		visited[current.to] = true

		for _, neighbor := range graph[current.to] {
			if !visited[neighbor.to] {
				queue = append(queue, Edge{neighbor.to, current.weight * neighbor.weight})
			}
		}
	}

	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Edge) // построение графа

	// заполнение графа
	for i, eq := range equations {
		graph[eq[0]] = append(graph[eq[0]], Edge{to: eq[1], weight: values[i]})
		graph[eq[1]] = append(graph[eq[1]], Edge{to: eq[0], weight: 1 / values[i]})
	}

	// обработка запросов
	results := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := graph[query[0]]; !ok {
			results[i] = -1.0
			continue
		}
		if _, ok := graph[query[1]]; !ok {
			results[i] = -1.0
			continue
		}

		results[i] = bfs(graph, query[0], query[1])
	}

	return results
}
