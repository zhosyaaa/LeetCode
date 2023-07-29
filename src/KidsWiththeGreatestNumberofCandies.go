package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, 0, len(candies))
	max := candies[0]
	for i := 0; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies > max {
			ans = append(ans, true)
			continue
		}
		ans = append(ans, false)
	}
	return ans
}
