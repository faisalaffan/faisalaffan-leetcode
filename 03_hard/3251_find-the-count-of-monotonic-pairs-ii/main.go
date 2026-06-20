package main

// LeetCode #3251: Find the Count of Monotonic Pairs II
// https://leetcode.com/problems/find-the-count-of-monotonic-pairs-ii/
// Difficulty: Hard
//
// Same problem as 3250 but with larger constraints: n ≤ 2000, nums[i] ≤ 1000.
// DP with prefix sums optimized to O(n * max(nums)).
//
// Given nums, count pairs (arr1, arr2) where:
//   - arr1 non-decreasing
//   - arr2 non-increasing
//   - arr1[i] + arr2[i] = nums[i]
//
// Time: O(n * maxVal), Space: O(maxVal)

import "fmt"

func main() {
	// Example 1: [2,3,2] => 4
	fmt.Println(countOfPairsII([]int{2, 3, 2}))
	// Example 2: [5,5,5,5] => 126
	fmt.Println(countOfPairsII([]int{5, 5, 5, 5}))
	// Example 3: [1,2,3,4] => 5
	fmt.Println(countOfPairsII([]int{1, 2, 3, 4}))
	// Example 4: single element
	fmt.Println(countOfPairsII([]int{10}))
	// Example 5: [0,0,0] => 1
	fmt.Println(countOfPairsII([]int{0, 0, 0}))
}

const MODII = 1_000_000_007

func countOfPairsII(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0] + 1
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	// dp[j] = ways for current position where arr1[i] = j
	dp := make([]int64, maxVal+1)
	for j := 0; j <= nums[0]; j++ {
		dp[j] = 1
	}

	for i := 1; i < n; i++ {
		a, b := nums[i-1], nums[i]

		// prefix sums
		prefix := make([]int64, maxVal+2)
		for j := 0; j <= maxVal; j++ {
			prefix[j+1] = (prefix[j] + dp[j]) % MODII
		}

		ndp := make([]int64, maxVal+1)

		// For each possible arr1[i] = j (0 <= j <= b):
		// We need arr1[i-1] = k where:
		//   0 <= k <= a (arr2[i-1] >= 0)
		//   k <= j (non-decreasing arr1)
		//   k <= j + a - b (non-increasing arr2: a-k >= b-j)
		for j := 0; j <= b; j++ {
			maxK := j
			if j+a-b < maxK {
				maxK = j + a - b
			}
			if maxK > a {
				maxK = a
			}
			if maxK >= 0 {
				ndp[j] = prefix[maxK+1]
			}
		}

		dp = ndp
	}

	var ans int64
	for j := 0; j <= nums[n-1]; j++ {
		ans = (ans + dp[j]) % MODII
	}
	return int(ans)
}
