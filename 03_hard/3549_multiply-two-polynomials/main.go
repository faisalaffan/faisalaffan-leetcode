package main

// LeetCode #3549: Multiply Two Polynomials
// https://leetcode.com/problems/multiply-two-polynomials/
// Difficulty: Hard [Paid]
//
// Given two polynomials represented as arrays of coefficients (index = power),
// return their product as an array of coefficients.
//
// Approach: Standard polynomial multiplication O(n*m).

import "fmt"

func main() {
	// Example 1
	fmt.Println(multiply([]int{1, 2, 3}, []int{4, 5}))
	// Example 2
	fmt.Println(multiply([]int{1, 1}, []int{1, 1}))
	// Example 3: constant polynomial
	fmt.Println(multiply([]int{2}, []int{3, 4}))
	// Edge: with zeros
	fmt.Println(multiply([]int{0, 1}, []int{1, 0, 1}))
}

func multiply(poly1 []int, poly2 []int) []int {
	if len(poly1) == 0 || len(poly2) == 0 {
		return []int{}
	}

	result := make([]int, len(poly1)+len(poly2)-1)
	for i, c1 := range poly1 {
		for j, c2 := range poly2 {
			result[i+j] += c1 * c2
		}
	}
	return result
}
