# 1690 — Stone Game Vii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func stoneGameVII(stones []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP, Prefix Sum

**Waktu:** O(n^2), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1690: Stone Game VII
// https://leetcode.com/problems/stone-game-vii/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import "fmt"

func stoneGameVII(stones []int) int {
	n := len(stones)
  // Alokasi slice
	prefix := make([]int, n+1)
	for i, v := range stones {
		prefix[i+1] = prefix[i] + v
	}

	// dp[i][j] = max score difference (current player - other player) for subarray i..j
  // Alokasi slice
	dp := make([]int, n)

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			// Remove left: score = sum of rest, then subtract opponent's optimal
			removeLeft := (prefix[j+1] - prefix[i+1]) - dp[i+1]
			// Remove right: score = sum of rest, then subtract opponent's optimal
			removeRight := (prefix[j] - prefix[i]) - dp[i]
			dp[i] = max(removeLeft, removeRight)
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2})) // Expected: 6
	fmt.Println(stoneGameVII([]int{7, 90, 5, 1, 100, 10, 10, 2})) // Expected: 122
	fmt.Println(stoneGameVII([]int{1, 2})) // Expected: 2
}
```
