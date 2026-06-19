package main

// LeetCode #2375: Construct Smallest Number From DI String
// https://leetcode.com/problems/construct-smallest-number-from-di-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Use stack: push numbers 1..n+1. On 'I' or end, pop stack to build result.

import "fmt"

func main() {
	fmt.Println(smallestNumber("II"))  // "123"
	fmt.Println(smallestNumber("DI"))  // "231"
	fmt.Println(smallestNumber("DDD")) // "4321"
}

func smallestNumber(pattern string) string {
	n := len(pattern)
	stack := make([]int, 0, n+1)
	res := make([]byte, 0, n+1)
	for i := 0; i <= n; i++ {
		stack = append(stack, i+1)
		if i == n || pattern[i] == 'I' {
			for len(stack) > 0 {
				res = append(res, byte('0'+stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(res)
}
