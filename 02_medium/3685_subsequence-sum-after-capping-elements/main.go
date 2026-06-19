package main

// LeetCode #3685: Subsequence Sum After Capping Elements
// https://leetcode.com/problems/subsequence-sum-after-capping-elements/
// Difficulty: Medium
// Time: O(n*k + n^2) | Space: O(k + n)

import "fmt"

func subsequenceSumAfterCappingElements(nums []int, k int) []bool {
	n := len(nums)
	ans := make([]bool, n)

	dp := make([]bool, k+1)
	dp[0] = true

	maxVal := n
	cnt := make([]int, maxVal+1)
	for _, v := range nums {
		if v <= maxVal {
			cnt[v]++
		} else {
			cnt[maxVal]++
		}
	}

	cntGe := make([]int, maxVal+2)
	for x := maxVal; x >= 1; x-- {
		cntGe[x] = cntGe[x+1] + cnt[x]
	}

	for x := 1; x <= n; x++ {
		ge := cntGe[x]
		ok := false
		maxM := ge
		if maxM > k/x {
			maxM = k / x
		}
		for m := 0; m <= maxM; m++ {
			if dp[k-m*x] {
				ok = true
				break
			}
		}
		ans[x-1] = ok

		c := cnt[x]
		if c == 0 {
			continue
		}
		p := 1
		for c > 0 {
			take := p
			if c < take {
				take = c
			}
			w := take * x
			for s := k; s >= w; s-- {
				if !dp[s] && dp[s-w] {
					dp[s] = true
				}
			}
			c -= take
			p <<= 1
		}
	}

	return ans
}

func main() {
	fmt.Println(subsequenceSumAfterCappingElements([]int{1, 2, 3}, 3))
	fmt.Println(subsequenceSumAfterCappingElements([]int{2, 4, 6}, 6))
	fmt.Println(subsequenceSumAfterCappingElements([]int{1, 1, 1}, 2))
}
