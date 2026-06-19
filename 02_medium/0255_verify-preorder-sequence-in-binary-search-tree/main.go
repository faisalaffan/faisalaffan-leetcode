package main

// LeetCode #255: Verify Preorder Sequence in Binary Search Tree
// https://leetcode.com/problems/verify-preorder-sequence-in-binary-search-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import "fmt"

func verifyPreorder(preorder []int) bool {
	stack := []int{}
	lower := ^int(^uint(0) >> 1) // math.MinInt

	for _, val := range preorder {
		if val < lower {
			return false
		}
		for len(stack) > 0 && val > stack[len(stack)-1] {
			lower = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, val)
	}

	return true
}

func main() {
	fmt.Println(verifyPreorder([]int{5, 2, 1, 3, 6}))
	fmt.Println(verifyPreorder([]int{5, 2, 6, 1, 3}))
	fmt.Println(verifyPreorder([]int{1, 2, 3}))
}
