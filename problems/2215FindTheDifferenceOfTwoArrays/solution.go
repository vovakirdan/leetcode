package find_the_difference_of_two_arrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]bool{}
	set2 := map[int]bool{}

	for _, elem := range nums1 {
		set1[elem] = true
	}

	for _, elem := range nums2 {
		set2[elem] = true
	}

	res1 := []int{}
	res2 := []int{}

	for num := range set1 {
		if !set2[num] {
			res1 = append(res1, num)
		}
	}

	for num := range set2 {
		if !set1[num] {
			res2 = append(res2, num)
		}
	}

	return [][]int{res1, res2}
}
