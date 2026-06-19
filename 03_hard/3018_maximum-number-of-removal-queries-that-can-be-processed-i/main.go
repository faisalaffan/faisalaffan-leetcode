package main

// LeetCode #3018: Maximum Number of Removal Queries That Can Be Processed I
// https://leetcode.com/problems/maximum-number-of-removal-queries-that-can-be-processed-i/
// Difficulty: Hard [Paid]
//
// We have an array nums and a list of queries. We process queries in order.
// At each step, we can remove either the first or last element of nums.
// A query can be processed if the removed element >= query value.
// Find the maximum number of queries we can process.
//
// Approach: Interval DP
//   dp[l][r] = maximum number of queries processed when the remaining
//   subarray is nums[l..r] (inclusive). We process queries in order,
//   and at each step we try removing either nums[l-1] or nums[r+1].

import "fmt"

func maximumProcessableQueries(nums []int, queries []int) int {
	n := len(nums)
	m := len(queries)

	// dp[l][r] = max queries processed with remaining subarray nums[l..r]
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// Base case: single element
	for i := 0; i < n; i++ {
		if m > 0 && nums[i] >= queries[0] {
			dp[i][i] = 1
		} else {
			dp[i][i] = 0
		}
	}

	ans := 0
	for length := 1; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			val := 0

			// Try removing nums[l-1] (left of current subarray)
			if l > 0 {
				prev := dp[l-1][r]
				if prev > val {
					val = prev
				}
				if prev >= 0 && prev < m && nums[l-1] >= queries[prev] {
					if prev+1 > val {
						val = prev + 1
					}
				}
			}

			// Try removing nums[r+1] (right of current subarray)
			if r+1 < n {
				prev := dp[l][r+1]
				if prev > val {
					val = prev
				}
				if prev >= 0 && prev < m && nums[r+1] >= queries[prev] {
					if prev+1 > val {
						val = prev + 1
					}
				}
			}

			dp[l][r] = val
			if val > ans {
				ans = val
			}
		}
	}

	// Final check: try to process one more query if possible
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			val := dp[l][r]
			if val > ans {
				ans = val
			}
			if l > 0 && val < m && nums[l-1] >= queries[val] && val+1 > ans {
				ans = val + 1
			}
			if r+1 < n && val < m && nums[r+1] >= queries[val] && val+1 > ans {
				ans = val + 1
			}
		}
	}

	return ans
}

func main() {
	// Test 1
	fmt.Println("Test 1:", maximumProcessableQueries([]int{1, 2, 3, 4}, []int{1, 2}))

	// Test 2: Can't process any
	fmt.Println("Test 2:", maximumProcessableQueries([]int{1, 1, 1}, []int{5, 5}))

	// Test 3: All queries processable
	fmt.Println("Test 3:", maximumProcessableQueries([]int{10, 20, 30}, []int{1, 2, 3}))

	// Test 4: Single element, single query
	fmt.Println("Test 4:", maximumProcessableQueries([]int{5}, []int{3}))
	fmt.Println("Test 5:", maximumProcessableQueries([]int{5}, []int{6}))

	// Test 6: Multiple possibilities
	fmt.Println("Test 6:", maximumProcessableQueries([]int{5, 1, 3, 2, 4}, []int{2, 3, 1}))
}
