package main

// LeetCode #1700: Number of Students Unable to Eat Lunch
// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountStudents(students []int, sandwiches []int) int {
	count := [2]int{0, 0}
	for _, s := range students {
		count[s]++
	}
	for _, sandwich := range sandwiches {
		if count[sandwich] == 0 {
			break
		}
		count[sandwich]--
	}
	return count[0] + count[1]
}

func main() {
	fmt.Println(CountStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(CountStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}
