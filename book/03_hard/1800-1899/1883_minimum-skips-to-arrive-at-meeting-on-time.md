# 1883 — Minimum Skips To Arrive At Meeting On Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minSkipsToArriveAtMeetingOnTime(dist []int, speed int, hoursBefore int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1883: Minimum Skips to Arrive at Meeting On Time
// https://leetcode.com/problems/minimum-skips-to-arrive-at-meeting-on-time/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func minSkipsToArriveAtMeetingOnTime(dist []int, speed int, hoursBefore int) int {
	n := len(dist)
	// dp[i][j] = minimum time (in terms of distance units multiplied to avoid float)
	// to reach dist[i] with j skips. We use the fact that skipping means we avoid
	// the "round up" (arrive at a rest stop).
	// Time is represented as integer scaled by speed, but we need to handle rounding.
	//
	// Let's use the approach: dp[i][j] = minimum "delay" (extra time beyond raw travel)
	// represented as a value such that actual_time = (total_dist + dp[i][j]) / speed
	// No, let's use a cleaner approach.
	//
	// When we don't skip: time = ceil(time_before + dist[i] / speed)
	// When we skip: time = time_before + dist[i] / speed (no rounding)
	// where time_before is the time entering rest stop i.
	//
	// Use dp[i][j] = the total time (in units of 1/speed, i.e., dist units) after
	// reaching dist[i] with j skips, rounding up at each step except skipped ones.
	// But we need to track rounding status precisely.
	//
	// Better approach: dp[i][j] = the minimum "actual arrival time in hours" (as a rational)
	// after i-th road with j skips. Represent as ceil of a fraction.
	//
	// Let's use the technique from the editorial: use dp[i][j] = minimum time to complete
	// first i roads with j skips, but store time as an integer representing the "rounded"
	// value. The trick: when we don't skip, we round up. When we skip, we don't.
	// The check at the end: if dp[n][j] <= hoursBefore * speed, it's possible with j skips.
	// Because dp[n][j] stores total distance without the "rounding factor" of the last step.
	//
	// Actually the cleanest way: dp[i][j] = minimum number of "extra minutes" or such.
	// Let me just follow the well-known solution pattern:
	//
	// Use dp[i][j] = minimum time to cross first i roads with j skips.
	// Time is measured in units of speed (1 unit = 1/speed hours).
	// When crossing road i (0-indexed), we travel dist[i] units.
	// Without skip: new_time = ((dp[i-1][j] + dist[i-1] + speed - 1) / speed) * speed ...
	// No, this is getting messy. Let me use the simplest known working approach.
	//
	// We store dp[j] after processing each road, where dp[j] = total time in "distance units",
	// and we round up at each step unless we skip. At the end, check if dp[j] <= hoursBefore * speed.
	//
	// dp[j] = floor(total_time * speed) where total_time is accumulated time so far
	// with j skips. When we don't skip: dp[j] = ((dp[j] + dist[i] + speed - 1) / speed) * speed
	// When we skip: dp[j] = dp[j] + dist[i]
	// Wait no. Let me think in terms of the actual time in hours.
	//
	// If we track the "time elapsed so far" as a value where we keep the integer part
	// (hours passed) and the remainder (distance into the current hour).
	// Actually, let me just use the standard double/float with epsilon approach
	// but do it with integer math.
	//
	// The key insight from the editorial:
	// Let dp[i][j] = minimum time to finish first i roads with j skips, but we DON'T
	// apply the rounding up on the last segment.
	// dp[i][j] = min(
	//   skip on i: dp[i-1][j-1] + dist[i-1],                                 // if j > 0
	//   no skip on i: ceil(dp[i-1][j] + dist[i-1])                             // round up
	// )
	// where ceil(x) means round up to next multiple of 1 (in integer terms).
	// But how to represent "ceil" cleanly?
	//
	// Represent time in "fractional hours * speed" = distance units.
	// Let dp[j] = accumulated time in distance units after processing some roads.
	// Without rounding, total is just sum of dist.
	// With rounding, at each road the total gets rounded up to the next whole hour.
	//
	// Round up to next hour: ceil_to_hour(t) = ((t + speed - 1) / speed) * speed
	// where t is in distance units.
	//
	// So for road with distance d:
	// No skip: new_dp[j] = ((dp[j] + d + speed - 1) / speed) * speed
	// Skip:    new_dp[j] = dp[j-1] + d  (or old dp[j-1] + d)
	//
	// At the end, we have total_time in distance units. We need dp[n][j] <= hoursBefore * speed.
	// But we didn't round up on the last segment (the destination isn't a "rest stop").
	// Ah wait, we shouldn't round up on the final arrival. Let me reconsider.
	//
	// The roads end with the destination. We round up at each rest stop (which is between roads).
	// There are n-1 rest stops (after roads 0..n-2). Road n-1 leads directly to the destination.
	// So we round up for roads 0 to n-2 (arriving at rest stops), but NOT for road n-1.
	//
	// This means the last step is always "skip" automatically.
	// dp[n-1][j] represents the time to reach road n-1's start (which is a rest stop).
	// Then total_time = dp[n-1][j] + dist[n-1] (no rounding).
	// We need total <= hoursBefore * speed.
	//
	// For i from 0 to n-2:
	//   new_dp[j] = min(
	//     j > 0 ? old_dp[j-1] + dist[i] : INF,
	//     ((old_dp[j] + dist[i] + speed - 1) / speed) * speed
	//   )
	// Then final check: dp[j] + dist[n-1] <= hoursBefore * speed.

	// Handle: if sum of dist > hoursBefore * speed, impossible even with all skips
	totalDist := 0
	for _, d := range dist {
		totalDist += d
	}
	if totalDist > hoursBefore*speed {
		return -1
	}

	INF := math.MaxInt32
  // Alokasi slice
	dp := make([]int, n)
  // Range loop
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	// Process roads 0 to n-2 (arriving at rest stops)
	for i := 0; i < n-1; i++ {
		d := dist[i]
		// Process skips from high to low so we don't reuse same row
  // Alokasi slice
		newDp := make([]int, n)
		for j := range newDp {
			newDp[j] = INF
		}
		for j := 0; j <= i+1; j++ {
			// Option 1: no skip (round up)
			if dp[j] < INF {
				rounded := ((dp[j] + d + speed - 1) / speed) * speed
				if rounded < newDp[j] {
					newDp[j] = rounded
				}
			}
			// Option 2: skip (no round), can only do if j > 0
			if j > 0 && dp[j-1] < INF {
				noRound := dp[j-1] + d
				if noRound < newDp[j] {
					newDp[j] = noRound
				}
			}
		}
		dp = newDp
	}

	// Final road
	last := dist[n-1]
	for j := 0; j < n; j++ {
		if dp[j] < INF && dp[j]+last <= hoursBefore*speed {
			return j
		}
	}
	return -1
}

func main() {
	// Example: dist=[1,3,2], speed=4, hoursBefore=2 -> 1
	fmt.Println(minSkipsToArriveAtMeetingOnTime([]int{1, 3, 2}, 4, 2))

	// Additional test
	fmt.Println(minSkipsToArriveAtMeetingOnTime([]int{7, 3, 5, 5}, 2, 10))
}
```
