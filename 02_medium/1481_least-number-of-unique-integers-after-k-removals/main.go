package main

// LeetCode #1481: Least Number of Unique Integers after K Removals
// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1)) // 1

	// Test case 2
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3)) // 2

	// Test case 3
	fmt.Println(findLeastNumOfUniqueInts([]int{1, 2, 3, 4, 5}, 5)) // 0
}

// Time: O(n log n) for sorting frequencies
// Space: O(n) for frequency map
func findLeastNumOfUniqueInts(arr []int, k int) int {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}

	sort.Ints(counts)

	remaining := k
	uniqueCount := len(counts)

	for _, c := range counts {
		if remaining >= c {
			remaining -= c
			uniqueCount--
		} else {
			break
		}
	}

	return uniqueCount
}
