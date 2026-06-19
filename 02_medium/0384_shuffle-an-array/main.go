package main

// LeetCode #384: Shuffle an Array
// https://leetcode.com/problems/shuffle-an-array/
// Difficulty: Medium
// Time: O(n) per shuffle | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	orig := make([]int, len(nums))
	copy(orig, nums)
	return Solution{original: orig}
}

func (s *Solution) Reset() []int {
	result := make([]int, len(s.original))
	copy(result, s.original)
	return result
}

func (s *Solution) Shuffle() []int {
	result := make([]int, len(s.original))
	copy(result, s.original)
	// Fisher-Yates
	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func main() {
	sol := Constructor([]int{1, 2, 3})
	fmt.Println("Reset:", sol.Reset())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Reset:", sol.Reset())
}
