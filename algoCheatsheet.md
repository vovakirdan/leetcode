# 🧠 Алгоритмическая шпаргалка: памятка для уверенного решения задач

## ✨ 1. Предисловие

> Идеальное решение — не стартовая точка, а конечная цель. Начни просто. Реши в лоб. Оптимизируй потом.

Ты не обязан сразу знать O(1)-вариант или писать с двух указателей. Главное — **не застревать**. Эта шпаргалка — напоминание, с чего начать и куда копать.

---

## 📦 2. Часто встречающиеся кейсы и что с ними делать

### 🔍 2.1 Поиск

### Двоичный поиск (O(log n))

Применим, если:

- массив отсортирован
- нужно найти границу, максимум, условие

**Скелет:**

```go
left, right := 0, len(nums)-1
for left < right {
    mid := (left + right) / 2
    if условие {
        right = mid
    } else {
        left = mid + 1
    }
}
return left
```

---

### 📃 Sliding Window Cheatsheet

#### 🔹 Определение:
Sliding Window — это метод двух указателей `left` и `right`, которые движутся по массиву, определяя "окно" из элементов.

---

#### 🌐 Типовые шаблоны:

#### 1. Сумма/длина подмассива:
```go
left := 0
sum := 0
for right := 0; right < len(nums); right++ {
    sum += nums[right]
    for sum > target {
        sum -= nums[left]
        left++
    }
    // update max/min length/etc
}
```

#### 2. Макс/мин окно с не более K ошибок:
```go
left := 0
count := 0
for right := 0; right < len(nums); right++ {
    if nums[right] == 0 {
        count++
    }
    for count > 1 {
        if nums[left] == 0 {
            count--
        }
        left++
    }
    // update max window length: right - left
}
```

#### 3. Символы/счетчики в строке (map/rune)
```go
charCount := map[byte]int{}
left := 0
for right := 0; right < len(s); right++ {
    charCount[s[right]]++
    for invalid(charCount) {
        charCount[s[left]]--
        left++
    }
    // update result
}
```

---

## ✅ Ключевое:
- `right` всегда идёт вперёд
- `left` сдвигается, если нарушено условие
- Окно: `[left, right]`
- Часто нужно: `right - left + 1` (или `right - left`, если исключает что-то)

---

### 🌲 2.3 Деревья, графы: DFS и BFS

#### DFS (поиск в глубину, стек или рекурсия)

Хорошо подходит для:

- обхода дерева
- графов (рекурсивно)
- поиска путей, связности

**Пример (рекурсивный DFS на дереве):**

```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }
    dfs(node.Left)
    dfs(node.Right)
}
```

**DFS по графу с visited:**

```go
func dfs(v string, visited map[string]bool) {
    visited[v] = true
    for _, nei := range graph[v] {
        if !visited[nei] {
            dfs(nei, visited)
        }
    }
}
```

---

#### BFS (поиск в ширину, очередь)

Хорошо подходит для:

- поиска кратчайшего пути (в графе)
- задач "по уровням"

**Пример:**

```go
queue := []Node{start}
visited := map[Node]bool{start: true}

for len(queue) > 0 {
    size := len(queue)
    for i := 0; i < size; i++ {
        node := queue[0]
        queue = queue[1:]
        for _, nei := range graph[node] {
            if !visited[nei] {
                visited[nei] = true
                queue = append(queue, nei)
            }
        }
    }
}
```

---

### 🧺 2.4 Set / Map в Go

```go
set := map[int]struct{}{}
set[val] = struct{}{}
if _, ok := set[val]; ok { ... }
delete(set, val)
```

- `map[int]struct{}` — эффективнее, чем `map[int]bool`
- Используется для поиска, уникальности, сравнения

---

### 2.5 🔹 MinHeap (Мини-куча)

### 📌 Когда использовать:

- Нужно **отслеживать k наибольших / наименьших элементов**
- Часто используется в задачах:
  - `Kth Largest/Smallest Element`
  - `Top K frequent`
  - `Merge K sorted lists`
  - Потоковые данные (streaming) — поддерживать Top K на лету

### ✅ Типовая структура:

```go
import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}
```

### 🛠 Использование:

```go
h := &MinHeap{}
heap.Init(h)

heap.Push(h, 3)
heap.Push(h, 1)
heap.Push(h, 2)

min := heap.Pop(h).(int)  // → 1
```

### 📏 Сложность:
- `heap.Push`, `heap.Pop` — O(log n)
- `heap.Init` — O(n)

---

### 2.6 🔹 Quickselect

### 📌 Когда использовать:

- Нужно найти **k-й по величине (или по порядку) элемент**
- Альтернатива сортировке — **в среднем O(n)**
- Задачи:
  - `Kth Largest Element`
  - `Median of unsorted array`
  - `Top K frequent (варианты)`

### ✅ Скелет реализации:

```go
func quickSelect(nums []int, left, right, k int) int {
    if left == right {
        return nums[left]
    }

    pivotIndex := partition(nums, left, right)

    if pivotIndex == k {
        return nums[pivotIndex]
    } else if pivotIndex < k {
        return quickSelect(nums, pivotIndex+1, right, k)
    } else {
        return quickSelect(nums, left, pivotIndex-1, k)
    }
}

func partition(nums []int, left, right int) int {
    pivot := nums[right]
    i := left

    for j := left; j < right; j++ {
        if nums[j] < pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }

    nums[i], nums[right] = nums[right], nums[i]
    return i
}
```

### 🔍 Пример вызова:

```go
// Чтобы найти k-й по величине (отсортировать по убыванию):
index := len(nums) - k
res := quickSelect(nums, 0, len(nums)-1, index)
```

### 📏 Сложность:

- **Среднее время:** O(n)
- **Худшее время:** O(n²), если плохо выбрать пивот
- **Память:** O(1), если in-place

---

## ⚔️ MinHeap vs Quickselect — когда что:

| Сценарий                            | Лучше использовать |
|------------------------------------|---------------------|
| Найти **k-й элемент**              | Quickselect         |
| **Отслеживать k лучших** онлайн    | MinHeap (размер k)  |
| Обновляемые потоки / стрим         | MinHeap             |
| Множественные запросы по k         | Сортировка один раз |

---

### 💡 2.5 Часто используемые паттерны

- **Stack** — если надо обрабатывать "слева направо + удалять/отменять" (`remove stars`, `valid parentheses`, `monotonic stack`)
- **Prefix Sum / Diff Array** — если нужны быстрые диапазонные суммы
- **Dynamic Programming (DP)** — если подзадачи перекрываются
- **Backtracking** — если нужно попробовать все варианты (перестановки, sudoku)

---

## 🔚 3. Послесловие

Ты не обязан сразу помнить всё. Эта памятка — как ручной тормоз: тянешь, когда буксуешь.

Главное — **не бояться начать** и **идти итеративно**:

1. Понял задачу
2. Решил как умеешь
3. Оптимизировал при необходимости

> «Идеальный код — тот, который работает. А потом уже читается и ускоряется.»

