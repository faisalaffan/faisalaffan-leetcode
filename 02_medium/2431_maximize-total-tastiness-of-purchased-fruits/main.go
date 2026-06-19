package main

// LeetCode #2431: Maximize Total Tastiness of Purchased Fruits
// https://leetcode.com/problems/maximize-total-tastiness-of-purchased-fruits/
// Difficulty: Medium
// Time: O(n * budget * coupon) | Space: O(budget * coupon)
// Knapsack DP: dp[b][c] = max tastiness with b budget and c coupons remaining.

import "fmt"

type Fruit struct {
	price, tastiness int
}

func main() {
	fruits := []Fruit{{2, 3}, {3, 6}, {5, 10}}
	fmt.Println(maxTastiness(fruits, 10, 1)) // 16 (buy 3rd with coupon, 1st and 2nd normally)

	fruits2 := []Fruit{{1, 5}, {2, 3}, {3, 6}}
	fmt.Println(maxTastiness(fruits2, 5, 2)) // 14
}

func maxTastiness(fruits []Fruit, budget int, couponCount int) int {
	dp := make([][]int, budget+1)
	for b := range dp {
		dp[b] = make([]int, couponCount+1)
		for c := range dp[b] {
			dp[b][c] = -1
		}
	}
	dp[0][0] = 0
	ans := 0

	for _, f := range fruits {
		newDp := make([][]int, budget+1)
		for b := range newDp {
			newDp[b] = make([]int, couponCount+1)
			copy(newDp[b], dp[b])
		}

		for b := budget; b >= 0; b-- {
			for c := 0; c <= couponCount; c++ {
				if dp[b][c] < 0 {
					continue
				}
				// buy normally
				if b+f.price <= budget {
					if dp[b][c]+f.tastiness > newDp[b+f.price][c] {
						newDp[b+f.price][c] = dp[b][c] + f.tastiness
					}
				}
				// buy with coupon (half price)
				if c < couponCount {
					cp := f.price / 2
					if b+cp <= budget {
						if dp[b][c]+f.tastiness > newDp[b+cp][c+1] {
							newDp[b+cp][c+1] = dp[b][c] + f.tastiness
						}
					}
				}
			}
		}
		dp = newDp
	}

	for b := 0; b <= budget; b++ {
		for c := 0; c <= couponCount; c++ {
			if dp[b][c] > ans {
				ans = dp[b][c]
			}
		}
	}
	return ans
}
