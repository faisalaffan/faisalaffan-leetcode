package main

// LeetCode #177: Nth Highest Salary
// https://leetcode.com/problems/nth-highest-salary/
// Difficulty: Medium
// Time: O(n log n) average, Space: O(n) for quickselect

import (
	"fmt"
	"sort"
)

func nthHighestSalary(salaries []int, n int) int {
	if n <= 0 || n > len(salaries) {
		return 0
	}

	// Use sort + deduplicate
	seen := make(map[int]bool)
	unique := []int{}
	for _, s := range salaries {
		if !seen[s] {
			seen[s] = true
			unique = append(unique, s)
		}
	}

	if n > len(unique) {
		return 0
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	return unique[n-1]
}

func main() {
	fmt.Println(nthHighestSalary([]int{100, 200, 300, 200}, 2))
	fmt.Println(nthHighestSalary([]int{100, 100}, 2))
	fmt.Println(nthHighestSalary([]int{60, 70, 80, 90, 100}, 3))
}
