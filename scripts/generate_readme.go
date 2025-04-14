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
	Slug   string
	Title  string
	Level  string
	Link   string
}

func extractMetadata(readmePath string) (string, string, string) {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return "Unknown", "Unknown", ""
	}

	// Поиск заголовка и уровня сложности (ищем "Easy", "Medium", "Hard")
	titleRe := regexp.MustCompile(`(?m)^#\s+(.*)`)
	levelRe := regexp.MustCompile(`(?i)Level:\s*(Easy|Medium|Hard)`)
	linkRe := regexp.MustCompile(`\[Ссылка на задачу\]\((https://leetcode\.com/problems/[^)]+)\)`)

	title := "Unknown"
	level := "Unknown"
	link := ""

	if m := titleRe.FindSubmatch(content); m != nil {
		title = string(m[1])
	}
	if m := levelRe.FindSubmatch(content); m != nil {
		level = cases.Title(language.English).String(string(m[1]))
	}
	if m := linkRe.FindSubmatch(content); m != nil {
		link = string(m[1])
	}

	return title, level, link
}

func extractNumberFromSlug(slug string) int {
	re := regexp.MustCompile(`^(\d+)`)
	if m := re.FindStringSubmatch(slug); m != nil {
		num, _ := strconv.Atoi(m[1])
		return num
	}
	return 0
}

func main() {
	problemsDir := "problems"
	var problems []Problem

	entries, err := os.ReadDir(problemsDir)
	if err != nil {
		panic(err)
	}

	skipped := 0

	for _, entry := range entries {
		if entry.IsDir() {
			slug := entry.Name()
			readmePath := filepath.Join(problemsDir, slug, "README.md")
			title, level, link := extractMetadata(readmePath)

			if title == "Unknown" || level == "Unknown" || link == "" {
				fmt.Fprintf(os.Stderr, "⚠️  Пропущено: %s — не хватает title/level/link\n", slug)
				skipped++
				continue
			}

			number := extractNumberFromSlug(slug)
			problems = append(problems, Problem{Number: number, Slug: slug, Title: title, Level: level, Link: link})
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	var builder strings.Builder
	builder.WriteString("# LeetCode Solutions in Go\n\n")
	builder.WriteString("🎯 Этот репозиторий содержит решения задач с [LeetCode](https://leetcode.com), написанные на Go. Каждое решение включает объяснение, код и тесты.\n\n")
	builder.WriteString("## 📦 Фичи\n\n")
	builder.WriteString("- 🧠 Структура по задачам (`problems/<номер><Название>`)\n")
	builder.WriteString("- ✅ Тесты через `go test ./...`\n")
	builder.WriteString("- 🛠 Генерация задач скриптом\n")
	builder.WriteString("- 🔄 Автоматическая генерация таблицы задач в `README.md`\n")
	builder.WriteString("- 🧪 Автопроверки при коммите (`gofmt`, `go vet`, `README`)\n\n")

	builder.WriteString("## 🛠 Makefile команды\n\n")
	builder.WriteString("| Команда | Описание |\n")
	builder.WriteString("|--------|----------|\n")
	builder.WriteString("| `make [new/new-solution/new-problem] number=123 title=MyTitle level=Easy` | Создать новую задачу |\n")
	builder.WriteString("| `make update-readme` | Перегенерировать README.md |\n")
	builder.WriteString("| `make test` | Прогнать все тесты |\n")
	builder.WriteString("| `make check` | Проверка gofmt и go vet |\n")
	builder.WriteString("| `make install-hooks` | Установить pre-commit hook |\n\n")

	builder.WriteString("## 🔗 Задачи\n\n")
	builder.WriteString("| # | Название задачи | Уровень | Решение | Описание |\n")
	builder.WriteString("|---|------------------|---------|---------|----------|\n")

	for i, p := range problems {
		builder.WriteString(fmt.Sprintf(
			"| %d | [%d. %s](%s) | %s | [Code](problems/%s/solution.go) | [Explanation](problems/%s/README.md) |\n",
			i+1, p.Number, p.Title, p.Link, p.Level, p.Slug, p.Slug))
	}

	err = os.WriteFile("README.md", []byte(builder.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("✅ README.md успешно обновлён. Пропущено задач: %d\n", skipped)
}
