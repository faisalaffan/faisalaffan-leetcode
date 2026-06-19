package main

// LeetCode #2834: Find the Minimum Possible Sum of a Beautiful Array
// https://leetcode.com/problems/find-the-minimum-possible-sum-of-a-beautiful-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int {
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[target-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(2, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(3, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(5, 5))
}
