# 3147 — Taking Maximum Energy From The Mystic Dungeon

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumEnergy(energy []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n)  |  **Ruang:** O(k)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3147: Taking Maximum Energy From the Mystic Dungeon
// https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/
// Difficulty: Medium
// Time: O(n) | Space: O(k)

import "fmt"

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
  // Alokasi slice
	dp := make([]int, n)
  // Range loop
	for i := range dp {
		dp[i] = energy[i]
	}

	for i := k; i < n; i++ {
		if dp[i-k] > 0 {
			dp[i] += dp[i-k]
		}
	}

	ans := dp[n-1]
	for i := n - k - 1; i >= 0; i -= k {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumEnergy([]int{5, 2, -10, -5, 1}, 3)) // Expected: 3
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2))        // Expected: -1
}
```
