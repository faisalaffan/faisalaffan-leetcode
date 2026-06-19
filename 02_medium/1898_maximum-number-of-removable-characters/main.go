package main

// LeetCode #1898: Maximum Number of Removable Characters
// https://leetcode.com/problems/maximum-number-of-removable-characters/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxRemovals("abcacb", "ab", []int{3, 1, 0}))
	fmt.Println(MaxRemovals("abcbddddd", "abcd", []int{3, 2, 1, 4, 5, 6}))
	fmt.Println(MaxRemovals("abcab", "abc", []int{0, 1, 2, 3, 4}))
}

// Time: O((n+m) log k) where n = len(s), m = len(p), k = len(removable)
// Space: O(n)
func MaxRemovals(s string, p string, removable []int) int {
	left, right := 0, len(removable)
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if canForm(s, p, removable, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func canForm(s string, p string, removable []int, k int) bool {
	removed := make([]bool, len(s))
	for i := 0; i < k; i++ {
		removed[removable[i]] = true
	}

	j := 0
	for i := 0; i < len(s) && j < len(p); i++ {
		if !removed[i] && s[i] == p[j] {
			j++
		}
	}
	return j == len(p)
}
