package main

// LeetCode #1130: Minimum Cost Tree From Leaf Values
// https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/
// Difficulty: Medium
//
// Approach: Monotonic decreasing stack (greedy)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))    // 32
	fmt.Println(mctFromLeafValues([]int{4, 11}))      // 44
}

func mctFromLeafValues(arr []int) int {
	stack := make([]int, 0)
	result := 0

	for _, v := range arr {
		for len(stack) > 0 && stack[len(stack)-1] <= v {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				result += mid * v
			} else {
				if stack[len(stack)-1] < v {
					result += mid * stack[len(stack)-1]
				} else {
					result += mid * v
				}
			}
		}
		stack = append(stack, v)
	}

	for len(stack) > 1 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result += last * stack[len(stack)-1]
	}

	return result
}
