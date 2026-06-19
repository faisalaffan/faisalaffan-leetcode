package main

// LeetCode #3143: Maximum Points Inside the Square
// https://leetcode.com/problems/maximum-points-inside-the-square/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxPointsInsideSquare(points [][]int, s string) int {
	type item struct {
		dist int
		tag  byte
	}
	n := len(points)
	arr := make([]item, n)
	for i, p := range points {
		d := max(abs(p[0]), abs(p[1]))
		arr[i] = item{d, s[i]}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].dist < arr[j].dist
	})

	seen := make(map[byte]bool)
	ans := 0
	i := 0
	for i < n {
		j := i
		for j < n && arr[j].dist == arr[i].dist {
			if seen[arr[j].tag] {
				return ans
			}
			j++
		}
		for k := i; k < j; k++ {
			seen[arr[k].tag] = true
		}
		ans = len(seen)
		i = j
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(maxPointsInsideSquare([][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, "abdca"))
}
