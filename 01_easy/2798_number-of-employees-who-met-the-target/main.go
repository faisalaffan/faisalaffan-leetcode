package main

// LeetCode #2798: Number of Employees Who Met the Target
// https://leetcode.com/problems/number-of-employees-who-met-the-target/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{5, 1, 4, 2, 2}, 6))
}

func NumberOfEmployeesWhoMetTheTarget(hours []int, target int) int {
	count := 0
	for _, h := range hours {
		if h >= target {
			count++
		}
	}
	return count
}
