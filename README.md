# LeetCode Solutions in Go

🎯 Этот репозиторий содержит решения задач с [LeetCode](https://leetcode.com), написанные на Go. Каждое решение включает объяснение, код и тесты.

#### 📚 [Алгоритмическая шпаргалка](algoCheatsheet.md)

#### 📚 [Шпаргалка по Go](deepGo.md)

#### 📚 [Лайвкодинг кейсы](livecodingCases.md)

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
| 2 | [103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/) | Medium | Tree, Breadth-First Search, Binary Tree | [Code](problems/100-199/103BinaryTreeZigzagLevelOrderTraversal/solution.go) | [Explanation](problems/100-199/103BinaryTreeZigzagLevelOrderTraversal/README.md) |
| 3 | [136. Single Number](https://leetcode.com/problems/single-number) | Easy | Array, Bit Manipulation | [Code](problems/100-199/136SingleNumber/solution.go) | [Explanation](problems/100-199/136SingleNumber/README.md) |
| 4 | [162. Find Peak Element](https://leetcode.com/problems/find-peak-element) | Medium | Array, Binary Search | [Code](problems/100-199/162FindPeakElement/solution.go) | [Explanation](problems/100-199/162FindPeakElement/README.md) |
| 5 | [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view) | Medium | Tree, Depth-First Search, Breadth-First Search, Binary Tree | [Code](problems/100-199/199BinaryTreeRightSideView/solution.go) | [Explanation](problems/100-199/199BinaryTreeRightSideView/README.md) |
| 6 | [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree) | Medium | Hash Table, String, Design, Trie | [Code](problems/200-299/208ImplementTriePrefixTree/solution.go) | [Explanation](problems/200-299/208ImplementTriePrefixTree/README.md) |
| 7 | [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) | Medium | Array, Divide and Conquer, Sorting, Heap (Priority Queue), Quickselect | [Code](problems/200-299/215KthLargestElementInAnArray/solution.go) | [Explanation](problems/200-299/215KthLargestElementInAnArray/README.md) |
| 8 | [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) | Medium | Array, Prefix Sum | [Code](problems/200-299/238ProductOfArrayExceptSelf/solution.go) | [Explanation](problems/200-299/238ProductOfArrayExceptSelf/README.md) |
| 9 | [334. Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence) | Medium | Array, Greedy | [Code](problems/300-399/334IncreasingTripletSubsequence/solution.go) | [Explanation](problems/300-399/334IncreasingTripletSubsequence/README.md) |
| 10 | [338. Counting Bits](https://leetcode.com/problems/counting-bits/) | Easy | Dynamic Programming, Bit Manipulation | [Code](problems/300-399/338CountingBits/solution.go) | [Explanation](problems/300-399/338CountingBits/README.md) |
| 11 | [394. Decode String](https://leetcode.com/problems/decode-string/) | Medium | String, Stack, Recursion | [Code](problems/300-399/394DecodeString/solution.go) | [Explanation](problems/300-399/394DecodeString/README.md) |
| 12 | [680. Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/) | Easy | Two Pointers, String, Greedy | [Code](problems/600-699/680ValidPalindromeII/solution.go) | [Explanation](problems/600-699/680ValidPalindromeII/README.md) |
| 13 | [724. Find Pivot Index](https://leetcode.com/problems/find-pivot-index/) | Easy | Array, Prefix Sum | [Code](problems/700-799/724FindPivotIndex/solution.go) | [Explanation](problems/700-799/724FindPivotIndex/README.md) |
| 14 | [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) | Medium | Array, Breadth-First Search, Matrix | [Code](problems/900-999/994RottingOranges/solution.go) | [Explanation](problems/900-999/994RottingOranges/README.md) |
| 15 | [1071. Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/) | Easy | Math, String | [Code](problems/1000-1099/1071GreatestCommonDivisorOfStrings/solution.go) | [Explanation](problems/1000-1099/1071GreatestCommonDivisorOfStrings/README.md) |
| 16 | [1207. Unique Number Of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/) | Easy | Array, Hash Table | [Code](problems/1200-1299/1207UniqueNumberOfOccurrences/solution.go) | [Explanation](problems/1200-1299/1207UniqueNumberOfOccurrences/README.md) |
| 17 | [1431. Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) | Easy | Array | [Code](problems/1400-1499/1431KidsWithTheGreatestNumberOfCandies/solution.go) | [Explanation](problems/1400-1499/1431KidsWithTheGreatestNumberOfCandies/README.md) |
| 18 | [1493. Longest Subarray of 1's After Deleting One Element](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/) | Medium | Array, Dynamic Programming, Sliding Window | [Code](problems/1400-1499/1493LongestSubarrayOf1sAfterDeletingOneElement/solution.go) | [Explanation](problems/1400-1499/1493LongestSubarrayOf1sAfterDeletingOneElement/README.md) |
| 19 | [1768. Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) | Easy | Two Pointers, String | [Code](problems/1700-1799/1768MergeStringsAlternately/solution.go) | [Explanation](problems/1700-1799/1768MergeStringsAlternately/README.md) |
| 20 | [2215. Find The Difference Of Two Arrays](https://leetcode.com/problems/find-the-difference-of-two-arrays/) | Easy | Array, Hash Table | [Code](problems/2200-2299/2215FindTheDifferenceOfTwoArrays/solution.go) | [Explanation](problems/2200-2299/2215FindTheDifferenceOfTwoArrays/README.md) |
| 21 | [2390. Removing Stars From a String](https://leetcode.com/problems/removing-stars-from-a-string) | Medium | String, Stack, Simulation | [Code](problems/2300-2399/2390RemovingStarsFromAString/solution.go) | [Explanation](problems/2300-2399/2390RemovingStarsFromAString/README.md) |
