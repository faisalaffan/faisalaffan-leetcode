# 1434 — Number Of Ways To Wear Different Hats To Each Other

## Deskripsi

**Soal:** [1434. Number Of Ways To Wear Different Hats To Each Other](https://leetcode.com/problems/number-of-ways-to-wear-different-hats-to-each-other/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func numberWays(hats [][]int) int`

## Solusi Go

```go
package main

// LeetCode #1434: Number of Ways to Wear Different Hats to Each Other
// https://leetcode.com/problems/number-of-ways-to-wear-different-hats-to-each-other/
// Difficulty: Hard

import "fmt"

const mod1434 = 1_000_000_007

func numberWays(hats [][]int) int {
	n := len(hats)
	// Map each hat (1..40) to people who like it
  // Membuat slice 2D untuk DP/tabel
	hatToPeople := make([][]int, 41)
	for person, list := range hats {
		for _, hat := range list {
			hatToPeople[hat] = append(hatToPeople[hat], person)
		}
	}

	totalMasks := 1 << n
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, totalMasks)
	dp[0] = 1

	for hat := 1; hat <= 40; hat++ {
		if len(hatToPeople[hat]) == 0 {
			continue
		}
		// Iterate masks in reverse to avoid reusing the same hat
		for mask := totalMasks - 1; mask >= 0; mask-- {
			for _, person := range hatToPeople[hat] {
				if mask&(1<<person) != 0 {
					continue
				}
				nextMask := mask | (1 << person)
				dp[nextMask] = (dp[nextMask] + dp[mask]) % mod1434
			}
		}
	}
	return dp[totalMasks-1]
}

func main() {
	// Example: hats = [[3,4],[4,5],[5]] -> 1
	fmt.Println(numberWays([][]int{{3, 4}, {4, 5}, {5}}))
}
```
