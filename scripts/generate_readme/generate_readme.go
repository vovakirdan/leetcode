package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Problem struct {
	Number int
	Path   string // полный путь от problems/... до задачи
	Title  string
	Level  string
	Link   string
	Tags   []string
}

func extractMetadata(readmePath string) (string, string, string, []string) {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return "Unknown", "Unknown", "", nil
	}

	titleRe := regexp.MustCompile(`(?m)^#\s+(.*)`)
	levelRe := regexp.MustCompile(`(?i)Level:\s*(Easy|Medium|Hard)`)
	linkRe := regexp.MustCompile(`\[Ссылка на задачу\]\((https://leetcode\.com/problems/[^)]+)\)`)
	tagsRe := regexp.MustCompile(`(?s)## 🏷 Теги:\s*\n((?:- [^\n]+\n)+)(?:\n---)?`)

	title, level, link := "Unknown", "Unknown", ""
	var tags []string

	if m := titleRe.FindSubmatch(content); m != nil {
		title = string(m[1])
	}
	if m := levelRe.FindSubmatch(content); m != nil {
		level = cases.Title(language.English).String(string(m[1]))
	}
	if m := linkRe.FindSubmatch(content); m != nil {
		link = string(m[1])
	}

	// Поиск всех строк после заголовка "## Теги:"
	if tagBlock := tagsRe.FindStringSubmatch(string(content)); len(tagBlock) > 1 {
		lines := strings.Split(tagBlock[1], "\n")
		for _, line := range lines {
			clean := strings.TrimSpace(line)
			clean = strings.TrimPrefix(clean, "-")
			clean = strings.TrimSpace(clean)
			if clean != "" {
				tags = append(tags, clean)
			}
		}
	}

	return title, level, link, tags
}

func extractNumberFromFolder(name string) int {
	re := regexp.MustCompile(`^(\d+)`)
	if m := re.FindStringSubmatch(name); m != nil {
		num, _ := strconv.Atoi(m[1])
		return num
	}
	return 0
}

func generateReadme() {
	var problems []Problem
	skipped := 0

	err := filepath.WalkDir("problems", func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() {
			return nil
		}

		readmePath := filepath.Join(path, "README.md")
		if _, err := os.Stat(readmePath); err != nil {
			return nil
		}

		// Папка задачи — последняя часть пути
		_, taskDir := filepath.Split(path)
		title, level, link, tags := extractMetadata(readmePath)

		if title == "Unknown" || level == "Unknown" || link == "" {
			fmt.Fprintf(os.Stderr, "⚠️  Пропущено: %s — нет title/level/link\n", path)
			skipped++
			return nil
		}

		number := extractNumberFromFolder(taskDir)
		relativePath := strings.TrimPrefix(path, "problems/")

		problems = append(problems, Problem{
			Number: number,
			Path:   relativePath,
			Title:  title,
			Level:  level,
			Link:   link,
			Tags:   tags,
		})

		return nil
	})
	if err != nil {
		panic(err)
	}

	// Сортировка по номеру задачи
	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	// Генерация README.md
	var builder strings.Builder
	builder.WriteString("# LeetCode Solutions in Go\n\n")
	builder.WriteString("🎯 Этот репозиторий содержит решения задач с [LeetCode](https://leetcode.com), написанные на Go. Каждое решение включает объяснение, код и тесты.\n\n")

	builder.WriteString("## 📦 Фичи\n\n")
	builder.WriteString("- 🧠 Задачи структурированы в папках `problems/<range>/<номер><Название>`\n")
	builder.WriteString("- ✅ Тесты запускаются через `go test ./...`\n")
	builder.WriteString("- 🔄 README обновляется автоматически `make update-readme`\n")
	builder.WriteString("- 🏷 Теги загружаются через `make fetch-tags`\n")
	builder.WriteString("- ⚙️ Хуки при коммите: `gofmt`, `go vet`, генерация README\n\n")

	builder.WriteString("## 🛠 Makefile команды\n\n")
	builder.WriteString("| Команда | Описание |\n")
	builder.WriteString("|--------|----------|\n")
	builder.WriteString("| `make new number=123 title=MyTitle level=Easy` | Создать новую задачу |\n")
	builder.WriteString("| `make update-readme` | Перегенерировать README.md |\n")
	builder.WriteString("| `make test` | Прогнать все тесты |\n")
	builder.WriteString("| `make check` | Проверка gofmt и go vet |\n")
	builder.WriteString("| `make install-hooks` | Установить pre-commit hook |\n")
	builder.WriteString("| `make fetch-tags` | Загрузить теги в README каждой задачи |\n\n")

	builder.WriteString("## 🔗 Задачи\n\n")
	builder.WriteString("| # | Название | Уровень | Теги | Решение | Описание |\n")
	builder.WriteString("|---|----------|---------|------|---------|----------|\n")

	for i, p := range problems {
		tags := "-"
		if len(p.Tags) > 0 {
			tags = strings.Join(p.Tags, ", ")
		}
		builder.WriteString(fmt.Sprintf(
			"| %d | [%d. %s](%s) | %s | %s | [Code](problems/%s/solution.go) | [Explanation](problems/%s/README.md) |\n",
			i+1, p.Number, p.Title, p.Link, p.Level, tags, p.Path, p.Path))
	}

	err = os.WriteFile("README.md", []byte(builder.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("✅ README.md успешно обновлён. Пропущено задач: %d\n", skipped)
}

func main() {
	generateReadme()
}
