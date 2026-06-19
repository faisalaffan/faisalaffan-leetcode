package main

// LeetCode #3224: Minimum Array Changes to Make Differences Equal
// https://leetcode.com/problems/minimum-array-changes-to-make-differences-equal/
// Difficulty: Medium
// Time: O(n + k) | Space: O(k)

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minChanges(nums []int, k int) int {
	n := len(nums)
	diff := make([]int, k+2)

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		curDiff := abs(a - b)

		// One change: can achieve any diff from 0 to max(a, b, k-a, k-b)
		maxReach := max(max(a, b), max(k-a, k-b))
		// One change can achieve any diff in [0, maxReach]
		diff[0]++
		if maxReach+1 <= k {
			diff[maxReach+1]--
		}

		// Zero changes: only curDiff
		diff[curDiff]--
		diff[curDiff+1]++
	}

	ans := n
	cur := 0
	for i := 0; i <= k; i++ {
		cur += diff[i]
		if cur < ans {
			ans = cur
		}
	}
	return ans
}

func main() {
	fmt.Println(minChanges([]int{1, 0, 1, 2, 4, 3}, 4)) // Expected: 2
	fmt.Println(minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6)) // Expected: ?
}
