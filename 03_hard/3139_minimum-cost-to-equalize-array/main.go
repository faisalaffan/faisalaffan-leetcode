package main

// LeetCode #3139: Minimum Cost to Equalize Array
// https://leetcode.com/problems/minimum-cost-to-equalize-array/
// Difficulty: Hard
//
// You can increment one element by 1 (cost1) or two different elements by 1 (cost2).
// Find min cost to make all elements equal, modulo 1e9+7.

import (
	"fmt"
	"math"
)

const MOD = 1000000007

func main() {
	// Example: nums=[4,1], cost1=5, cost2=2 -> 15
	nums := []int{4, 1}
	cost1 := 5
	cost2 := 2
	fmt.Println(minCostToEqualizeArray(nums, cost1, cost2))
}

func minCostToEqualizeArray(nums []int, cost1, cost2 int) int {
	n := len(nums)
	minVal, maxVal := nums[0], nums[0]
	var sum int64 = 0
	for _, v := range nums {
		sum += int64(v)
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}

	if n == 1 {
		return 0
	}

	// If cost1*2 <= cost2, just use cost1 for all increments
	if cost1*2 <= cost2 {
		totalCost := int64(0)
		for _, v := range nums {
			totalCost += int64(maxVal-v) * int64(cost1)
		}
		return int(totalCost % MOD)
	}

	ans := int64(math.MaxInt64)
	// Check targets from maxVal upward. The optimal T is bounded.
	limit := maxVal + n*2 + 5

	for target := maxVal; target <= limit; target++ {
		totalIncs := int64(target)*int64(n) - sum
		maxDeficit := int64(target - minVal)

		// Max pairs = min(totalIncs/2, totalIncs - maxDeficit)
		pairs := totalIncs / 2
		if pairs > totalIncs-maxDeficit {
			pairs = totalIncs - maxDeficit
		}
		cost := pairs*int64(cost2) + (totalIncs-2*pairs)*int64(cost1)
		if cost < ans {
			ans = cost
		}
	}

	return int(ans % MOD)
}
