package main

// LeetCode #3041: Maximize Consecutive Elements in an Array After Modification
// https://leetcode.com/problems/maximize-consecutive-elements-in-an-array-after-modification/
// Difficulty: Hard
//
// Approach: DP with hash map
// Sort nums first. For each element x, we can either keep it (x) or increment by 1 (x+1).
// f[v] = longest consecutive sequence ending at value v using processed elements.
// Process each element once; each element can extend f[x-1] (keeping x) or f[x] (as x+1).

import (
	"fmt"
	"sort"
)

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)
	f := make(map[int]int)
	ans := 0
	for _, x := range nums {
		old := f[x]
		f[x] = max(f[x], f[x-1]+1)
		f[x+1] = max(f[x+1], old+1)
	}
	for _, v := range f {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println("Example 1:", maxSelectedElements([]int{2, 1, 5, 1, 1}))
	// Expected: 3

	// Example 2
	fmt.Println("Example 2:", maxSelectedElements([]int{1, 4, 7, 10}))
	// Expected: 1

	// Single element
	fmt.Println("Single:", maxSelectedElements([]int{5}))
	// Expected: 1

	// Two elements same value
	fmt.Println("Two same:", maxSelectedElements([]int{3, 3}))
	// Expected: 2 (chain: 3, 4)

	// Large gap
	fmt.Println("Large gap:", maxSelectedElements([]int{1, 100, 200}))
	// Expected: 1

	// Descending input
	fmt.Println("Descending:", maxSelectedElements([]int{3, 2, 1}))
	// Expected: 3 (chain: 1, 2, 3)

	// All same value
	fmt.Println("All same:", maxSelectedElements([]int{5, 5, 5, 5}))
	// Expected: 2 (chain: 5, 6)

	// Negative numbers
	fmt.Println("Negative:", maxSelectedElements([]int{-1, 0, 1}))
	// Expected: 3 (chain: -1, 0, 1)

	// Duplicates forming chain
	fmt.Println("Duplicates:", maxSelectedElements([]int{1, 1, 2, 2, 3, 3}))
	// Expected: 5 (chain: 1, 2, 3, 4, can we get 5? Let's see)
	// Possibilities: 1→1, 1→2, 2→3, 2→4, 3→5, 3→6 → chain 1,2,3,4,5 length 5
}
