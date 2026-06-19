package main

// LeetCode #2365: Task Scheduler II
// https://leetcode.com/problems/task-scheduler-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Track last completion day for each task type. If within cooldown, advance day.

import "fmt"

func main() {
	fmt.Println(taskSchedulerII([]int{1, 2, 1, 2, 3, 1}, 3)) // 9
	fmt.Println(taskSchedulerII([]int{5, 8, 8, 5}, 2))       // 6
}

func taskSchedulerII(tasks []int, space int) int64 {
	last := make(map[int]int64)
	var day int64 = 0
	for _, t := range tasks {
		day++
		if prev, ok := last[t]; ok && day-prev <= int64(space) {
			day = prev + int64(space) + 1
		}
		last[t] = day
	}
	return day
}
