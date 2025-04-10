package singlenumber

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // XOR: removes duplicates, leaves the single number
	}
	return result
}
