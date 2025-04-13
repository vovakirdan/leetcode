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
}

func extractTitleAndLevel(readmePath string) (string, string) {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return "Unknown", "Unknown"
	}

	// Поиск заголовка и уровня сложности (ищем "Easy", "Medium", "Hard")
	titleRe := regexp.MustCompile(`(?m)^#\s+(.*)`)
	levelRe := regexp.MustCompile(`(?i)Level:\s*(Easy|Medium|Hard)`)

	title := "Unknown"
	level := "Unknown"

	if m := titleRe.FindSubmatch(content); m != nil {
		title = string(m[1])
	}
	if m := levelRe.FindSubmatch(content); m != nil {
		level = cases.Title(language.English).String(string(m[1]))
	}

	return title, level
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
			title, level := extractTitleAndLevel(readmePath)

			if title == "Unknown" || level == "Unknown" {
				fmt.Fprintf(os.Stderr, "⚠️  Пропущено: %s — отсутствует корректный заголовок или уровень сложности\n", slug)
				skipped++
				continue
			}

			number := extractNumberFromSlug(slug)
			problems = append(problems, Problem{Number: number, Slug: slug, Title: title, Level: level})
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	var builder strings.Builder
	builder.WriteString("# LeetCode Solutions in Go\n\n")
	builder.WriteString("Решения задач с [leetcode.com](https://leetcode.com) с пояснением, кодом и тестами.\n\n")
	builder.WriteString("## 🔗 Задачи\n\n")
	builder.WriteString("| # | Название задачи | Уровень | Решение | Описание |\n")
	builder.WriteString("|---|------------------|---------|---------|----------|\n")

	for i, p := range problems {
		// Получим slug из названия задачи
		slug := strings.ToLower(regexp.MustCompile(`([a-z])([A-Z])`).ReplaceAllString(p.Title, "$1-$2"))

		builder.WriteString(fmt.Sprintf(
			"| %d | [%d. %s](https://leetcode.com/problems/%s/) | %s | [Code](problems/%s/solution.go) | [Explanation](problems/%s/README.md) |\n",
			i+1, p.Number, p.Title, slug, p.Level, p.Slug, p.Slug))
	}

	err = os.WriteFile("README.md", []byte(builder.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("✅ README.md успешно обновлён. Пропущено задач: %d\n", skipped)
}
