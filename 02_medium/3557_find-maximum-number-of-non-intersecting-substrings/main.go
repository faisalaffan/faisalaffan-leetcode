package main

// LeetCode #3557: Find Maximum Number of Non Intersecting Substrings
// https://leetcode.com/problems/find-maximum-number-of-non-intersecting-substrings/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindMaximumNumberOfNonIntersectingSubstrings("abcba"))
	// Test case 2
	fmt.Println("Test 2:", FindMaximumNumberOfNonIntersectingSubstrings("abac"))
	// Test case 3
	fmt.Println("Test 3:", FindMaximumNumberOfNonIntersectingSubstrings("a"))
}

func FindMaximumNumberOfNonIntersectingSubstrings(s string) int {
	n := len(s)
	count := 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		// Try to find a substring starting at i that doesn't intersect with used ones
		for j := i + 1; j <= n; j++ {
			overlap := false
			for k := i; k < j; k++ {
				if used[k] {
					overlap = true
					break
				}
			}
			if !overlap {
				// Check if substring s[i:j] is valid (e.g., palindrome check)
				isPal := true
				for l, r := i, j-1; l < r; l, r = l+1, r-1 {
					if s[l] != s[r] {
						isPal = false
						break
					}
				}
				if isPal {
					for k := i; k < j; k++ {
						used[k] = true
					}
					count++
					break
				}
			}
		}
	}
	return count
}
