package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	seen := make(map[int]bool)

	for _, num := range arr {
		freq[num]++
	}

	for _, count := range freq {
		if seen[count] {
			return false
		}
		seen[count] = true
	}

	return true
}
