package main

// LeetCode #1502: Can Make Arithmetic Progression From Sequence
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
// Difficulty: Easy
//
// LeetCode submission: func canMakeArithmeticProgression(arr []int) bool

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{3, 5, 1})) // true
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{1, 2, 4})) // false
}

// Time: O(n log n), Space: O(1)
func CanMakeArithmeticProgressionFromSequence(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
