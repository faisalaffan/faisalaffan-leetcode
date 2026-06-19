package main

// LeetCode #2197: Replace Non-Coprime Numbers in Array
// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/
// Difficulty: Hard
//
// Stack merge: for each number, push onto stack. While stack has >= 2 elements
// and gcd(stack[-2], stack[-1]) > 1, pop top two, push lcm.

import (
	"fmt"
)

func main() {
	// Example: [6,4,3,2,7,6,2] => [12,7,6]
	fmt.Println(replaceNonCoprimeNumbers([]int{6, 4, 3, 2, 7, 6, 2}))
	// Example: [2,2,1,1,3,3,3] => [2,1,1,3]
	fmt.Println(replaceNonCoprimeNumbers([]int{2, 2, 1, 1, 3, 3, 3}))
	// Example: [1,1,1,1] => [1,1,1,1]
	fmt.Println(replaceNonCoprimeNumbers([]int{1, 1, 1, 1}))
	// Example: [12,18,6] => [36]
	fmt.Println(replaceNonCoprimeNumbers([]int{12, 18, 6}))
	// Example: single
	fmt.Println(replaceNonCoprimeNumbers([]int{7}))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func replaceNonCoprimeNumbers(nums []int) []int {
	stack := make([]int, 0, len(nums))

	for _, x := range nums {
		stack = append(stack, x)
		for len(stack) >= 2 {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			g := gcd(a, b)
			if g == 1 {
				break
			}
			// Replace with LCM
			stack = stack[:len(stack)-2]
			stack = append(stack, lcm(a, b))
		}
	}

	return stack
}
