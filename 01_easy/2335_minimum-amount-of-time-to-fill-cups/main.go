package main

// LeetCode #2335: Minimum Amount of Time to Fill Cups
// https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups/
// Difficulty: Easy
// Time O(1) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumAmountOfTimeToFillCups([]int{1, 4, 2})) // 4
	fmt.Println(MinimumAmountOfTimeToFillCups([]int{5, 4, 4})) // 5
}

func MinimumAmountOfTimeToFillCups(amount []int) int {
	sort.Ints(amount)
	a, b, c := amount[0], amount[1], amount[2]
	if a+b <= c {
		return c
	}
	return (a+b+c+1)/2
}
