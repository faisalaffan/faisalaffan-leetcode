package main

// LeetCode #3745: Maximize Expression of Three Elements
// https://leetcode.com/problems/maximize-expression-of-three-elements/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MaximizeExpressionOfThreeElements([]int{1, 4, 2, 5}))
	fmt.Println(MaximizeExpressionOfThreeElements([]int{-2, 0, 5, -2, 4}))
}

// Time: O(n)
// Space: O(1)
func MaximizeExpressionOfThreeElements(nums []int) int {
	max1, max2 := math.MinInt32, math.MinInt32
	min1 := math.MaxInt32

	for _, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
		if num < min1 {
			min1 = num
		}
	}

	return max1 + max2 - min1
}
