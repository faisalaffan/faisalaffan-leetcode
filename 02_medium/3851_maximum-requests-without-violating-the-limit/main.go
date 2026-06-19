package main

// LeetCode #3851: Maximum Requests Without Violating the Limit
// https://leetcode.com/problems/maximum-requests-without-violating-the-limit/
// Difficulty: Medium [Paid]
// Time: O(N log N) | Space: O(N)
// Approach: For each user, group requests by time. Use sliding window to
// find max requests that can be kept per user without exceeding k in any
// window of size `window`.

import (
	"fmt"
	"sort"
)

func MaximumRequestsWithoutViolatingTheLimit(requests [][]int, k int, window int) int {
	// Group requests by user
	userRequests := make(map[int][]int)
	for _, r := range requests {
		user, time := r[0], r[1]
		userRequests[user] = append(userRequests[user], time)
	}

	total := 0

	for _, times := range userRequests {
		sort.Ints(times)
		// Use DP to find max requests we can keep
		// For each request, we can either keep it or drop it
		n := len(times)
		dp := make([]int, n+1)
		for i := 0; i < n; i++ {
			// Drop this request
			dp[i+1] = max(dp[i+1], dp[i])
			// Keep this request, find how many we can keep that end before times[i]-window
			// We want the first j where times[j] > times[i] - window - 1
			j := sort.Search(n, func(x int) bool { return times[x] > times[i]-window-1 })
			count := i - j + 1
			if count <= k {
				dp[i+1] = max(dp[i+1], dp[j]+count)
			}
		}
		total += dp[n]
	}

	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	req1 := [][]int{{1, 1}, {2, 1}, {1, 7}, {2, 8}}
	fmt.Println(MaximumRequestsWithoutViolatingTheLimit(req1, 1, 4)) // Expected: 4

	// Example 2
	req2 := [][]int{{1, 2}, {1, 5}, {1, 2}, {1, 6}}
	fmt.Println(MaximumRequestsWithoutViolatingTheLimit(req2, 2, 5)) // Expected: 2

	// Example 3
	req3 := [][]int{{1, 1}, {2, 5}, {1, 2}, {3, 9}}
	fmt.Println(MaximumRequestsWithoutViolatingTheLimit(req3, 1, 1)) // Expected: 3
}
