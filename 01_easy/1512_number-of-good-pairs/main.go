package main

// LeetCode #1512: Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/
// Difficulty: Easy
//
// LeetCode submission: func numIdenticalPairs(nums []int) int

import "fmt"

func main() {
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3, 1, 1, 3})) // 4
	fmt.Println(NumberOfGoodPairs([]int{1, 1, 1, 1}))        // 6
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3}))           // 0
}

// Time: O(n), Space: O(n)
func NumberOfGoodPairs(nums []int) int {
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v]
		freq[v]++
	}
	return count
}
