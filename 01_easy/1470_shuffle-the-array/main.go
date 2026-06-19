package main

// LeetCode #1470: Shuffle the Array
// https://leetcode.com/problems/shuffle-the-array/
// Difficulty: Easy
//
// LeetCode submission: func shuffle(nums []int, n int) []int

import "fmt"

func main() {
	fmt.Println(ShuffleTheArray([]int{2, 5, 1, 3, 4, 7}, 3)) // [2 3 5 4 1 7]
	fmt.Println(ShuffleTheArray([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4)) // [1 4 2 3 3 2 4 1]
}

// Time: O(n), Space: O(n)
func ShuffleTheArray(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}
