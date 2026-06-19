package main

// LeetCode #495: Teemo Attacking
// https://leetcode.com/problems/teemo-attacking/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func TeemoAttacking(timeSeries []int, duration int) int {
	total := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		total += min(duration, timeSeries[i+1]-timeSeries[i])
	}
	if len(timeSeries) > 0 {
		total += duration
	}
	return total
}

func main() {
	fmt.Println(TeemoAttacking([]int{1, 4}, 2))
	fmt.Println(TeemoAttacking([]int{1, 2}, 2))
}
