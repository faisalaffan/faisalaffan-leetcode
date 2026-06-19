package main

// LeetCode #3444: Minimum Increments for Target Multiples in an Array
// https://leetcode.com/problems/minimum-increments-for-target-multiples-in-an-array/
// Difficulty: Hard
//
// DP over bitmask of targets. For each array element, compute cost to make it
// divisible by each subset S of targets (cost = nearest multiple of lcm(S)).
// Then dp[mask] = min cost to cover mask using processed elements.

import "fmt"

const INF = 1 << 60

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func minIncrementsForTargetMultiples(nums []int, target []int) int {
	n := len(nums)
	m := len(target)
	M := 1 << m

	// precomputeLCM[mask] = lcm of targets in mask
	lcmMask := make([]int, M)
	lcmMask[0] = 1
	for mask := 1; mask < M; mask++ {
		// find lowest set bit
		lsb := mask & -mask
		bit := 0
		for lsb>>bit != 1 {
			bit++
		}
		prev := mask ^ lsb
		if prev == 0 {
			lcmMask[mask] = target[bit]
		} else {
			l := lcm(lcmMask[prev], target[bit])
			// Clamp to avoid overflow (targets ≤ 50, nums ≤ 1e9)
			if l > 1_000_000_000 {
				l = 1_000_000_001
			}
			lcmMask[mask] = l
		}
	}

	// For each element, compute cost to cover each mask
	// cost[i][mask] = min increments to make nums[i] divisible by lcmMask[mask]
	elemCost := make([][]int, n)
	for i, x := range nums {
		elemCost[i] = make([]int, M)
		elemCost[i][0] = 0
		for mask := 1; mask < M; mask++ {
			l := lcmMask[mask]
			if l > 1_000_000_000 {
				elemCost[i][mask] = INF
				continue
			}
			// Nearest multiple of l ≥ x
			rem := x % l
			if rem == 0 {
				elemCost[i][mask] = 0
			} else {
				elemCost[i][mask] = l - rem
			}
		}
	}

	dp := make([]int, M)
	for i := 1; i < M; i++ {
		dp[i] = INF
	}

	for _, cost := range elemCost {
		ndp := make([]int, M)
		copy(ndp, dp)
		for oldMask := 0; oldMask < M; oldMask++ {
			if dp[oldMask] == INF {
				continue
			}
			for s := 1; s < M; s++ {
				if cost[s] == INF {
					continue
				}
				newMask := oldMask | s
				candidate := dp[oldMask] + cost[s]
				if candidate < ndp[newMask] {
					ndp[newMask] = candidate
				}
			}
		}
		dp = ndp
	}

	return dp[M-1]
}

func main() {
	// Test: nums=[1,2,3], target=[4,2,6] -> expected 4
	fmt.Printf("[1,2,3] target=[4,2,6] -> %d (expected 4)\n",
		minIncrementsForTargetMultiples([]int{1, 2, 3}, []int{4, 2, 6}))

	// Test: nums=[2,3,5], target=[3,5] -> expected ?
	// target 3: 2→3 cost 1; target 5: already 5 cost 0 => total 1
	fmt.Printf("[2,3,5] target=[3,5] -> %d\n",
		minIncrementsForTargetMultiples([]int{2, 3, 5}, []int{3, 5}))

	// Test: nums=[1], target=[2] -> 1→2 cost 1
	fmt.Printf("[1] target=[2] -> %d (expected 1)\n",
		minIncrementsForTargetMultiples([]int{1}, []int{2}))

	// Test: nums=[4,8,12], target=[3] -> need multiple of 3: 4→6 cost 2, 8→9 cost 1, 12→12 cost 0. Min 0.
	fmt.Printf("[4,8,12] target=[3] -> %d (expected 0)\n",
		minIncrementsForTargetMultiples([]int{4, 8, 12}, []int{3}))
}
