#!/bin/bash

# Перемещает задачи из problems/ в подпапки типа 0-99, 100-199 и т.д.

base="problems"

for folder in "$base"/*/; do
    folder=${folder%/}
    name=$(basename "$folder")

    # Пропустить подпапки уже вида "0-99"
    if [[ "$name" =~ ^[0-9]+-[0-9]+$ ]]; then
        continue
    fi

    # Извлекаем номер задачи из имени (до первой буквы)
    if [[ "$name" =~ ^([0-9]+) ]]; then
        number=${BASH_REMATCH[1]}
        start=$(( (number / 100) * 100 ))
        end=$(( start + 99 ))
        range_folder="${base}/${start}-${end}"

        # Создаем папку, если не существует
        mkdir -p "$range_folder"

        # Перемещаем задачу
        mv "$base/$name" "$range_folder/"
        echo "✅ Перемещено: $name → $range_folder/"
    fi
done

echo "🎉 Перемещение завершено!"
