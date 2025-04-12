#!/bin/bash

# Проверка наличия аргументов
if [ $# -ne 1 ]; then
    echo "Использование: $0 <название_задачи>"
    exit 1
fi

# Преобразование названия задачи в slug
slug=$(echo "$1" | tr '[:upper:]' '[:lower:]' | sed 's/ /-/g')

# Создание директории для задачи
mkdir -p "problems/$slug"

# Создание файла README.md
touch "problems/$slug/README.md"

# создание go файла
touch "problems/$slug/solution.go"

# создание файла тестов
touch "problems/$slug/solution_test.go"

echo "✅ Задача $1 успешно создана в директории problems/$slug"
