package findpivotindex

func sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func pivotIndex(nums []int) int {
	total := sum(nums)
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum := total - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
