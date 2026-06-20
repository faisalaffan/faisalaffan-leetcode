# 3385 — Minimum Time To Break Locks Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumTimeToBreakLocksIi(strength []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3385: Minimum Time to Break Locks II
// https://leetcode.com/problems/minimum-time-to-break-locks-ii/
// Difficulty: Hard [Paid]
//
// Bitmask DP. Each lock broken adds 1 to power multiplier.
// time = ceil(strength[i] / power).

import "fmt"

func main() {
	fmt.Println(MinimumTimeToBreakLocksIi([]int{3, 4, 1}, 2))
}

func MinimumTimeToBreakLocksIi(strength []int, k int) int {
	n := len(strength)
	m := 1 << n
  // Alokasi slice
	dp := make([]int, m)
	for mask := 1; mask < m; mask++ {
		dp[mask] = 1 << 60
	}

	dp[0] = 0
	for mask := 0; mask < m; mask++ {
		broken := 0
		for b := 0; b < n; b++ {
			if mask>>b&1 == 1 {
				broken++
			}
		}
		power := k + broken
		for b := 0; b < n; b++ {
			if mask>>b&1 == 1 {
				continue
			}
			need := (strength[b] + power - 1) / power
			nmask := mask | (1 << b)
			if dp[mask]+need < dp[nmask] {
				dp[nmask] = dp[mask] + need
			}
		}
	}
	return dp[m-1]
}
```
