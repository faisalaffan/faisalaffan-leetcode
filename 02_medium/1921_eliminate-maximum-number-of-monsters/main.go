package main

// LeetCode #1921: Eliminate Maximum Number of Monsters
// https://leetcode.com/problems/eliminate-maximum-number-of-monsters/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(EliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
	fmt.Println(EliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
	fmt.Println(EliminateMaximum([]int{3, 2, 4}, []int{5, 3, 2}))
}

// Time: O(n log n), Space: O(n)
func EliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	time := make([]int, n)
	for i := 0; i < n; i++ {
		time[i] = (dist[i] + speed[i] - 1) / speed[i] // ceil division
	}
	sort.Ints(time)

	for i := 0; i < n; i++ {
		if time[i] <= i {
			return i
		}
	}
	return n
}
