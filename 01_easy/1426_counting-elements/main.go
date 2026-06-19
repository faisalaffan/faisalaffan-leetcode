package main

// LeetCode #1426: Counting Elements
// https://leetcode.com/problems/counting-elements/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func countElements(arr []int) int

import "fmt"

func main() {
	fmt.Println(CountingElements([]int{1, 2, 3}))       // 2
	fmt.Println(CountingElements([]int{1, 1, 3, 3, 5, 5, 7, 7})) // 0
	fmt.Println(CountingElements([]int{1, 1, 2, 2}))    // 2
}

// Time: O(n), Space: O(n)
func CountingElements(arr []int) int {
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		seen[v] = true
	}
	count := 0
	for _, v := range arr {
		if seen[v+1] {
			count++
		}
	}
	return count
}
