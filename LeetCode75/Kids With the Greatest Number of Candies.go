package leetcode75

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxNumber := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > maxNumber {
			maxNumber = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		res[i] = candies[i]+extraCandies >= maxNumber
	}
	return res
}
