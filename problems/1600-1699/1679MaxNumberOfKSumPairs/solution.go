package maxnumberofksumpairs

import (
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	count := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			count++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}
	return count
}
