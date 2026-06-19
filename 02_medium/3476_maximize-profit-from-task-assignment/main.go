package main

// LeetCode #3476: Maximize Profit from Task Assignment
// https://leetcode.com/problems/maximize-profit-from-task-assignment/
// Difficulty: Medium [Paid]
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximizeProfitFromTaskAssignment([]int{1, 3, 2, 4}, []int{2, 5, 3, 6}))
	// Test case 2
	fmt.Println("Test 2:", MaximizeProfitFromTaskAssignment([]int{5, 1, 3}, []int{10, 2, 5}))
	// Test case 3
	fmt.Println("Test 3:", MaximizeProfitFromTaskAssignment([]int{2}, []int{4}))
}

func MaximizeProfitFromTaskAssignment(difficulty []int, profit []int) int {
	type task struct {
		d int
		p int
	}
	n := len(difficulty)
	tasks := make([]task, n)
	for i := 0; i < n; i++ {
		tasks[i] = task{difficulty[i], profit[i]}
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].d < tasks[j].d || (tasks[i].d == tasks[j].d && tasks[i].p > tasks[j].p)
	})

	maxProfit := 0
	best := 0
	for i := 0; i < n; i++ {
		if tasks[i].p > best {
			best = tasks[i].p
		}
		maxProfit += best
	}
	return maxProfit
}
