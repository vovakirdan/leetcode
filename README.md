# LeetCode Solutions in Go

🎯 Этот репозиторий содержит решения задач с [LeetCode](https://leetcode.com), написанные на Go. Каждое решение включает объяснение, код и тесты.

## 📦 Фичи

- 🧠 Структура по задачам (`problems/<номер><Название>`)
- ✅ Тесты через `go test ./...`
- 🛠 Генерация задач скриптом
- 🔄 Автоматическая генерация таблицы задач в `README.md`
- 🧪 Автопроверки при коммите (`gofmt`, `go vet`, `README`)

## 🛠 Makefile команды

| Команда | Описание |
|--------|----------|
| `make new-solution number=123 title=MyTitle` | Создать новую задачу |
| `make update-readme` | Перегенерировать README.md |
| `make test` | Прогнать все тесты |
| `make check` | Проверка gofmt и go vet |
| `make install-hooks` | Установить pre-commit hook |

## 🔗 Задачи

| # | Название задачи | Уровень | Решение | Описание |
|---|------------------|---------|---------|----------|
| 1 | [62. Unique Paths](https://leetcode.com/problems/unique-paths/) | Medium | [Code](problems/62UniquePaths/solution.go) | [Explanation](problems/62UniquePaths/README.md) |
| 2 | [136. Single Number](https://leetcode.com/problems/single-number) | Easy | [Code](problems/136SingleNumber/solution.go) | [Explanation](problems/136SingleNumber/README.md) |
| 3 | [162. Find Peak Element](https://leetcode.com/problems/find-peak-element) | Medium | [Code](problems/162FindPeakElement/solution.go) | [Explanation](problems/162FindPeakElement/README.md) |
| 4 | [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view) | Medium | [Code](problems/199BinaryTreeRightSideView/solution.go) | [Explanation](problems/199BinaryTreeRightSideView/README.md) |
| 5 | [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree) | Medium | [Code](problems/208ImplementTriePrefixTree/solution.go) | [Explanation](problems/208ImplementTriePrefixTree/README.md) |
| 6 | [334. Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence) | Medium | [Code](problems/334IncreasingTripletSubsequence/solution.go) | [Explanation](problems/334IncreasingTripletSubsequence/README.md) |
| 7 | [338. Counting Bits](https://leetcode.com/problems/counting-bits/) | Easy | [Code](problems/338CountingBits/solution.go) | [Explanation](problems/338CountingBits/README.md) |
| 8 | [1071. Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/) | Easy | [Code](problems/1071GreatestCommonDivisorOfStrings/solution.go) | [Explanation](problems/1071GreatestCommonDivisorOfStrings/README.md) |
| 9 | [1207. Unique Number Of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/) | Easy | [Code](problems/1207UniqueNumberOfOccurrences/solution.go) | [Explanation](problems/1207UniqueNumberOfOccurrences/README.md) |
| 10 | [1431. Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) | Easy | [Code](problems/1431KidsWithTheGreatestNumberOfCandies/solution.go) | [Explanation](problems/1431KidsWithTheGreatestNumberOfCandies/README.md) |
| 11 | [1768. Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) | Easy | [Code](problems/1768MergeStringsAlternately/solution.go) | [Explanation](problems/1768MergeStringsAlternately/README.md) |
| 12 | [2215. Find The Difference Of Two Arrays](https://leetcode.com/problems/find-the-difference-of-two-arrays/) | Easy | [Code](problems/2215FindTheDifferenceOfTwoArrays/solution.go) | [Explanation](problems/2215FindTheDifferenceOfTwoArrays/README.md) |
| 13 | [2390. Removing Stars From a String](https://leetcode.com/problems/removing-stars-from-a-string) | Medium | [Code](problems/2390RemovingStarsFromAString/solution.go) | [Explanation](problems/2390RemovingStarsFromAString/README.md) |
