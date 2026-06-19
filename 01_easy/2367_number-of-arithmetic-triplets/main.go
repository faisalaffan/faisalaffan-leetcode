package main

// LeetCode #2367: Number of Arithmetic Triplets
// https://leetcode.com/problems/number-of-arithmetic-triplets/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(NumberOfArithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3)) // 2
	fmt.Println(NumberOfArithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))  // 2
}

func NumberOfArithmeticTriplets(nums []int, diff int) int {
	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		seen[n] = true
	}
	count := 0
	for _, n := range nums {
		if seen[n+diff] && seen[n+2*diff] {
			count++
		}
	}
	return count
}
