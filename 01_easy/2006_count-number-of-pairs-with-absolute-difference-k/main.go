package main

// LeetCode #2006: Count Number of Pairs With Absolute Difference K
// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 2, 2, 1}, 1))   // 4
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 3}, 3))          // 0
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{3, 2, 1, 5, 4}, 2)) // 3
}

// Time: O(n), Space: O(n)
func CountNumberOfPairsWithAbsoluteDifferenceK(nums []int, k int) int {
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v-k] + freq[v+k]
		freq[v]++
	}
	return count
}
