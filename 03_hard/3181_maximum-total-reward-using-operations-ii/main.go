package main

// LeetCode #3181: Maximum Total Reward Using Operations II
// https://leetcode.com/problems/maximum-total-reward-using-operations-ii/
// Difficulty: Hard
//
// You have an array rewardValues. In each operation you can pick an un-picked
// item with reward value r, but only if your current total x satisfies x < r.
// After picking, your total becomes x + r.
//
// Return the maximum possible total reward.
//
// Approach: sort unique rewards, then DP with a bitset (big.Int). For each
// reward r, we only extend totals that are < r. Use a big integer bitset where
// bit i is set iff total i is reachable.

import (
	"fmt"
	"math/big"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	// Remove duplicates and sort.
	sort.Ints(rewardValues)
	uniq := []int{rewardValues[0]}
	for i := 1; i < len(rewardValues); i++ {
		if rewardValues[i] != rewardValues[i-1] {
			uniq = append(uniq, rewardValues[i])
		}
	}

	bitset := new(big.Int)
	bitset.SetBit(bitset, 0, 1) // total 0 is always reachable
	maxReward := uniq[len(uniq)-1]

	for _, r := range uniq {
		// mask = bitset & ((1 << r) - 1)   -> keep only totals < r
		mask := new(big.Int)
		limit := new(big.Int).Lsh(big.NewInt(1), uint(r))
		limit.Sub(limit, big.NewInt(1))
		mask.And(bitset, limit)

		// bitset |= mask << r
		shifted := new(big.Int).Lsh(mask, uint(r))
		bitset.Or(bitset, shifted)
	}

	// Find the highest set bit.
	ans := 0
	// The maximum possible total is at most 2 * maxReward.
	for x := 2 * maxReward; x >= 0; x-- {
		if bitset.Bit(x) == 1 {
			ans = x
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(maxTotalReward([]int{1, 6, 4, 3, 2})) // expect: 12 (pick 3+4+5? no, 1+2+3+6=12)
	fmt.Println(maxTotalReward([]int{10, 15, 25}))    // expect: 50?
}
