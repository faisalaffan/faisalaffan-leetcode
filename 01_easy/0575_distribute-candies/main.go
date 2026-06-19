package main

// LeetCode #575: Distribute Candies
// https://leetcode.com/problems/distribute-candies/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func DistributeCandies(candyType []int) int {
	types := make(map[int]bool)
	for _, c := range candyType {
		types[c] = true
	}
	maxAllowed := len(candyType) / 2
	if len(types) < maxAllowed {
		return len(types)
	}
	return maxAllowed
}

func main() {
	fmt.Println(DistributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(DistributeCandies([]int{1, 1, 2, 3}))
	fmt.Println(DistributeCandies([]int{6, 6, 6, 6}))
}
