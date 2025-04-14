#!/bin/bash

# Проверка аргументов
if [ $# -lt 3 ]; then
    echo "Использование: $0 <номер_задачи> <название задачи в CamelCase> [Easy|Medium|Hard]"
    echo "Пример: $0 136 SingleNumber Easy"
    exit 1
fi

number="$1"
title_camel="$2"
difficulty="$3"

# Название задачи в camel case → пробельное → slug
title_human=$(echo "$title_camel" | sed -E 's/([a-z])([A-Z])/\1 \2/g')
slug=$(echo "$title_camel" | sed -E 's/([a-z])([A-Z])/\1-\2/g' | tr '[:upper:]' '[:lower:]')
pkg=$(echo "$title_camel" | tr '[:upper:]' '[:lower:]')

# Вычисление подпапки: 0-99, 100-199 и т.д.
start=$(( (number / 100) * 100 ))
end=$(( start + 99 ))
range_folder="problems/${start}-${end}"

# Папка задачи
task_folder="${range_folder}/${number}${title_camel}"

# Проверка на существование
if [ -d "$task_folder" ]; then
    echo "❌ Папка уже существует: $task_folder"
    exit 1
fi

# Создание иерархии
mkdir -p "$task_folder"

# README.md
cat > "$task_folder/README.md" <<EOF
# $title_human

Level: $difficulty

[Ссылка на задачу](https://leetcode.com/problems/$slug/)

## 🧠 Задача:

...

---

## 📌 Идея:

...

---

## 📏 Структура:

...

---

## 🔁 Шаги алгоритма:

...

---

## ⏱️ Сложность:

- Время:
- Память:

---

## 📝 Пример:

Вход:

\`\`\`

\`\`\`

Выход:

\`\`\`

\`\`\`

---

## 🏷 Теги:

_(заполнится после \`make fetch-tags\`)_

---

## 💡 Решение:

\`\`\`go

\`\`\`
EOF

# solution.go
cat > "$task_folder/solution.go" <<EOF
package $pkg

// TODO: реализовать решение задачи $title_human
EOF

# solution_test.go
cat > "$task_folder/solution_test.go" <<EOF
package $pkg

import "testing"

func Test${title_camel}(t *testing.T) {
	// TODO: реализовать тесты
}
EOF

echo "✅ Задача $number $title_human успешно создана в $task_folder"
