package rottingoranges

func orangesRotting(grid [][]int) int {
	queue := [][]int{} // очередь позиций гнилых апельсинов
	fresh := 0         // счётчик свежих апельсинов

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j}) // начальная гниль
			} else if grid[i][j] == 1 {
				fresh++ // считаем все свежие
			}
		}
	}

	dirs := [][]int{
		{0, 1},  // вправо
		{1, 0},  // вниз
		{0, -1}, // влево
		{-1, 0}, // вверх
	}

	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:]

			for _, dir := range dirs {
				ni := cell[0] + dir[0]
				nj := cell[1] + dir[1]

				if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					queue = append(queue, []int{ni, nj})
					fresh--
				}
			}
		}
		minutes++
	}

	if fresh != 0 {
		return -1
	}
	return minutes
}
