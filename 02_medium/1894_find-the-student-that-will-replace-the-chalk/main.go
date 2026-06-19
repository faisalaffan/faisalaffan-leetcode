package main

// LeetCode #1894: Find the Student that Will Replace the Chalk
// https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ChalkReplacer([]int{5, 1, 5}, 22))
	fmt.Println(ChalkReplacer([]int{3, 4, 1, 2}, 25))
	fmt.Println(ChalkReplacer([]int{5, 2, 3}, 9))
}

// Time: O(n), Space: O(1)
func ChalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, c := range chalk {
		sum += c
	}
	k %= sum

	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return 0
}
