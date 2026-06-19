package main

// LeetCode #2055: Plates Between Candles
// https://leetcode.com/problems/plates-between-candles/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	// Prefix sum of plates
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		if s[i] == '*' {
			prefix[i+1]++
		}
	}

	// Nearest candle to the left
	leftCandle := make([]int, n)
	last := -1
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			last = i
		}
		leftCandle[i] = last
	}

	// Nearest candle to the right
	rightCandle := make([]int, n)
	last = -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			last = i
		}
		rightCandle[i] = last
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		left, right := q[0], q[1]
		l := rightCandle[left]
		r := leftCandle[right]
		if l == -1 || r == -1 || l >= r {
			result[i] = 0
		} else {
			result[i] = prefix[r] - prefix[l]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
	// Expected: [2, 3]

	// Test case 2
	fmt.Println("Test 2:", platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}))
	// Expected: [9, 0, 0, 0, 0]

	// Test case 3
	fmt.Println("Test 3:", platesBetweenCandles("|*|", [][]int{{0, 2}}))
	// Expected: [1]
}
