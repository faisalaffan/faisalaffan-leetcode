package main

// LeetCode #1550: Three Consecutive Odds
// https://leetcode.com/problems/three-consecutive-odds/
// Difficulty: Easy
//
// LeetCode submission: func threeConsecutiveOdds(arr []int) bool

import "fmt"

func main() {
	fmt.Println(ThreeConsecutiveOdds([]int{2, 6, 4, 1}))  // false
	fmt.Println(ThreeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12})) // true
}

// Time: O(n), Space: O(1)
func ThreeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, v := range arr {
		if v%2 == 1 {
			count++
			if count == 3 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}
