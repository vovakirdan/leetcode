package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Tag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type QuestionData struct {
	Question struct {
		TopicTags []Tag `json:"topicTags"`
	} `json:"question"`
}

const graphqlURL = "https://leetcode.com/graphql"
const query = `
query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    topicTags {
      name
      slug
    }
  }
}`

func extractSlug(readme string) string {
	slugRe := regexp.MustCompile(`\[Ссылка на задачу\]\(https://leetcode.com/problems/([a-z0-9\-]+)/?\)`) // group 1 is slug
	if m := slugRe.FindStringSubmatch(readme); m != nil {
		return m[1]
	}
	return ""
}

func fetchTags(slug string) ([]string, error) {
	payload := map[string]interface{}{
		"operationName": "questionData",
		"query":         query,
		"variables": map[string]string{
			"titleSlug": slug,
		},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(payload)

	req, err := http.NewRequest("POST", graphqlURL, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", fmt.Sprintf("https://leetcode.com/problems/%s/", slug))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Log the response for debugging
	// log.Printf("Response for %s: %s", slug, string(body))

	var response struct {
		Data struct {
			Question struct {
				TopicTags []Tag `json:"topicTags"`
			} `json:"question"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	tags := []string{}
	for _, tag := range response.Data.Question.TopicTags {
		tags = append(tags, tag.Name)
	}
	return tags, nil
}

func updateReadmeWithTags(path string, tags []string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tagSection := "## 🏷 Теги:\n"
	for _, tag := range tags {
		tagSection += fmt.Sprintf("- %s\n", tag)
	}
	tagSection += "\n---"

	s := string(content)
	tagsRe := regexp.MustCompile(`(?s)## 🏷 Теги:.*?\n---`) // (?s) - dot matches newline

	if tagsRe.MatchString(s) {
		s = tagsRe.ReplaceAllString(s, tagSection)
	} else {
		// Вставить перед последним ---
		insertBefore := strings.LastIndex(s, "---")
		if insertBefore != -1 {
			s = s[:insertBefore] + tagSection + "\n" + s[insertBefore:]
		} else {
			s += "\n" + tagSection
		}
	}

	return os.WriteFile(path, []byte(s), 0644)
}

func Run() {
	baseDir := "problems"
	err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, "README.md") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		slug := extractSlug(string(content))
		if slug == "" {
			fmt.Fprintf(os.Stderr, "⚠️  Пропуск: %s (нет ссылки на задачу)\n", path)
			return nil
		}

		tags, err := fetchTags(slug)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Ошибка загрузки тегов для %s: %v\n", slug, err)
			return nil
		}

		if len(tags) == 0 {
			fmt.Printf("ℹ️  Нет тегов для %s\n", slug)
			return nil
		}

		err = updateReadmeWithTags(path, tags)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Ошибка обновления README %s: %v\n", path, err)
		} else {
			fmt.Printf("✅ Теги обновлены: %s\n", path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func RunForOneProblem(number string) {
	var problemDir string

	err := filepath.WalkDir("problems", func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() {
			return nil
		}
		if filepath.Base(filepath.Dir(path)) == "problems" {
			// Пропускаем верхнеуровневые диапазоны (0-99 и т.п.)
			return nil
		}
		if strings.HasPrefix(filepath.Base(path), number) {
			problemDir = path
			return filepath.SkipAll
		}
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Ошибка при поиске директории задачи: %v\n", err)
		return
	}

	if problemDir == "" {
		fmt.Fprintf(os.Stderr, "❌ Не найдена директория для задачи %s\n", number)
		return
	}

	readme := filepath.Join(problemDir, "README.md")
	if _, err := os.Stat(readme); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "⚠️  Пропущено: %s (файл README не найден)\n", readme)
		return
	}

	content, err := os.ReadFile(readme)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Ошибка чтения README: %v\n", err)
		return
	}

	slug := extractSlug(string(content))
	if slug == "" {
		fmt.Fprintf(os.Stderr, "❌ Не удалось извлечь slug из README\n")
		return
	}

	tags, err := fetchTags(slug)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Ошибка получения тегов: %v\n", err)
		return
	}

	if err := updateReadmeWithTags(readme, tags); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Ошибка обновления README: %v\n", err)
		return
	}

	fmt.Printf("✅ Теги обновлены для задачи %s\n", number)
}

func main() {
	if len(os.Args) > 1 {
		RunForOneProblem(os.Args[1])
	} else {
		Run()
	}
}
