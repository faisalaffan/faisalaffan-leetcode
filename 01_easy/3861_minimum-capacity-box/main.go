package main

// LeetCode #3861: Minimum Capacity Box
// https://leetcode.com/problems/minimum-capacity-box/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumCapacityBox([]int{1, 5, 3, 7}, 3))
	fmt.Println(MinimumCapacityBox([]int{3, 5, 4, 3}, 2))
	fmt.Println(MinimumCapacityBox([]int{4}, 5))
}

// Time: O(n)
// Space: O(1)
func MinimumCapacityBox(capacity []int, itemSize int) int {
	ans := -1
	for i, c := range capacity {
		if c >= itemSize && (ans == -1 || c < capacity[ans]) {
			ans = i
		}
	}
	return ans
}
