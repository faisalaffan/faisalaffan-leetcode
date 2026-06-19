package main

// LeetCode #3068: Find the Maximum Sum of Node Values
// https://leetcode.com/problems/find-the-maximum-sum-of-node-values/
// Difficulty: Hard
//
// Approach: Sort by gain
// For each node, we can optionally XOR its value with k (gain[i] = (nums[i]^k) - nums[i]).
// Each operation XORs two connected nodes, so the number of XORed nodes must be even.
// Sort gains descending; pair them up; add pair to total if pair sum > 0.
// Tree structure is irrelevant because any even-sized subset is achievable
// via path-based XOR operations.

import (
	"fmt"
	"sort"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	total := int64(0)
	gains := make([]int, len(nums))
	for i, v := range nums {
		total += int64(v)
		gains[i] = (v ^ k) - v
	}
	sort.Slice(gains, func(i, j int) bool {
		return gains[i] > gains[j]
	})
	for i := 0; i+1 < len(gains); i += 2 {
		pairSum := gains[i] + gains[i+1]
		if pairSum > 0 {
			total += int64(pairSum)
		}
	}
	return total
}

func main() {
	// Example: nums=[1,2,1], k=3, edges=[[0,1],[0,2]] -> 6
	fmt.Println(maximumValueSum([]int{1, 2, 1}, 3, [][]int{{0, 1}, {0, 2}}))
	// Example: nums=[2,3], k=7, edges=[[0,1]] -> 9
	fmt.Println(maximumValueSum([]int{2, 3}, 7, [][]int{{0, 1}}))
	// Example: nums=[7,8,9], k=1, edges=[[0,1],[1,2]] -> 24
	fmt.Println(maximumValueSum([]int{7, 8, 9}, 1, [][]int{{0, 1}, {1, 2}}))
}
