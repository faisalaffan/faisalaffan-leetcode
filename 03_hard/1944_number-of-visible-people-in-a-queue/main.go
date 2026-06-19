package main

// LeetCode #1944: Number of Visible People in a Queue
// https://leetcode.com/problems/number-of-visible-people-in-a-queue/
// Difficulty: Hard
// Monotonic stack from right. Count shorter people on stack, then +1 for the next taller.

import "fmt"

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	stack := make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		visible := 0
		// Pop everyone shorter than current height
		for len(stack) > 0 && stack[len(stack)-1] < heights[i] {
			stack = stack[:len(stack)-1]
			visible++
		}
		// If there's someone taller left, they are also visible
		if len(stack) > 0 {
			visible++
		}
		ans[i] = visible
		stack = append(stack, heights[i])
	}

	return ans
}

func main() {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9})) // Expected: [3 1 2 1 1 0]
	fmt.Println(canSeePersonsCount([]int{5, 1, 2, 3, 10}))     // Expected: [4 1 1 1 0]
}
