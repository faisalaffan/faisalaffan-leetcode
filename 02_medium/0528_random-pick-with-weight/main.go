package main

// LeetCode #528: Random Pick with Weight
// https://leetcode.com/problems/random-pick-with-weight/
// Difficulty: Medium
// Time: O(n) for init, O(log n) per pick
// Space: O(n)

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	sol := Constructor([]int{1, 3})
	fmt.Println(sol.PickIndex())
	fmt.Println(sol.PickIndex())
	fmt.Println(sol.PickIndex())
}

type Solution struct {
	prefixSum []int
	totalSum  int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, len(w))
	sum := 0
	for i, weight := range w {
		sum += weight
		prefixSum[i] = sum
	}
	return Solution{prefixSum: prefixSum, totalSum: sum}
}

func (s *Solution) PickIndex() int {
	target := rand.Intn(s.totalSum) + 1
	return sort.SearchInts(s.prefixSum, target)
}
