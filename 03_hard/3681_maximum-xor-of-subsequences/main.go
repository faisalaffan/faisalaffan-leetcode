package main

// LeetCode #3681: Maximum XOR of Subsequences
// https://leetcode.com/problems/maximum-xor-of-subsequences/
// Difficulty: Hard
//
// Given array nums, select two subsequences preserving order. Let X be XOR
// of first and Y of second. Maximize X XOR Y.
//
// Approach: Build linear basis of all numbers. Max XOR of two subsequences
// equals max XOR achievable from the linear basis (since we can assign
// any subset's XOR to one subsequence and the rest to the other).

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxXorSubsequences([]int{1, 2, 3}))
	// Example 2
	fmt.Println(maxXorSubsequences([]int{5, 2}))
	// Edge: single element
	fmt.Println(maxXorSubsequences([]int{7}))
	// Edge: all zeros
	fmt.Println(maxXorSubsequences([]int{0, 0, 0}))
}

func maxXorSubsequences(nums []int) int {
	// Build linear basis
	basis := make([]int, 0, 31)
	for _, v := range nums {
		x := v
		for _, b := range basis {
			x = min(x, x^b)
		}
		if x > 0 {
			// Insert x maintaining basis property
			basis = append(basis, x)
			// Sort descending
			for i := len(basis) - 1; i > 0; i-- {
				if basis[i] > basis[i-1] {
					basis[i], basis[i-1] = basis[i-1], basis[i]
				}
			}
		}
	}

	// Maximum XOR from basis
	// We want max xor of X and Y where X and Y are XORs of two disjoint
	// sets. Since X XOR Y = XOR of all elements in X union Y (but each
	// element appears once in either X or Y, not both).
	// Actually, since we can assign any element to either subsequence,
	// X XOR Y can be any XOR of a subset (where elements assigned to
	// second subsequence contribute their XOR if count is odd... wait)
	//
	// Actually, X = XOR of subset A, Y = XOR of subset B, A∩B=∅, A∪B⊆nums
	// X XOR Y = XOR of elements that appear in an odd number of {A, B}
	// = XOR of elements in exactly one of A or B.
	// = XOR of elements in A Δ B (symmetric difference)
	// Since A and B are arbitrary disjoint subsets, A Δ B is any subset.
	// So X XOR Y = XOR of any subset of nums.
	// Max XOR of any subset = max XOR from linear basis.

	result := 0
	for _, b := range basis {
		if result^b > result {
			result ^= b
		}
	}
	return result
}
