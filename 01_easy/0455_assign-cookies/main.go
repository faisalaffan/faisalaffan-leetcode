package main

// LeetCode #455: Assign Cookies
// https://leetcode.com/problems/assign-cookies/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n + m log m), Space: O(1)
func AssignCookies(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func main() {
	fmt.Println(AssignCookies([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(AssignCookies([]int{1, 2}, []int{1, 2, 3}))
}
