package main

// LeetCode #3771: Total Score of Dungeon Runs
// https://leetcode.com/problems/total-score-of-dungeon-runs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func totalScoreOfDungeonRuns(hp int, damage []int, requirement []int) int64 {
	n := len(damage)
	// suffix cumulative damage
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + damage[i]
	}

	// For each starting index i, binary search how many rooms we can pass
	// Condition: hp - (damage[i] + ... + damage[j]) >= requirement[j]
	// => hp - requirement[j] >= suf[i] - suf[j+1]
	// => suf[j+1] >= suf[i] - (hp - requirement[j])
	// => if suf[i] - suf[j+1] <= hp - requirement[j]

	// Need to find j >= i where we have enough hp after damage to meet requirement
	// Equivalent: suf[i] - suf[j+1] <= hp - requirement[j]
	// => suf[j+1] >= suf[i] - (hp - requirement[j])

	// Binary search approach: for each i, find valid j range
	// Since suffix sums are decreasing as j increases, we can track valid range

	var ans int64
	for i := 0; i < n; i++ {
		// Binary search for furthest j we can reach
		lo, hi := i, n-1
		best := -1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			damageSum := suf[i] - suf[mid+1]
			if hp-damageSum >= requirement[mid] {
				best = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if best != -1 {
			// We can pass rooms i through best
			// This contributes to scores at start positions <= i
			// Specifically, for start position s where s <= i <= best:
			// score(s) includes this room
			// Number of start positions that include room i = i - 0 + 1 = i+1
			// But only if best >= i (always true here)

			// Actually simpler: for each start i, we pass (best - i + 1) rooms
			ans += int64(best - i + 1)
		}
	}
	return ans
}

func binarySearch(arr []int, target int) int {
	return sort.SearchInts(arr, target)
}

func main() {
	fmt.Println(totalScoreOfDungeonRuns(11, []int{3, 6, 7}, []int{4, 2, 5}))
	fmt.Println(totalScoreOfDungeonRuns(5, []int{1, 2, 3}, []int{1, 1, 1}))
	fmt.Println(totalScoreOfDungeonRuns(1, []int{5}, []int{10}))
}
