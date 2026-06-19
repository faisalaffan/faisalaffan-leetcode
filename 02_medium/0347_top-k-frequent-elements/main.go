package main

// LeetCode #347: Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func topKFrequent(nums []int, k int) []int {
	// Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Bucket sort by frequency (index = frequency)
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	// Expected: [1, 2]

	// Test case 2
	fmt.Println("Test 2:", topKFrequent([]int{1}, 1))
	// Expected: [1]

	// Test case 3
	fmt.Println("Test 3:", topKFrequent([]int{1, 2, 3, 1, 2, 1}, 2))
	// Expected: [1, 2]
}
