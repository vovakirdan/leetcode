# Лайвкодинг кейсы
Здесь будут разобраны часто встречающиеся кейсы на собеседованиях.

## Навигация

1. [LRU Cache](#кейс-1-lru-cache) | [Код](#реализация-кейса-1-lru-cache)
2. [Конвейер (Pipeline)](#кейс-2-конвейер-pipeline) | [Код](#реализация-кейса-2-конвейер-pipeline)
3. [Мини-вебсервер](#кейс-3-мини-вебсервер) | [Код](#реализация-кейса-3-мини-вебсервер)
4. [Собственный JSON парсер](#кейс-4-собственный-json-парсер) | [Код](#реализация-кейса-4-собственный-json-парсер)
5. [Детект зацикливания в графе](#кейс-5-детект-зацикливания-в-графе) | [Код](#реализация-кейса-5-детект-зацикливания-в-графе)
6. [EventBus](#кейс-6-eventbus) | [Код](#реализация-кейса-6-eventbus)
7. [Кастомный мьютекс](#кейс-7-кастомный-мьютекс) | [Код](#реализация-кейса-7-кастомный-мьютекс)
8. [Мини-SQL интерпретатор](#кейс-8-мини-sql-интерпретатор) | [Код](#реализация-кейса-8-мини-sql-интерпретатор)
9. [Таймаут и отмена операций](#кейс-9-таймаут-и-отмена-операций) | [Код](#реализация-кейса-9-таймаут-и-отмена-операций)
10. [Трэдсейфовый счётчик](#кейс-10-трэдсейфовый-счётчик) | [Код](#реализация-кейса-10-трэдсейфовый-счётчик)

## Кейсы

### **Кейс 1: LRU Cache**
**Задача:**  
Реализовать структуру `LRUCache` с методами `Get(key int) int` и `Put(key, value int)`.

[**Код**](#реализация-кейса-1-lru-cache)

**Вопросы:**  
- Как обеспечить O(1) время на доступ?
- Какие структуры данных использовать?

**Разбор:**  
Хеш-таблица + двусвязный список. `map[int]*Node` и двусвязный список для хранения порядка использования. Удаление и вставка в список — за O(1).

---

### **Кейс 2: Конвейер (Pipeline)**
**Задача:**  
Сделать три стадии обработки чисел (умножить на 2 → возвести в квадрат → перевести в строку), используя каналы и горутины.

[**Код**](#реализация-кейса-2-конвейер-pipeline)

**Вопросы:**  
- Как избежать утечек горутин?
- Как организовать закрытие каналов?

**Разбор:**  
Каждая стадия читает из входного канала и пишет в выходной. Использовать `defer close(out)` внутри горутин, чтобы закрыть канал корректно.

---

### **Кейс 3: Мини-вебсервер**
**Задача:**  
Реализовать HTTP-сервер, который считает количество запросов от каждого IP и возвращает этот счётчик.

[**Код**](#реализация-кейса-3-мини-вебсервер)

**Вопросы:**  
- Как обеспечить конкурентный доступ к мапе?
- Какие структуры лучше использовать: `sync.Map`, `map+Mutex`, `atomic`?

**Разбор:**  
Для конкурентного доступа — `sync.Map` или `map[string]int` + `sync.Mutex`. Сложность в том, что IP нужно вычленить из `http.Request`.

---

### **Кейс 4: Собственный JSON парсер**
**Задача:**  
Реализовать парсинг JSON-объекта вида `{"a":1,"b":2}` в `map[string]int`.

[**Код**](#реализация-кейса-4-собственный-json-парсер)

**Вопросы:**  
- Почему `encoding/json` не подходит?
- Как организовать state machine для парсинга?

**Разбор:**  
Механика: state machine по символам строки. Тест на баланс скобок, экранирование, и правильную обработку ключей/значений.

---

### **Кейс 5: Детект зацикливания в графе**
**Задача:**  
Дан ориентированный граф. Проверить, содержит ли он цикл.

[**Код**](#реализация-кейса-5-детект-зацикливания-в-графе)

**Вопросы:**  
- Почему нужен visited + stackVisited?
- Как реализовать DFS без рекурсии?

**Разбор:**  
Классический DFS с двумя множествами: `visited` и `onStack`. Можно реализовать с использованием стека вручную (итеративный вариант).

---

### **Кейс 6: EventBus**
**Задача:**  
Реализовать `EventBus`, который позволяет подписчикам подписываться на события и получать уведомления.

[**Код**](#реализация-кейса-6-eventbus)

**Вопросы:**  
- Какую структуру использовать для хранения подписчиков?
- Как обрабатывать конкурентную отправку сообщений?

**Разбор:**  
Использовать `map[string][]chan any`. Для потокобезопасности — `sync.RWMutex`. Можно реализовать буферизированные каналы.

---

### **Кейс 7: Кастомный мьютекс**
**Задача:**  
Реализовать структуру, похожую на `sync.Mutex`, на основе каналов.

[**Код**](#реализация-кейса-7-кастомный-мьютекс)

**Вопросы:**  
- Как использовать каналы как блокировку?
- Какие edge-cases могут быть при использовании?

**Разбор:**  
`make(chan struct{}, 1)` — канал как семафор. `Lock` — это `lock <- struct{}{}`, `Unlock` — `<-lock`. Важно не делать double-lock.

---

### **Кейс 8: Мини-SQL интерпретатор**
**Задача:**  
Реализовать парсер для простых SQL-запросов: `SELECT name FROM users WHERE age > 20`.

[**Код**](#реализация-кейса-8-мини-sql-интерпретатор)

**Вопросы:**  
- Как разбивать строку на токены?
- Что проще — парсить вручную или использовать генератор?

**Разбор:**  
Сначала токенизация (`strings.Fields` и т.д.), потом state machine по ключевым словам. Хранить AST в виде структуры.

---

### **Кейс 9: Таймаут и отмена операций**
**Задача:**  
Реализовать функцию, которая делает запрос к серверу с таймаутом 2 секунды. Если нет ответа — отменяет.

[**Код**](#реализация-кейса-9-таймаут-и-отмена-операций)

**Вопросы:**  
- Как работает `context.WithTimeout`?
- Что произойдёт, если не отменить контекст?

**Разбор:**  
Использовать `ctx, cancel := context.WithTimeout(...)`. В defer — `cancel()`. Контекст передаётся в `http.NewRequestWithContext`.

---

### **Кейс 10: Трэдсейфовый счётчик**
**Задача:**  
Реализовать безопасный счётчик с инкрементом и чтением значений.

[**Код**](#реализация-кейса-10-трэдсейфовый-счётчик)

**Вопросы:**  
- Чем `atomic` лучше `Mutex`?
- Когда использовать `int32`, а когда `int64`?

**Разбор:**  
`atomic.AddInt64` / `atomic.LoadInt64` — эффективный и потокобезопасный способ. Лучше для частых инкрементов.

## Задачи

## **Реализация кейса 1: LRU Cache**

### **Задача:**
Реализовать структуру данных LRU Cache с методами:

```go
Get(key int) int     // возвращает значение по ключу или -1, если ключа нет
Put(key int, value int) // вставляет значение, при превышении размера удаляет наименее использованное
```

Емкость задаётся при создании. Все операции должны быть за O(1).

---

### **Код:**
```go
package main

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, found := c.cache[key]; found {
		c.moveToFront(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, found := c.cache[key]; found {
		node.value = value
		c.moveToFront(node)
		return
	}

	if len(c.cache) >= c.capacity {
		c.removeOldest()
	}

	node := &Node{key: key, value: value}
	c.cache[key] = node
	c.addToFront(node)
}

func (c *LRUCache) moveToFront(node *Node) {
	c.remove(node)
	c.addToFront(node)
}

func (c *LRUCache) addToFront(node *Node) {
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) removeOldest() {
	oldest := c.tail.prev
	c.remove(oldest)
	delete(c.cache, oldest.key)
}

// Пример использования
func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)           // удаляется ключ 2
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)           // удаляется ключ 1
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}
```

---

### **Объяснение:**

- Используем `map[int]*Node` — быстрое O(1) обращение по ключу.
- Двусвязный список (`head <-> ... <-> tail`) — чтобы быстро перемещать элементы и удалять самый старый.
- `moveToFront` делает элемент самым «свежим» при `Get` или `Put`.

---

### **Исправления (если бы были ошибки):**

Если забыть `delete(c.cache, oldest.key)`, то кеш "растёт", и память не освобождается.

Также возможна ошибка, если не инициализировать `head` и `tail` как связанные узлы: тогда `addToFront` или `remove` паникнут на `nil`.

---

## **Реализация кейса 2: Конвейер (Pipeline)**

### **Задача:**
Реализовать конвейер из трёх стадий обработки чисел:
1. Умножение на 2
2. Возведение в квадрат
3. Преобразование в строку

Каждая стадия должна быть отдельной горутиной, использующей каналы для передачи данных. Все данные должны быть обработаны корректно и без утечек горутин.

---

### **Код:**
```go
package main

import (
	"fmt"
	"strconv"
)

// Первая стадия: умножение на 2
func multiplyByTwo(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for x := range in {
			out <- x * 2
		}
	}()
	return out
}

// Вторая стадия: возведение в квадрат
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for x := range in {
			out <- x * x
		}
	}()
	return out
}

// Третья стадия: преобразование в строку
func toString(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for x := range in {
			out <- strconv.Itoa(x)
		}
	}()
	return out
}

func main() {
	// Исходные данные
	input := make(chan int)
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()

	// Построение пайплайна
	stage1 := multiplyByTwo(input)
	stage2 := square(stage1)
	output := toString(stage2)

	// Вывод результата
	for s := range output {
		fmt.Println(s)
	}
}
```

---

### **Объяснение:**

- Канал `input` генерирует числа от 1 до 5.
- Каждая стадия реализована как функция, принимающая канал на чтение и возвращающая канал на запись.
- Внутри каждой стадии запускается горутина, которая читает из входного канала, обрабатывает значение и отправляет в выходной канал.
- `defer close(out)` гарантирует, что канал будет закрыт после завершения обработки — это важно, чтобы **не было зависания в следующей стадии**.
- В `main()` — мы просто читаем из последнего канала `output`.

---

### **Вопросы:**

**1. Как избежать утечек горутин?**  
→ Обеспечить, что каждая горутина завершится. Для этого:
- Обязательно закрывать каналы (`defer close(out)`)
- Гарантировать, что `range` по входному каналу завершится — нужно **закрыть** предыдущий канал

**2. Как организовать закрытие каналов?**  
→ Используем `defer close(out)` внутри каждой горутины. Канал закрывается, когда входной канал полностью прочитан.

---

### **Исправления в коде (если бы были ошибки):**

**Ошибка:** забыли `close(out)` — следующая стадия может зависнуть в `range`, ожидая данные.  
**Ошибка:** отправка в закрытый канал — если `defer close(out)` стоит до `for` (вне горутины) или если `out` передаётся куда-то дальше после закрытия.

--- 

## **Реализация кейса 3: Мини-вебсервер**

### **Задача:**
Реализовать HTTP-сервер на Go, который:
- Хранит количество запросов от каждого IP-адреса клиента.
- При каждом новом запросе возвращает число обращений с этого IP.

---

### **Код:**
```go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

type Server struct {
	mu    sync.Mutex
	count map[string]int
}

func NewServer() *Server {
	return &Server{
		count: make(map[string]int),
	}
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "invalid remote addr", http.StatusInternalServerError)
		return
	}

	s.mu.Lock()
	s.count[ip]++
	n := s.count[ip]
	s.mu.Unlock()

	fmt.Fprintf(w, "Hello, %s! You've visited %d times.\n", ip, n)
}

func main() {
	server := NewServer()
	http.HandleFunc("/", server.handler)
	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

### **Объяснение:**

- `Server` содержит карту `map[string]int` — счётчики запросов по IP.
- `sync.Mutex` нужен, чтобы защитить доступ к мапе от гонок данных.
- IP получается через `net.SplitHostPort(r.RemoteAddr)` — отсекаем порт.
- В `handler`:
  - Получаем IP
  - Блокируем мьютекс → обновляем счётчик → разблокируем
  - Возвращаем ответ с количеством посещений

---

### **Вопросы:**

**1. Как обеспечить конкурентный доступ к мапе?**  
→ Использовать `sync.Mutex` или `sync.RWMutex`. Можно также `sync.Map`, но `Mutex` здесь проще и эффективнее (при малом числе горутин).

**2. Какие структуры лучше использовать: `sync.Map`, `map+Mutex`, `atomic`?**  
| Структура | Когда лучше использовать |
|----------|--------------------------|
| `map+Mutex` | Подходит в большинстве случаев, особенно если доступ сбалансирован |
| `sync.Map` | Удобен, если много читающих и редко пишут |
| `atomic` | Используется с `int32/int64`, но не для сложных структур как `map` |

---

### **Возможные ошибки и исправления:**

**Ошибка:** доступ к `map` без блокировки → может привести к гонке данных.  
**Решение:** использовать `sync.Mutex` или `sync.RWMutex`.

**Ошибка:** неправильный разбор IP — `RemoteAddr` содержит IP:порт.  
**Решение:** `net.SplitHostPort(r.RemoteAddr)` вернёт IP и порт по отдельности.

---

## **Реализация кейса 4: Собственный JSON парсер**

### **Задача:**
Реализовать собственный парсер JSON-объекта вида:  
```json
{"a":1,"b":2}
```

Результат — `map[string]int`.

---

### **Код:**
```go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseSimpleJSON(input string) (map[string]int, error) {
	result := make(map[string]int)
	input = strings.TrimSpace(input)

	if len(input) < 2 || input[0] != '{' || input[len(input)-1] != '}' {
		return nil, errors.New("invalid JSON object")
	}

	// Убираем внешние скобки
	content := input[1 : len(input)-1]

	// Простейший парсинг через state-machine
	var (
		key        string
		value      int
		inQuotes   bool
		expectKey  = true
		expectColon = false
		expectValue = false
		token      strings.Builder
	)

	i := 0
	for i < len(content) {
		ch := content[i]

		switch {
		case ch == '"':
			inQuotes = !inQuotes
			if !inQuotes && expectKey {
				key = token.String()
				token.Reset()
				expectKey = false
				expectColon = true
			}
			i++

		case inQuotes:
			token.WriteByte(ch)
			i++

		case ch == ':' && expectColon:
			expectColon = false
			expectValue = true
			i++

		case ch >= '0' && ch <= '9' && expectValue:
			start := i
			for i < len(content) && content[i] >= '0' && content[i] <= '9' {
				i++
			}
			numStr := content[start:i]
			val, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			value = val
			result[key] = value
			expectValue = false

		case ch == ',':
			if expectKey || expectColon || expectValue {
				return nil, errors.New("unexpected comma")
			}
			expectKey = true
			token.Reset()
			i++

		case ch == ' ' || ch == '\n' || ch == '\t':
			i++ // пропускаем пробелы

		default:
			return nil, fmt.Errorf("unexpected character: %q at pos %d", ch, i)
		}
	}

	if expectKey || expectColon || expectValue {
		return nil, errors.New("incomplete JSON object")
	}

	return result, nil
}

func main() {
	jsonStr := `{"a":1,"b":2,"c":42}`
	parsed, err := parseSimpleJSON(jsonStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed map:", parsed)
}
```

---

### **Объяснение:**

- Мы **не используем `encoding/json`**, потому что цель — научиться **писать парсер вручную**.
- Реализация использует **примитивный state machine**:
  - `"key"` → `:` → `value` → `,` или `}` — строгое ожидание этих состояний.
- `inQuotes` отслеживает, находимся ли внутри строки.
- `expectKey`, `expectColon`, `expectValue` — булевые флаги, определяющие, что допустимо в данный момент.
- После получения ключа и значения — они добавляются в `map[string]int`.

---

### **Вопросы:**

**1. Почему `encoding/json` не подходит?**  
→ Потому что задача в лайвкодинге — реализовать свой парсер, понять, как работает лексер/парсер, и проанализировать, где возникают сложности.

**2. Как организовать state machine?**  
→ Через набор флагов `expectX` + пошаговый парсинг по каждому символу. Каждый переход изменяет контекст ожидания следующего символа.

---

### **Ошибки и исправления:**

**Ошибка:** неправильный парсинг числа, если оно больше 1 цифры.  
→ **Решение:** использовать цикл для сбора всех цифр `while content[i] >= '0' && content[i] <= '9'`.

**Ошибка:** забыли сбросить флаги `expectKey`, `expectValue`, `expectColon`.  
→ **Решение:** сбрасываем после каждого успешного шага.

**Ошибка:** парсер не обрабатывает вложенные объекты или строки с экранированием.  
→ **Решение:** это ограничение **упрощённой** реализации — обсуждаем как часть "урезанного" DSL.

---

## **Реализация кейса 5: Детект зацикливания в графе**

### **Задача:**
Дан ориентированный граф, представленный в виде списка смежности:  
`map[int][]int`

Нужно определить, содержит ли граф **хотя бы один цикл**.

---

### **Код (рекурсивный DFS):**
```go
package main

import "fmt"

func hasCycle(graph map[int][]int) bool {
	visited := make(map[int]bool)
	onStack := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(node int) bool {
		if onStack[node] {
			return true // нашли цикл
		}
		if visited[node] {
			return false
		}

		visited[node] = true
		onStack[node] = true

		for _, neighbor := range graph[node] {
			if dfs(neighbor) {
				return true
			}
		}

		onStack[node] = false
		return false
	}

	for node := range graph {
		if !visited[node] && dfs(node) {
			return true
		}
	}
	return false
}

func main() {
	graph := map[int][]int{
		0: {1},
		1: {2},
		2: {0}, // цикл: 0 → 1 → 2 → 0
		3: {4},
		4: {},
	}

	fmt.Println("Has cycle:", hasCycle(graph)) // true
}
```

---

### **Объяснение:**

- `visited` — помечаем, что вершина уже обработана.
- `onStack` — показывает, что вершина находится в **текущем стеке вызова**. Если снова попадаем на неё — нашли цикл.
- `dfs` — рекурсивная функция обхода. На входе ставим `onStack = true`, на выходе — `false`.
- Проверяем каждый узел, если он ещё не был посещён.

---

### **Вопросы:**

**1. Почему нужны `visited` и `onStack`?**  
- `visited` — чтобы не обрабатывать один и тот же узел дважды.
- `onStack` — помогает поймать **обратное ребро**, то есть цикл: если во время обхода попадаем на узел, уже находящийся в текущем стеке.

**2. Как реализовать DFS без рекурсии?**  
→ Вместо вызовов — использовать **ручной стек** и отслеживать статусы (enter/exit). Ниже — пример.

---

### **Код (итеративный DFS):**
```go
func hasCycleIterative(graph map[int][]int) bool {
	type Frame struct {
		node    int
		visited bool
	}
	visited := make(map[int]bool)
	onStack := make(map[int]bool)

	for start := range graph {
		if visited[start] {
			continue
		}

		stack := []Frame{{start, false}}

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top.visited {
				onStack[top.node] = false
				continue
			}

			if onStack[top.node] {
				return true
			}

			visited[top.node] = true
			onStack[top.node] = true
			stack = append(stack, Frame{top.node, true}) // выход

			for i := len(graph[top.node]) - 1; i >= 0; i-- {
				neighbor := graph[top.node][i]
				if !visited[neighbor] {
					stack = append(stack, Frame{neighbor, false})
				} else if onStack[neighbor] {
					return true
				}
			}
		}
	}
	return false
}
```

---

### **Разбор (итеративный вариант):**

- Используем структуру `Frame`, чтобы различать вход и выход узла.
- При первом заходе на узел → ставим `onStack = true`.
- При выходе — снимаем `onStack = false`.
- Обрабатываем соседей в обратном порядке, чтобы сохранить правильный порядок обхода.

---

### **Ошибки и исправления:**

**Ошибка:** забыли снять `onStack = false` после обработки — приведёт к ложноположительным срабатываниям.  
**Решение:** обязательно удалять узел из стека после выхода из `dfs`.

**Ошибка:** при итеративной реализации перепутан порядок обработки узлов.  
**Решение:** использовать enter/exit фреймы (`visited bool`), как в `Frame`.

---

## **Реализация кейса 6: EventBus**

### **Задача:**
Реализовать простой `EventBus`, который:
- Позволяет подписаться на событие по имени (`Subscribe(event string) <-chan any`)
- Позволяет отправить событие всем подписчикам (`Publish(event string, data any)`)
- Гарантирует безопасность при конкурентной работе

---

### **Код:**
```go
package main

import (
	"fmt"
	"sync"
)

type EventBus struct {
	subscribers map[string][]chan any
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan any),
	}
}

// Subscribe возвращает канал, в который будут приходить события
func (eb *EventBus) Subscribe(event string) <-chan any {
	ch := make(chan any, 10) // буферизированный канал
	eb.mu.Lock()
	eb.subscribers[event] = append(eb.subscribers[event], ch)
	eb.mu.Unlock()
	return ch
}

// Publish отправляет событие всем подписчикам
func (eb *EventBus) Publish(event string, data any) {
	eb.mu.RLock()
	subs := eb.subscribers[event]
	eb.mu.RUnlock()

	for _, ch := range subs {
		select {
		case ch <- data:
		default:
			// если канал переполнен — пропускаем
			// можно логировать или обрабатывать иначе
		}
	}
}

// Optional: Unsubscribe
func (eb *EventBus) Unsubscribe(event string, target <-chan any) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	chans := eb.subscribers[event]
	for i, ch := range chans {
		if ch == target {
			eb.subscribers[event] = append(chans[:i], chans[i+1:]...)
			close(ch)
			break
		}
	}
}

// Пример использования
func main() {
	bus := NewEventBus()

	sub := bus.Subscribe("hello")

	go func() {
		for msg := range sub {
			fmt.Println("Received:", msg)
		}
	}()

	bus.Publish("hello", "world")
	bus.Publish("hello", 42)

	// Подождём для вывода
	select {}
}
```

---

### **Объяснение:**

- Используется `map[string][]chan any` — один ключ = список каналов-подписчиков.
- `sync.RWMutex`:
  - `RLock` — при публикации (только чтение)
  - `Lock` — при подписке/отписке (модификация)
- Буферизированные каналы (например, `chan any, 10`) — защищают от блокировки при медленных потребителях.
- В `Publish`: если канал полон — не блокируемся, можно логировать или делать retry.

---

### **Вопросы:**

**1. Какую структуру использовать для хранения подписчиков?**  
→ `map[string][]chan any` — позволяет поддерживать множественные подписки на одно событие.

**2. Как обрабатывать конкурентную отправку сообщений?**  
→ Через `sync.RWMutex`:
- `RLock` для отправки — безопасно, потому что не меняем map
- `Lock` для добавления/удаления подписчиков
- Каналы — буферизированные, чтобы избежать блокировки при `Publish`

---

### **Ошибки и исправления:**

**Ошибка:** отправка в небуферизированный канал без получения — может заблокировать `Publish`.  
**Решение:** используем буфер или отправляем через `select` с `default`, чтобы не блокировать.

**Ошибка:** забыли закрыть канал при `Unsubscribe` → утечка горутин.  
**Решение:** `close(ch)` при удалении из подписчиков.

---

## **Реализация кейса 7: Кастомный мьютекс**

### **Задача:**
Реализовать собственную структуру `MyMutex`, аналог `sync.Mutex`, используя только **каналы**.  
Поддерживаем методы:
- `Lock()`
- `Unlock()`

---

### **Код:**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type MyMutex struct {
	ch chan struct{}
}

func NewMyMutex() *MyMutex {
	m := &MyMutex{
		ch: make(chan struct{}, 1), // семафор
	}
	m.ch <- struct{}{} // сразу "разрешено"
	return m
}

func (m *MyMutex) Lock() {
	<-m.ch // блокируемся, если занято
}

func (m *MyMutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
		// успешно разблокировали
	default:
		panic("unlock of unlocked mutex")
	}
}

func main() {
	m := NewMyMutex()
	counter := 0
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				m.Lock()
				counter++
				fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)
				m.Unlock()
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
```

---

### **Объяснение:**

- Канал `make(chan struct{}, 1)` работает как **семафор с 1 разрешением**:
  - `Lock()` — забирает разрешение (`<-m.ch`), блокирует если пусто
  - `Unlock()` — возвращает разрешение (`m.ch <- struct{}{}`)
- При инициализации `NewMyMutex()` мы сразу кладём `struct{}{}` в канал — разрешая первую блокировку.
- Используем `select` в `Unlock()` с `default` — если `Unlock` вызван дважды подряд, будет panic.

---

### **Вопросы:**

**1. Как использовать каналы как блокировку?**  
→ Через канал ёмкостью 1:
- Забрал значение → заблокировал
- Вернул значение → разблокировал

**2. Какие edge-cases могут быть при использовании?**
- ❗ **Повторный Unlock без Lock** → panic (или утечка, если не отлавливать)
- ❗ Потерянный Unlock (если не вызван после Lock) → горутина навсегда зависнет
- ❗ Нет проверки владельца (в отличие от `sync.Mutex`, можно Unlock из другой горутины)

---

### **Ошибки и исправления:**

**Ошибка:** забыли инициализировать `m.ch <- struct{}{}` → `Lock()` никогда не разблокируется  
**Решение:** при `NewMyMutex()` сразу добавляем один элемент в канал

**Ошибка:** `Unlock()` без `select` может повесить горутину, если канал уже полон  
**Решение:** используем `select { case ch <- struct{}{}: default: panic(...) }`

---

## **Реализация кейса 8: Мини-SQL интерпретатор**

### **Задача:**
Реализовать парсер простого SQL-запроса:
```
SELECT name FROM users WHERE age > 20
```
Результат — **структура запроса** (`AST`), содержащая:
- список полей (например, `["name"]`)
- имя таблицы (`users`)
- фильтр (`age > 20`)

---

### **Код:**
```go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Структура для фильтра
type Condition struct {
	Field    string
	Operator string
	Value    int
}

// Структура SQL-запроса
type Query struct {
	Fields    []string
	TableName string
	Where     *Condition
}

func parseSQL(query string) (*Query, error) {
	tokens := strings.Fields(query)
	if len(tokens) < 4 {
		return nil, errors.New("too short for valid SQL")
	}

	// Начальная проверка
	if strings.ToUpper(tokens[0]) != "SELECT" {
		return nil, errors.New("query must start with SELECT")
	}

	// State machine
	state := "SELECT"
	result := &Query{}

	i := 1
	for i < len(tokens) {
		token := tokens[i]

		switch state {
		case "SELECT":
			if strings.ToUpper(token) == "FROM" {
				state = "FROM"
				i++
				continue
			}
			result.Fields = append(result.Fields, strings.Trim(token, ","))
			i++

		case "FROM":
			result.TableName = token
			i++
			if i < len(tokens) && strings.ToUpper(tokens[i]) == "WHERE" {
				state = "WHERE"
				i++
			} else {
				state = "END"
			}

		case "WHERE":
			if i+2 >= len(tokens) {
				return nil, errors.New("invalid WHERE clause")
			}
			field := tokens[i]
			operator := tokens[i+1]
			valueStr := tokens[i+2]
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return nil, errors.New("expected numeric value in WHERE")
			}
			result.Where = &Condition{
				Field:    field,
				Operator: operator,
				Value:    value,
			}
			i += 3
			state = "END"

		case "END":
			i++

		default:
			return nil, errors.New("unknown state")
		}
	}

	return result, nil
}

func main() {
	query := `SELECT name FROM users WHERE age > 20`
	parsed, err := parseSQL(query)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}
	fmt.Printf("Parsed query:\nFields: %v\nTable: %s\nWhere: %+v\n",
		parsed.Fields, parsed.TableName, parsed.Where)
}
```

---

### **Объяснение:**

- Используем `strings.Fields` для **токенизации** — быстро и просто, работает для плоских SQL без кавычек/скобок.
- Дальше идёт **state machine**:
  - `"SELECT"` — парсим поля до `FROM`
  - `"FROM"` — имя таблицы
  - `"WHERE"` — фильтр: `поле`, `оператор`, `число`
- Храним результат в структуре `Query` и условие в `Condition`.

---

### **Вопросы:**

**1. Как разбивать строку на токены?**  
→ В простейшем случае — через `strings.Fields()`.  
Для более сложного — лексер с анализом символов и кавычек.

**2. Что проще: парсить вручную или использовать генератор?**  
| Подход | Когда использовать |
|--------|--------------------|
| Вручную | Просто, быстро, гибко. Подходит для мини-Domain Language |
| Генератор (`goyacc`, `antlr`) | При необходимости полного SQL или сложных выражений/группировок |

---

### **Ошибки и исправления:**

**Ошибка:** не обрабатываются запятые между полями (например: `SELECT a, b`).  
→ **Решение:** обрезать `,` с конца каждого поля `strings.Trim(token, ",")`

**Ошибка:** неправильное количество токенов в WHERE  
→ **Решение:** проверка `i+2 >= len(tokens)` до разбора условия

**Ошибка:** без учёта регистра (`select` не сработает)  
→ **Решение:** использовать `strings.ToUpper(...)` для ключевых слов

---

## **Реализация кейса 9: Таймаут и отмена операций**

### **Задача:**
Реализовать функцию, которая отправляет HTTP-запрос к серверу с **таймаутом 2 секунды**.  
Если сервер не отвечает — операция отменяется.

---

### **Код:**
```go
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchWithTimeout(url string) error {
	// Создаём контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // важно: всегда вызывать, чтобы освободить ресурсы

	// Создаём запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	// Отправляем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Если контекст отменён — проверим это явно
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("request cancelled due to timeout")
		}
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	fmt.Println("Response:", string(body))
	return nil
}

func main() {
	err := fetchWithTimeout("http://httpbin.org/delay/3") // задержка 3 сек
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```

---

### **Объяснение:**

- `context.WithTimeout` создаёт контекст, который будет отменён через 2 секунды.
- `defer cancel()` **обязательно** — он освобождает ресурсы таймера (и может отменить запрос по завершению функции).
- Контекст передаётся в `http.NewRequestWithContext`, чтобы `http.Client` мог его уважать.
- Если сервер не ответит за 2 секунды — запрос прервётся с ошибкой `context deadline exceeded`.

---

### **Вопросы:**

**1. Как работает `context.WithTimeout`?**  
→ Он создаёт обёртку над `context.Context`, которая автоматически завершится через заданный `timeout`. После этого:
- `ctx.Err()` возвращает `context.DeadlineExceeded`
- Все функции, уважающие контекст, будут отменены

**2. Что произойдёт, если не вызвать `cancel()`?**  
→ Утечка ресурса: таймер будет висеть в памяти, даже если запрос уже завершился.  
Всегда делай `defer cancel()`!

---

### **Ошибки и исправления:**

**Ошибка:** забыли `defer cancel()` → утечка контекста  
→ **Решение:** всегда использовать `defer cancel()` сразу после создания контекста

**Ошибка:** не передали контекст в `NewRequestWithContext`  
→ **Решение:** нельзя использовать обычный `http.NewRequest`, нужно `NewRequestWithContext`

**Ошибка:** неправильно обрабатываем ошибку таймаута  
→ **Решение:** проверять `ctx.Err() == context.DeadlineExceeded`, а не только `err != nil`

---

## **Реализация кейса 10: Трэдсейфовый счётчик**

### **Задача:**
Реализовать **безопасный счётчик**, который:
- Поддерживает инкремент и чтение текущего значения
- Работает корректно в многопоточном окружении

---

### **Код:**
```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	value int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
}
```

---

### **Объяснение:**

- `atomic.AddInt64(&c.value, 1)` — атомарно увеличивает счётчик на 1
- `atomic.LoadInt64(&c.value)` — атомарно читает значение
- Такой способ быстрее и эффективнее, чем `sync.Mutex`, особенно если операция простая (например, инкремент)

---

### **Вопросы:**

**1. Чем `atomic` лучше `Mutex`?**  
→ Для простых операций (инкремент, считывание) `atomic`:
- Быстрее: не требует блокировок
- Безопасен: использует низкоуровневые атомарные инструкции процессора
- Идеален для счётчиков, статистик, метрик

**2. Когда использовать `int32`, а когда `int64`?**  
→ Зависит от:
- **Архитектуры**: на 32-битной системе предпочтительно `int32`
- **Диапазона значений**: если значение может превышать ~2 млрд — лучше использовать `int64`
- В многопоточном коде **нельзя использовать `int` с atomic**, только фиксированные типы (`int32`, `int64`, `uint32` и т.д.)

---

### **Ошибки и исправления:**

**Ошибка:** использование `int` в `atomic`-функциях  
→ **Решение:** использовать `int64` или `int32` — строго типизировано

**Ошибка:** микс `atomic` и обычного доступа к переменной  
→ **Решение:** **всегда** использовать только `atomic` функции для доступа и изменения переменной — иначе гонки

**Ошибка:** использовать `atomic` для комплексных операций (например, `if x == 5 { x = 10 }`)  
→ **Решение:** для сложных операций — использовать `sync.Mutex`

---
