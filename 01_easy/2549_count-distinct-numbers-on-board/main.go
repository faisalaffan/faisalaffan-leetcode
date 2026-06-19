package main

// LeetCode #2549: Count Distinct Numbers on Board
// https://leetcode.com/problems/count-distinct-numbers-on-board/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountDistinctNumbersOnBoard(5)) // 4
	fmt.Println(CountDistinctNumbersOnBoard(2)) // 1
	fmt.Println(CountDistinctNumbersOnBoard(1)) // 1
}

func CountDistinctNumbersOnBoard(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}
