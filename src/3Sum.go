package main

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i, val := range nums {
		if i > 0 && val == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := val + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{val, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return res
}
