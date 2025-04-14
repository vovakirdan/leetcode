#!/bin/bash

# Проверка аргументов
if [ $# -lt 3 ]; then
    echo "Использование: $0 <номер_задачи> <название задачи в CamelCase> [Easy|Medium|Hard]"
    echo "Пример: $0 136 SingleNumber Easy"
    exit 1
fi

# Извлечение номера и названия
number="$1"
title_camel="$2"
difficulty="$3"
folder="${number}${title_camel}"

# Проверка существования папки
if [ -d "problems/$folder" ]; then
    echo "❌ Ошибка: Папка problems/$folder уже существует"
    exit 1
fi

# slug для ссылки на leetcode (только название задачи)
slug=$(echo "$title_camel" | sed -E 's/([a-z])([A-Z])/\1-\2/g' | tr '[:upper:]' '[:lower:]')

# название пакета — только из названия (в нижнем регистре)
pkg=$(echo "$title_camel" | tr '[:upper:]' '[:lower:]')

# Заголовок для README (разделяем CamelCase пробелами)
title_human=$(echo "$title_camel" | sed -E 's/([a-z])([A-Z])/\1 \2/g')

# Создание структуры
mkdir -p "problems/$folder"

# README.md
cat > "problems/$folder/README.md" <<EOF
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

## 📄 Пример:

Вход:

```

```

Выход:

```

```

---

## 📝 Решение:

```go

```

EOF

# solution.go
cat > "problems/$folder/solution.go" <<EOF
package $pkg

// TODO: реализовать решение задачи $title_human
EOF

# solution_test.go
cat > "problems/$folder/solution_test.go" <<EOF
package $pkg

import "testing"

func Test${title_camel}(t *testing.T) {
	// TODO: реализовать тесты
}
EOF

echo "✅ Задача $number $title_human успешно создана в problems/$folder"