package main

// LeetCode #3027: Find the Number of Ways to Place People II
// https://leetcode.com/problems/find-the-number-of-ways-to-place-people-ii/
// Difficulty: Hard

import (
	"fmt"
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	ans := 0
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	for i := 0; i < len(points)-1; i++ {
		xMax := math.MaxInt32
		yMin := math.MinInt32
		for j := i + 1; j < len(points); j++ {
			if points[j][0] > points[i][0]-1 && points[j][0] < xMax &&
				points[j][1] > yMin && points[j][1] < points[i][1]+1 {
				ans++
				xMax = points[j][0]
				yMin = points[j][1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfPairs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(numberOfPairs([][]int{{1, 2}, {2, 1}, {3, 1}}))
}
