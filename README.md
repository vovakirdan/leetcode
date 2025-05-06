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
| 4 | [151. Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/) | Medium | Two Pointers, String | [Code](problems/100-199/151ReverseWordsInAString/solution.go) | [Explanation](problems/100-199/151ReverseWordsInAString/README.md) |
| 5 | [162. Find Peak Element](https://leetcode.com/problems/find-peak-element) | Medium | Array, Binary Search | [Code](problems/100-199/162FindPeakElement/solution.go) | [Explanation](problems/100-199/162FindPeakElement/README.md) |
| 6 | [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view) | Medium | Tree, Depth-First Search, Breadth-First Search, Binary Tree | [Code](problems/100-199/199BinaryTreeRightSideView/solution.go) | [Explanation](problems/100-199/199BinaryTreeRightSideView/README.md) |
| 7 | [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree) | Medium | Hash Table, String, Design, Trie | [Code](problems/200-299/208ImplementTriePrefixTree/solution.go) | [Explanation](problems/200-299/208ImplementTriePrefixTree/README.md) |
| 8 | [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) | Medium | Array, Divide and Conquer, Sorting, Heap (Priority Queue), Quickselect | [Code](problems/200-299/215KthLargestElementInAnArray/solution.go) | [Explanation](problems/200-299/215KthLargestElementInAnArray/README.md) |
| 9 | [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) | Medium | Array, Prefix Sum | [Code](problems/200-299/238ProductOfArrayExceptSelf/solution.go) | [Explanation](problems/200-299/238ProductOfArrayExceptSelf/README.md) |
| 10 | [328. Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/) | Medium | Linked List | [Code](problems/300-399/328OddEvenLinkedList/solution.go) | [Explanation](problems/300-399/328OddEvenLinkedList/README.md) |
| 11 | [334. Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence) | Medium | Array, Greedy | [Code](problems/300-399/334IncreasingTripletSubsequence/solution.go) | [Explanation](problems/300-399/334IncreasingTripletSubsequence/README.md) |
| 12 | [338. Counting Bits](https://leetcode.com/problems/counting-bits/) | Easy | Dynamic Programming, Bit Manipulation | [Code](problems/300-399/338CountingBits/solution.go) | [Explanation](problems/300-399/338CountingBits/README.md) |
| 13 | [345. Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/) | Easy | Two Pointers, String | [Code](problems/300-399/345ReverseVowelsOfAString/solution.go) | [Explanation](problems/300-399/345ReverseVowelsOfAString/README.md) |
| 14 | [374. Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower/) | Easy | Binary Search, Interactive | [Code](problems/300-399/374GuessNumberHigherOrLower/solution.go) | [Explanation](problems/300-399/374GuessNumberHigherOrLower/README.md) |
| 15 | [392. Is Subsequence](https://leetcode.com/problems/is-subsequence/) | Easy | Two Pointers, String, Dynamic Programming | [Code](problems/300-399/392IsSubsequence/solution.go) | [Explanation](problems/300-399/392IsSubsequence/README.md) |
| 16 | [394. Decode String](https://leetcode.com/problems/decode-string/) | Medium | String, Stack, Recursion | [Code](problems/300-399/394DecodeString/solution.go) | [Explanation](problems/300-399/394DecodeString/README.md) |
| 17 | [399. Evaluate Division](https://leetcode.com/problems/evaluate-division/) | Medium | Array, String, Depth-First Search, Breadth-First Search, Union Find, Graph, Shortest Path | [Code](problems/300-399/399EvaluateDivision/solution.go) | [Explanation](problems/300-399/399EvaluateDivision/README.md) |
| 18 | [443. String Compression](https://leetcode.com/problems/string-compression/) | Medium | Two Pointers, String | [Code](problems/400-499/443StringCompression/solution.go) | [Explanation](problems/400-499/443StringCompression/README.md) |
| 19 | [450. Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/) | Medium | Tree, Binary Search Tree, Binary Tree | [Code](problems/400-499/450DeleteNodeInABst/solution.go) | [Explanation](problems/400-499/450DeleteNodeInABst/README.md) |
| 20 | [605. Can Place Flowers](https://leetcode.com/problems/can-place-flowers/) | Easy | Array, Greedy | [Code](problems/600-699/605CanPlaceFlowers/solution.go) | [Explanation](problems/600-699/605CanPlaceFlowers/README.md) |
| 21 | [680. Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/) | Easy | Two Pointers, String, Greedy | [Code](problems/600-699/680ValidPalindromeII/solution.go) | [Explanation](problems/600-699/680ValidPalindromeII/README.md) |
| 22 | [724. Find Pivot Index](https://leetcode.com/problems/find-pivot-index/) | Easy | Array, Prefix Sum | [Code](problems/700-799/724FindPivotIndex/solution.go) | [Explanation](problems/700-799/724FindPivotIndex/README.md) |
| 23 | [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) | Medium | Array, Breadth-First Search, Matrix | [Code](problems/900-999/994RottingOranges/solution.go) | [Explanation](problems/900-999/994RottingOranges/README.md) |
| 24 | [1071. Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/) | Easy | Math, String | [Code](problems/1000-1099/1071GreatestCommonDivisorOfStrings/solution.go) | [Explanation](problems/1000-1099/1071GreatestCommonDivisorOfStrings/README.md) |
| 25 | [1207. Unique Number Of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/) | Easy | Array, Hash Table | [Code](problems/1200-1299/1207UniqueNumberOfOccurrences/solution.go) | [Explanation](problems/1200-1299/1207UniqueNumberOfOccurrences/README.md) |
| 26 | [1431. Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) | Easy | Array | [Code](problems/1400-1499/1431KidsWithTheGreatestNumberOfCandies/solution.go) | [Explanation](problems/1400-1499/1431KidsWithTheGreatestNumberOfCandies/README.md) |
| 27 | [1493. Longest Subarray of 1's After Deleting One Element](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/) | Medium | Array, Dynamic Programming, Sliding Window | [Code](problems/1400-1499/1493LongestSubarrayOf1sAfterDeletingOneElement/solution.go) | [Explanation](problems/1400-1499/1493LongestSubarrayOf1sAfterDeletingOneElement/README.md) |
| 28 | [1679. Max Number of K-Sum Pairs](https://leetcode.com/problems/max-number-of-k-sum-pairs/) | Medium | Array, Hash Table, Two Pointers, Sorting | [Code](problems/1600-1699/1679MaxNumberOfKSumPairs/solution.go) | [Explanation](problems/1600-1699/1679MaxNumberOfKSumPairs/README.md) |
| 29 | [1768. Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) | Easy | Two Pointers, String | [Code](problems/1700-1799/1768MergeStringsAlternately/solution.go) | [Explanation](problems/1700-1799/1768MergeStringsAlternately/README.md) |
| 30 | [2215. Find The Difference Of Two Arrays](https://leetcode.com/problems/find-the-difference-of-two-arrays/) | Easy | Array, Hash Table | [Code](problems/2200-2299/2215FindTheDifferenceOfTwoArrays/solution.go) | [Explanation](problems/2200-2299/2215FindTheDifferenceOfTwoArrays/README.md) |
| 31 | [2390. Removing Stars From a String](https://leetcode.com/problems/removing-stars-from-a-string) | Medium | String, Stack, Simulation | [Code](problems/2300-2399/2390RemovingStarsFromAString/solution.go) | [Explanation](problems/2300-2399/2390RemovingStarsFromAString/README.md) |
| 32 | [2542. Maximum Subsequence Score](https://leetcode.com/problems/maximum-subsequence-score/) | Medium | Array, Greedy, Sorting, Heap (Priority Queue) | [Code](problems/2500-2599/2542MaximumSubsequenceScore/solution.go) | [Explanation](problems/2500-2599/2542MaximumSubsequenceScore/README.md) |
