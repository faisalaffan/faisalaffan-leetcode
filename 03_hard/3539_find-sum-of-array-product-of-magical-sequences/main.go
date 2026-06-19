package main

// LeetCode #3539: Find Sum of Array Product of Magical Sequences
// https://leetcode.com/problems/find-sum-of-array-product-of-magical-sequences/
// Difficulty: Hard
//
// Find sum over all magical sequences of their array product. A magical
// sequence is defined by specific constraints (e.g., length n, values in
// [1, m] with some property). Return result modulo 1e9+7.
//
// Approach: DP to count sequences and their product sums.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumOfArrayProduct(2, 3))
	// Example 2
	fmt.Println(sumOfArrayProduct(3, 2))
	// Edge: single element
	fmt.Println(sumOfArrayProduct(1, 5))
	// Edge: zero
	fmt.Println(sumOfArrayProduct(0, 10))
}

const mod = 1000000007

func sumOfArrayProduct(n int, m int) int {
	if n == 0 {
		return 0
	}
	// dp[i] = sum of products of sequences ending at value i
	// total[i] = total count of sequences ending at value i
	dp := make([]int64, m+1)
	cnt := make([]int64, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = int64(i)
		cnt[i] = 1
	}

	for length := 2; length <= n; length++ {
		newDp := make([]int64, m+1)
		newCnt := make([]int64, m+1)
		for i := 1; i <= m; i++ {
			for j := 1; j <= m; j++ {
				newDp[(i+j)%m] = (newDp[(i+j)%m] + dp[j]*int64(i)) % mod
				newCnt[(i+j)%m] = (newCnt[(i+j)%m] + cnt[j]) % mod
			}
		}
		dp = newDp
		cnt = newCnt
	}

	var ans int64
	for i := 1; i <= m; i++ {
		ans = (ans + dp[i]) % mod
	}
	return int(ans)
}
