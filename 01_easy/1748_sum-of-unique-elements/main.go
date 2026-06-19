package main

// LeetCode #1748: Sum of Unique Elements
// https://leetcode.com/problems/sum-of-unique-elements/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func SumOfUnique(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sum := 0
	for num, count := range freq {
		if count == 1 {
			sum += num
		}
	}
	return sum
}

func main() {
	fmt.Println(SumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(SumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(SumOfUnique([]int{1, 2, 3, 4, 5}))
}
