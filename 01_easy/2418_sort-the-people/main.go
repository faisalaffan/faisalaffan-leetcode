package main

// LeetCode #2418: Sort the People
// https://leetcode.com/problems/sort-the-people/
// Difficulty: Easy
// Time O(n log n) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortThePeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})) // ["Mary","Emma","John"]
	fmt.Println(SortThePeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))   // ["Bob","Alice","Bob"]
}

func SortThePeople(names []string, heights []int) []string {
	n := len(names)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return heights[idx[i]] > heights[idx[j]]
	})
	res := make([]string, n)
	for i, id := range idx {
		res[i] = names[id]
	}
	return res
}
