package main

// LeetCode #2891: Method Chaining
// https://leetcode.com/problems/method-chaining/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we filter animals with weight > 100, sort by weight,
// and rename the weight column.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: methodChaining
	// Input: [name, species, age, weight]
	animals := [][]int{{1, 1, 5, 50}, {2, 1, 3, 120}, {3, 2, 4, 150}, {4, 1, 2, 80}}
	fmt.Println(MethodChaining(animals))
	// [[2 1 3 120] [3 2 4 150]]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: methodChaining
func MethodChaining(animals [][]int) [][]int {
	// Filter: weight > 100 (column index 3)
	filtered := [][]int{}
	for _, a := range animals {
		if a[3] > 100 {
			filtered = append(filtered, a)
		}
	}
	// Sort by weight ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i][3] < filtered[j][3]
	})
	return filtered
}
