package main

// LeetCode #3289: The Two Sneaky Numbers of Digitville
// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{0, 1, 1, 0}))
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{0, 3, 2, 1, 3, 2}))
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}

// TheTwoSneakyNumbersOfDigitville returns the two numbers that appear twice in the array.
// Time: O(n). Space: O(n).
func TheTwoSneakyNumbersOfDigitville(nums []int) []int {
	seen := make(map[int]int)
	result := []int{}
	for _, num := range nums {
		seen[num]++
		if seen[num] == 2 {
			result = append(result, num)
		}
	}
	return result
}
