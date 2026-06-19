package main

// LeetCode #2892: Minimizing Array After Replacing Pairs With Their Product
// https://leetcode.com/problems/minimizing-array-after-replacing-pairs-with-their-product/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func MinimizingArrayAfterReplacingPairsWithTheirProduct(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// Greedy: merge adjacent elements whose product is <= their sum
	// Actually, we merge if product > maxElement (since product will grow)
	// For positive integers > 1, product grows quickly
	// If any element is 1, it can always be merged (product = other * 1, sum = other + 1)

	result := make([]int, 0, n)
	for _, v := range nums {
		result = append(result, v)
		for len(result) >= 2 {
			last := len(result) - 1
			// Merge if product <= sum
			p := result[last] * result[last-1]
			s := result[last] + result[last-1]
			if p <= s {
				merged := result[last] * result[last-1]
				result = result[:last-1]
				result = append(result, merged)
			} else {
				break
			}
		}
	}

	return len(result)
}

func main() {
	fmt.Println(MinimizingArrayAfterReplacingPairsWithTheirProduct([]int{2, 3, 3}))
	fmt.Println(MinimizingArrayAfterReplacingPairsWithTheirProduct([]int{1, 2, 1, 2}))
}
