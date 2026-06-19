package main

// LeetCode #2797: Partial Function with Placeholders
// https://leetcode.com/problems/partial-function-with-placeholders/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JS problem, adapted to Go. Partially applies arguments.

import "fmt"

func main() {
	add := func(a, b, c int) int { return a + b + c }
	partial := PartialFunctionWithPlaceholders(add, 1, nil, 3)
	fmt.Println(partial(2))
}

func PartialFunctionWithPlaceholders(fn func(int, int, int) int, args ...interface{}) func(int) int {
	return func(x int) int {
		realArgs := [3]int{}
		argIdx := 0
		for i, a := range args {
			if a == nil {
				realArgs[i] = x
			} else {
				realArgs[i] = a.(int)
				argIdx++
			}
		}
		return fn(realArgs[0], realArgs[1], realArgs[2])
	}
}
