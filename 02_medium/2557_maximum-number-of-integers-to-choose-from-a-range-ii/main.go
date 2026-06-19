package main

// LeetCode #2557: Maximum Number of Integers to Choose From a Range II
// https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxCount(banned []int, n int, maxSum int) int {
	sort.Ints(banned)
	bannedSet := make(map[int]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	count := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if bannedSet[i] {
			continue
		}
		if sum+i > maxSum {
			break
		}
		sum += i
		count++
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxCount([]int{1, 6, 5}, 5, 6))
	// Expected: 2

	// Test case 2: large banned set
	fmt.Println("Test 2:", maxCount([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15, 40))
	// Expected: 3 (pick 11,12,13)

	// Test case 3
	fmt.Println("Test 3:", maxCount([]int{}, 5, 15))
	// Expected: 5 (1+2+3+4+5=15)
}
