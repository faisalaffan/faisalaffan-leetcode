package main

// LeetCode #176: Second Highest Salary
// https://leetcode.com/problems/second-highest-salary/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

import "math"

func secondHighestSalary(salaries []int) int {
	first, second := math.MinInt32, math.MinInt32

	for _, s := range salaries {
		if s > first {
			second = first
			first = s
		} else if s > second && s < first {
			second = s
		}
	}

	if second == math.MinInt32 {
		return 0
	}
	return second
}

func main() {
	fmt.Println(secondHighestSalary([]int{100, 200, 300}))
	fmt.Println(secondHighestSalary([]int{100, 100}))
	fmt.Println(secondHighestSalary([]int{50, 100, 100, 75}))
}
