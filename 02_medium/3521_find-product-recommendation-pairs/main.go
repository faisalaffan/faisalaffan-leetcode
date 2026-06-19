package main

// LeetCode #3521: Find Product Recommendation Pairs
// https://leetcode.com/problems/find-product-recommendation-pairs/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindProductRecommendationPairs([]int{1, 2, 3, 4, 5}, 5))
	// Test case 2
	fmt.Println("Test 2:", FindProductRecommendationPairs([]int{1, 1, 1, 1}, 2))
	// Test case 3
	fmt.Println("Test 3:", FindProductRecommendationPairs([]int{1, 2, 3}, 7))
}

func FindProductRecommendationPairs(products []int, target int) [][]int {
	// Find pairs that sum to target
	var result [][]int
	seen := make(map[int]bool)
	for _, p := range products {
		complement := target - p
		if seen[complement] {
			result = append(result, []int{complement, p})
		}
		seen[p] = true
	}
	if result == nil {
		return [][]int{}
	}
	return result
}
