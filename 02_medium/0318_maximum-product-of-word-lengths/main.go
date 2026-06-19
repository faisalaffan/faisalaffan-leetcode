package main

// LeetCode #318: Maximum Product of Word Lengths
// https://leetcode.com/problems/maximum-product-of-word-lengths/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func maxProduct(words []string) int {
	bits := make([]int, len(words))
	for i, w := range words {
		mask := 0
		for _, ch := range w {
			mask |= 1 << (ch - 'a')
		}
		bits[i] = mask
	}

	maxProd := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bits[i]&bits[j] == 0 {
				prod := len(words[i]) * len(words[j])
				if prod > maxProd {
					maxProd = prod
				}
			}
		}
	}
	return maxProd
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	// Expected: 16 (abcw * xtfn)

	// Test case 2
	fmt.Println("Test 2:", maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	// Expected: 4 (ab * cd)

	// Test case 3
	fmt.Println("Test 3:", maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
	// Expected: 0
}
