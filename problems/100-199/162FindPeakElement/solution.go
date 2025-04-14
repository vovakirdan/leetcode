package findpeakelement

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid+1] {
			right = mid // пик слева
		} else {
			left = mid + 1 // пик справа
		}
	}

	return left
}
