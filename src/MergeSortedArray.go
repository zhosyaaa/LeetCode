package main

import (
	"sort"
	"strings"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// Valid Palindrome
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !isValidChar(s[left]) {
			left++
		}
		for left < right && !isValidChar(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isValidChar(ch uint8) bool {
	return (ch >= 65 && ch <= 90) || (ch >= 48 && ch <= 57) || (ch >= 97 && ch <= 122)
}
