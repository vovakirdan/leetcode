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
	slugRe := regexp.MustCompile(`\[Ссылка на задачу\]\(https://leetcode.com/problems/([a-z0-9\-]+)/\)`) // group 1 is slug
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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data QuestionData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	tags := []string{}
	for _, tag := range data.Question.TopicTags {
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

	s := string(content)
	tagsRe := regexp.MustCompile(`(?m)^## 🏷 Теги:.*?(?:(?:\n---)|\z)`) // замена существующей секции

	if tagsRe.MatchString(s) {
		s = tagsRe.ReplaceAllString(s, tagSection+"\n---")
	} else {
		// Вставить перед последним ---
		insertBefore := strings.LastIndex(s, "---")
		if insertBefore != -1 {
			s = s[:insertBefore] + tagSection + "\n" + s[insertBefore:]
		} else {
			s += "\n" + tagSection + "\n---"
		}
	}

	return os.WriteFile(path, []byte(s), 0644)
}

func UpdateReadmeWithTags() {
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

func main() {
	UpdateReadmeWithTags()
}
