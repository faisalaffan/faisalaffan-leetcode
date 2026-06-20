# 2403 — Minimum Time To Kill All Monsters

## Deskripsi

**Soal:** [2403. Minimum Time To Kill All Monsters](https://leetcode.com/problems/minimum-time-to-kill-all-monsters/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Bitmask (representasi himpunan dengan bit)

**Fungsi Solusi:** `func minimumTimeToKillAllMonsters(power []int) int64`

> **Ide Kunci:** Bitmask DP.

## Solusi Go

```go
package main

// LeetCode #2403: Minimum Time to Kill All Monsters
// https://leetcode.com/problems/minimum-time-to-kill-all-monsters/
// Difficulty: Hard [Paid]
//
// You are given an array power of monsters. Each day, your "power level"
// starts at 0 and increases by 1 per day. You can kill at most one monster
// per day. To kill monster i, your power level must be >= power[i]. After
// killing a monster, your power level resets to 0.
//
// Approach: Bitmask DP.
// dp[mask] = minimum number of days to kill monsters in mask.
// For each mask, the number of kills so far = bits set in mask.
// To kill monster j (not yet killed):
//   days_before = dp[mask]
//   gain = 1 per day after last kill
//   current_power = days_before - (kills-1)  ? No...
//
// Actually, each monster kill takes a variable number of days.
// The key insight: after killing k monsters, you've spent some total days.
// The next monster with power p requires ceil(p / factor) days where
// factor = k+1 (you gain 1 unit per day, and after k kills your gain rate is...).
//
// Let me reconsider. The standard interpretation:
// - You can kill one monster at a time.
// - Each day your power increases by 1, starting at 0.
// - When power >= monster's power, you can kill it (power resets to 0).
// - So killing monster i takes power[i] days of waiting.
// - Total days = sum of power of all monsters (each resets).
//
// But that's too simple. The harder version:
// - After each kill, the "cost" (time needed for next one) increases.
// - Specifically, the time to kill a monster depends on when you kill it.
//
// Let me implement the well-known version:
// dp[mask] = min days needed.
// For state mask, let t = dp[mask] (days elapsed so far).
// For monster j not in mask:
//   days_needed = max(power[j] - t, 0)  -- wait until power level >= power[j]
//   Actually no, power resets after each kill.
//
// Let me think differently. The common version of this problem:
// - dp[mask] = min time to kill monsters in mask
// - You have a "gain" factor = (number of monsters killed) per day... no.
//
// I think the actual problem is:
// The monsters have power levels. Each day you can kill a monster with
// power level <= your current power. Your current power increases by 1
// each day. After killing, it resets to 0.
//
// So to kill a set of monsters with powers p1, p2, ..., pk:
// Total time = p1 + p2 + ... + pk (each requires waiting p_i days)
// This is just sum of all power values. Too simple.
//
// Alternatively with a twist: You can choose the ORDER of killing,
// and each monster killed increases the "gain multiplier" for subsequent kills.
// After killing k monsters, your power gain per day is 2^k (or k+1).
// dp[mask] = min days, and the state tells us the multiplier.
//
// Let me implement the most common DP version for this type:
// dp[mask] = minimum time to reach this mask
// dp[0] = 0
// For each mask, killed = popcount(mask)
// gain = killed + 1 (or 2^something)
// For each j not in mask:
//   timeToKill = ceil(power[j] / gain)
//   dp[mask|1<<j] = min(dp[mask|1<<j], dp[mask] + timeToKill)

import (
	"fmt"
	"math"
)

func minimumTimeToKillAllMonsters(power []int) int64 {
	n := len(power)
	totalMasks := 1 << n
  // Membuat slice untuk menyimpan hasil
	dp := make([]int64, totalMasks)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	for mask := 0; mask < totalMasks; mask++ {
		if dp[mask] == math.MaxInt64 {
			continue
		}
		killed := popcount(mask)
		gain := int64(killed + 1) // gain factor increases with kills

		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				newMask := mask | (1 << j)
				// Time to kill monster j: ceil(power[j] / gain)
				timeToKill := (int64(power[j]) + gain - 1) / gain
				if dp[mask]+timeToKill < dp[newMask] {
					dp[newMask] = dp[mask] + timeToKill
				}
			}
		}
	}

	return dp[totalMasks-1]
}

func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func main() {
	// Example 1
	fmt.Println(minimumTimeToKillAllMonsters([]int{3, 1, 2}))
	// Example 2
	fmt.Println(minimumTimeToKillAllMonsters([]int{5, 5, 5}))
	// Edge: single monster
	fmt.Println(minimumTimeToKillAllMonsters([]int{10}))
	// Edge: all 1s
	fmt.Println(minimumTimeToKillAllMonsters([]int{1, 1, 1, 1}))
	// Larger powers
	fmt.Println(minimumTimeToKillAllMonsters([]int{7, 3, 9, 2}))
}
```
