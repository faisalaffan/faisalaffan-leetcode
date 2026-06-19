package main

// LeetCode #2552: Count Increasing Quadruplets
// https://leetcode.com/problems/count-increasing-quadruplets/
// Difficulty: Hard

import "fmt"

// countQuadruplets counts quadruplets (i,j,k,l) with i<j<k<l
// such that nums[i] < nums[k] < nums[j] < nums[l] (1324 pattern).
//
// For each pair (j,k) with j<k and nums[j] > nums[k]:
//   - Count i<j where nums[i] < nums[k] (from prefix)
//   - Count l>k where nums[l] > nums[j] (from suffix)
//   - Product of counts = quadruplets for this (j,k)
//
// Complexity: O(n^2) time, O(n^2) space
func countQuadruplets(nums []int) int64 {
	n := len(nums)

	// prefixLess[i][v] = count of elements before position i that are < v
	prefixLess := make([][]int, n+1)
	for i := range prefixLess {
		prefixLess[i] = make([]int, n+2)
	}
	for i := 0; i < n; i++ {
		copy(prefixLess[i+1], prefixLess[i])
		for v := nums[i] + 1; v <= n; v++ {
			prefixLess[i+1][v]++
		}
	}

	// suffixGreater[i][v] = count of elements after position i that are > v
	suffixGreater := make([][]int, n+1)
	for i := range suffixGreater {
		suffixGreater[i] = make([]int, n+2)
	}
	for i := n - 1; i >= 0; i-- {
		copy(suffixGreater[i], suffixGreater[i+1])
		for v := 1; v < nums[i]; v++ {
			suffixGreater[i][v]++
		}
	}

	var result int64
	for j := 1; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				left := prefixLess[j][nums[k]]
				right := suffixGreater[k][nums[j]]
				result += int64(left) * int64(right)
			}
		}
	}
	return result
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: [1,3,2,4,5] ->", countQuadruplets([]int{1, 3, 2, 4, 5})) // 2

	// Additional test cases
	fmt.Println("Test 2: [1,2,3,4] ->", countQuadruplets([]int{1, 2, 3, 4}))     // 0
	fmt.Println("Test 3: [5,4,3,2,1] ->", countQuadruplets([]int{5, 4, 3, 2, 1})) // 0
	fmt.Println("Test 4: [1,4,3,5,2] ->", countQuadruplets([]int{1, 4, 3, 5, 2}))
	fmt.Println("Test 5: [2,5,3,4,1] ->", countQuadruplets([]int{2, 5, 3, 4, 1}))
}
