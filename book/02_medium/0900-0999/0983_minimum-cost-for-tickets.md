# 0983 — Minimum Cost For Tickets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func mincostTickets(days []int, costs []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(n) where n is the range of days (last travel day)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #983: Minimum Cost For Tickets
// https://leetcode.com/problems/minimum-cost-for-tickets/
// Difficulty: Medium
//
// Approach: DP (bottom-up) over travel days
// Time: O(n) where n is the range of days (last travel day)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})) // 11
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15})) // 17
	fmt.Println(mincostTickets([]int{1, 2, 3}, []int{2, 7, 15})) // 6
}

func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
  // Alokasi slice
	dp := make([]int, lastDay+1)
  // HashMap: O(1) lookup
	travelSet := make(map[int]bool)
	for _, d := range days {
		travelSet[d] = true
	}

	for i := 1; i <= lastDay; i++ {
		if !travelSet[i] {
			dp[i] = dp[i-1]
			continue
		}
		one := dp[i-1] + costs[0]
		seven := dp[max(0, i-7)] + costs[1]
		thirty := dp[max(0, i-30)] + costs[2]
		dp[i] = min(one, min(seven, thirty))
	}

	return dp[lastDay]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
