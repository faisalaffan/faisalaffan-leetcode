package main

// LeetCode #3086: Minimum Moves to Pick K Ones
// https://leetcode.com/problems/minimum-moves-to-pick-k-ones/
// Difficulty: Hard
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumMoves(nums []int, k int) int64 {
	// Collect positions of all 1s
	pos := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			pos = append(pos, i)
		}
	}

	n := len(pos)
	if n < k {
		// Not enough ones — would need to convert zeros (not possible in this version)
		return -1
	}

	// Prefix sums for O(1) interval sum
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(pos[i])
	}

	var result int64 = math.MaxInt64

	// Sliding window of size k over positions
	for r := k; r <= n; r++ {
		l := r - k
		mid := l + k/2
		medianPos := int64(pos[mid])

		// Left side cost: sum(median - pos[i]) for i in [l, mid-1]
		leftCount := int64(mid - l)
		leftSum := medianPos*leftCount - (prefix[mid] - prefix[l])

		// Right side cost: sum(pos[i] - median) for i in [mid+1, r-1]
		rightCount := int64(r - mid - 1)
		rightSum := (prefix[r] - prefix[mid+1]) - medianPos*rightCount

		total := leftSum + rightSum
		if total < result {
			result = total
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumMoves([]int{1, 0, 0, 1, 1, 0, 1}, 3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minimumMoves([]int{1, 1, 0, 1}, 2))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minimumMoves([]int{1, 1, 1}, 2))
	// Expected: 1
}
