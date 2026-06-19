package main

// LeetCode #3740: Minimum Distance Between Three Equal Elements I
// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-i/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1, 2, 1, 1, 3}))
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1, 1, 2, 3, 2, 1, 2}))
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1}))
}

// Time: O(n)
// Space: O(n)
func MinimumDistanceBetweenThreeEqualElementsI(nums []int) int {
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	ans := math.MaxInt
	for _, indices := range pos {
		if len(indices) < 3 {
			continue
		}
		for i := 0; i <= len(indices)-3; i++ {
			dist := 2 * (indices[i+2] - indices[i])
			if dist < ans {
				ans = dist
			}
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
