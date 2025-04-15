# LeetCode Solutions in Go

🎯 Этот репозиторий содержит решения задач с [LeetCode](https://leetcode.com), написанные на Go. Каждое решение включает объяснение, код и тесты.

## 📦 Фичи

- 🧠 Задачи структурированы в папках `problems/<range>/<номер><Название>`
- ✅ Тесты запускаются через `go test ./...`
- 🔄 README обновляется автоматически `make update-readme`
- 🏷 Теги загружаются через `make fetch-tags`
- ⚙️ Хуки при коммите: `gofmt`, `go vet`, генерация README

## 🛠 Makefile команды

| Команда | Описание |
|--------|----------|
| `make new number=123 title=MyTitle level=Easy` | Создать новую задачу |
| `make update-readme` | Перегенерировать README.md |
| `make test` | Прогнать все тесты |
| `make check` | Проверка gofmt и go vet |
| `make install-hooks` | Установить pre-commit hook |
| `make fetch-tags` | Загрузить теги в README каждой задачи |

## 🔗 Задачи

| # | Название | Уровень | Теги | Решение | Описание |
|---|----------|---------|------|---------|----------|
| 1 | [62. Unique Paths](https://leetcode.com/problems/unique-paths/) | Medium | Math, Dynamic Programming, Combinatorics | [Code](problems/0-99/62UniquePaths/solution.go) | [Explanation](problems/0-99/62UniquePaths/README.md) |
| 2 | [136. Single Number](https://leetcode.com/problems/single-number) | Easy | Array, Bit Manipulation | [Code](problems/100-199/136SingleNumber/solution.go) | [Explanation](problems/100-199/136SingleNumber/README.md) |
| 3 | [162. Find Peak Element](https://leetcode.com/problems/find-peak-element) | Medium | Array, Binary Search | [Code](problems/100-199/162FindPeakElement/solution.go) | [Explanation](problems/100-199/162FindPeakElement/README.md) |
| 4 | [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view) | Medium | Tree, Depth-First Search, Breadth-First Search, Binary Tree | [Code](problems/100-199/199BinaryTreeRightSideView/solution.go) | [Explanation](problems/100-199/199BinaryTreeRightSideView/README.md) |
| 5 | [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree) | Medium | Hash Table, String, Design, Trie | [Code](problems/200-299/208ImplementTriePrefixTree/solution.go) | [Explanation](problems/200-299/208ImplementTriePrefixTree/README.md) |
| 6 | [238. Product Of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) | Medium | Array, Prefix Sum | [Code](problems/200-299/238ProductOfArrayExceptSelf/solution.go) | [Explanation](problems/200-299/238ProductOfArrayExceptSelf/README.md) |
| 7 | [334. Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence) | Medium | Array, Greedy | [Code](problems/300-399/334IncreasingTripletSubsequence/solution.go) | [Explanation](problems/300-399/334IncreasingTripletSubsequence/README.md) |
| 8 | [338. Counting Bits](https://leetcode.com/problems/counting-bits/) | Easy | Dynamic Programming, Bit Manipulation | [Code](problems/300-399/338CountingBits/solution.go) | [Explanation](problems/300-399/338CountingBits/README.md) |
| 9 | [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) | Medium | Array, Breadth-First Search, Matrix | [Code](problems/900-999/994RottingOranges/solution.go) | [Explanation](problems/900-999/994RottingOranges/README.md) |
| 10 | [1071. Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/) | Easy | Math, String | [Code](problems/1000-1099/1071GreatestCommonDivisorOfStrings/solution.go) | [Explanation](problems/1000-1099/1071GreatestCommonDivisorOfStrings/README.md) |
| 11 | [1207. Unique Number Of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/) | Easy | Array, Hash Table | [Code](problems/1200-1299/1207UniqueNumberOfOccurrences/solution.go) | [Explanation](problems/1200-1299/1207UniqueNumberOfOccurrences/README.md) |
| 12 | [1431. Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) | Easy | Array | [Code](problems/1400-1499/1431KidsWithTheGreatestNumberOfCandies/solution.go) | [Explanation](problems/1400-1499/1431KidsWithTheGreatestNumberOfCandies/README.md) |
| 13 | [1768. Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) | Easy | Two Pointers, String | [Code](problems/1700-1799/1768MergeStringsAlternately/solution.go) | [Explanation](problems/1700-1799/1768MergeStringsAlternately/README.md) |
| 14 | [2215. Find The Difference Of Two Arrays](https://leetcode.com/problems/find-the-difference-of-two-arrays/) | Easy | Array, Hash Table | [Code](problems/2200-2299/2215FindTheDifferenceOfTwoArrays/solution.go) | [Explanation](problems/2200-2299/2215FindTheDifferenceOfTwoArrays/README.md) |
| 15 | [2390. Removing Stars From a String](https://leetcode.com/problems/removing-stars-from-a-string) | Medium | String, Stack, Simulation | [Code](problems/2300-2399/2390RemovingStarsFromAString/solution.go) | [Explanation](problems/2300-2399/2390RemovingStarsFromAString/README.md) |
