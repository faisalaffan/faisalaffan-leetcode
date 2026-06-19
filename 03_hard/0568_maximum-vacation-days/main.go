package main

// LeetCode #568: Maximum Vacation Days
// https://leetcode.com/problems/maximum-vacation-days/
// Difficulty: Hard

import (
	"fmt"
)

func main() {
	flights := [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}
	days := [][]int{
		{1, 3, 1},
		{6, 0, 3},
		{3, 3, 3},
	}
	fmt.Println(maxVacationDays(flights, days)) // Expected: 12
}

func maxVacationDays(flights [][]int, days [][]int) int {
	n := len(flights)   // cities
	k := len(days[0])   // weeks

	// prev[j] = max vacation days ending at city j for current week
	prev := make([]int, n)
	for j := 0; j < n; j++ {
		// Week 0: can we reach city j?
		if j == 0 || flights[0][j] == 1 {
			prev[j] = days[j][0]
		} else {
			prev[j] = -1
		}
	}

	for w := 1; w < k; w++ {
		cur := make([]int, n)
		for j := 0; j < n; j++ {
			cur[j] = -1
		}
		for j := 0; j < n; j++ {
			for i := 0; i < n; i++ {
				if prev[i] >= 0 && (i == j || flights[i][j] == 1) {
					if prev[i]+days[j][w] > cur[j] {
						cur[j] = prev[i] + days[j][w]
					}
				}
			}
		}
		prev = cur
	}

	ans := 0
	for j := 0; j < n; j++ {
		if prev[j] > ans {
			ans = prev[j]
		}
	}
	return ans
}
