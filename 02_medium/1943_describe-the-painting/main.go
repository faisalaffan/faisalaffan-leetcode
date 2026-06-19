package main

// LeetCode #1943: Describe the Painting
// https://leetcode.com/problems/describe-the-painting/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {4, 7, 7}, {1, 7, 9}}))
	fmt.Println(SplitPainting([][]int{{1, 7, 9}, {6, 8, 15}, {8, 10, 7}}))
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {1, 4, 7}, {4, 7, 1}, {4, 7, 11}}))
}

// Time: O(n log n), Space: O(n)
func SplitPainting(segments [][]int) [][]int64 {
	diff := make(map[int]int64)
	endpoints := make(map[int]bool)

	for _, seg := range segments {
		start, end, color := seg[0], seg[1], seg[2]
		diff[start] += int64(color)
		diff[end] -= int64(color)
		endpoints[start] = true
		endpoints[end] = true
	}

	// Sort unique endpoints
	points := make([]int, 0, len(endpoints))
	for p := range endpoints {
		points = append(points, p)
	}
	sort.Ints(points)

	result := make([][]int64, 0)
	var sum int64 = 0
	for i := 0; i < len(points)-1; i++ {
		sum += diff[points[i]]
		if sum != 0 {
			result = append(result, []int64{int64(points[i]), int64(points[i+1]), sum})
		}
	}
	return result
}
