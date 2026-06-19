package main

// LeetCode #1228: Missing Number In Arithmetic Progression
// https://leetcode.com/problems/missing-number-in-arithmetic-progression/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{5, 7, 11, 13}))  // 9
	fmt.Println(missingNumber([]int{15, 13, 12}))    // 14
}

// LeetCode submission: missingNumber
func missingNumber(arr []int) int {
	n := len(arr)
	diff := (arr[n-1] - arr[0]) / n
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] == arr[0]+mid*diff {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[0] + diff*lo
}
