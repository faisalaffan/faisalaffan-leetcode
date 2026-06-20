package main

// LeetCode #3930: Power Update After K-th Largest Insertion II
// https://leetcode.com/problems/power-update-after-k-th-largest-insertion-ii/
// Difficulty: Hard [Paid]
//
// Given array nums and queries [index, val]. For each query, set
// nums[index] += val. After each update, compute the "power": the
// XOR of all elements that are multiples of a given integer p.
// Return power after each query.
//
// Approach: Maintain current total XOR of elements that are
// multiples of p. For each query, check if old value changes
// affect the XOR, then apply update and check new value.

import "fmt"

func main() {
	// Example 1
	fmt.Println(powerUpdate([]int{2, 4, 6, 8}, 2, [][]int{{1, 2}, {0, 2}}))
	// Example 2
	fmt.Println(powerUpdate([]int{1, 3, 5}, 3, [][]int{{0, 2}, {2, 1}}))
	// Edge: single element
	fmt.Println(powerUpdate([]int{6}, 2, [][]int{{0, 2}}))
}

func powerUpdate(nums []int, p int, queries [][]int) []int {
	n := len(nums)
	totalXor := 0
	included := make([]bool, n)
	for i, v := range nums {
		if v%p == 0 {
			totalXor ^= v
			included[i] = true
		}
	}

	ans := make([]int, len(queries))
	for idx, q := range queries {
		i, val := q[0], q[1]
		old := nums[i]
		nums[i] += val

		if included[i] {
			totalXor ^= old
		}
		if nums[i]%p == 0 {
			totalXor ^= nums[i]
			included[i] = true
		} else {
			included[i] = false
		}

		ans[idx] = totalXor
	}

	return ans
}
