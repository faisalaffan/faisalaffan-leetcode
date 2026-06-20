# 2188 — Minimum Time To Finish The Race

## Deskripsi

**Soal:** [2188. Minimum Time To Finish The Race](https://leetcode.com/problems/minimum-time-to-finish-the-race/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #2188: Minimum Time to Finish the Race
// https://leetcode.com/problems/minimum-time-to-finish-the-race/
// Difficulty: Hard
//
// DP: best[l] = min time for l consecutive laps with ONE tire (no change).
// dp[k] = min over j of dp[k-j] + changeTime + best[j].

import "fmt"

func main() {
	fmt.Println(minimumFinishTime([][]int{{2, 3}, {3, 4}}, 5, 4))            // 21
	fmt.Println(minimumFinishTime([][]int{{1, 10}, {2, 2}, {3, 4}}, 2, 5))   // 13
	fmt.Println(minimumFinishTime([][]int{{3, 4}}, 2, 3))                    // 13
}

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	const INF = 1 << 60

  // Membuat slice untuk menyimpan hasil
	best := make([]int, numLaps+1)
  // Iterasi seluruh elemen
	for i := range best {
		best[i] = INF
	}

	for _, tire := range tires {
		f, r := tire[0], tire[1]
		total := 0
		lapTime := f
		for l := 1; l <= numLaps; l++ {
			if total+lapTime >= INF {
				break
			}
			total += lapTime
			if total < best[l] {
				best[l] = total
					}
			if int64(lapTime)*int64(r) >= INF {
				break
			}
			lapTime *= r
		}
	}

  // Membuat slice untuk menyimpan hasil
	dp := make([]int, numLaps+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = -changeTime

	for i := 1; i <= numLaps; i++ {
		for j := 1; j <= i; j++ {
			if best[j] >= INF {
				continue
			}
			cand := dp[i-j] + changeTime + best[j]
			if cand < dp[i] {
				dp[i] = cand
			}
		}
	}
	return dp[numLaps]
}

func MinimumTimeToFinishTheRace() any {
	return minimumFinishTime([][]int{{2, 3}, {3, 4}}, 5, 4)
}
```
