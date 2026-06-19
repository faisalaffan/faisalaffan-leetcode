package main

// LeetCode #3160: Find the Number of Distinct Colors Among the Balls
// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func queryResults(limit int, queries [][]int) []int {
	ballColor := make(map[int]int)
	colorCount := make(map[int]int)
	ans := make([]int, len(queries))

	for i, q := range queries {
		ball, color := q[0], q[1]

		if prev, ok := ballColor[ball]; ok {
			colorCount[prev]--
			if colorCount[prev] == 0 {
				delete(colorCount, prev)
			}
		}

		ballColor[ball] = color
		colorCount[color]++
		ans[i] = len(colorCount)
	}
	return ans
}

func main() {
	fmt.Println(queryResults(4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}})) // Expected: [1, 2, 2, 3]
	fmt.Println(queryResults(4, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}})) // Expected: [1, 2, 2, 3, 4]
}
