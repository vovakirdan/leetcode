package kidswiththegreatestnumberofcandies

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxNum := slices.Max(candies)
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		res[i] = candies[i]+extraCandies >= maxNum
	}
	return res
}
