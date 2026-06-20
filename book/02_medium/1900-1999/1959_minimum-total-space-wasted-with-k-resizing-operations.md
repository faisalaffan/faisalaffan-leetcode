# 1959 — Minimum Total Space Wasted With K Resizing Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinSpaceWastedKResizing(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(k * n^2), Space: O(k * n)  
**Kompleksitas Ruang:** O(k * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1959: Minimum Total Space Wasted With K Resizing Operations
// https://leetcode.com/problems/minimum-total-space-wasted-with-k-resizing-operations/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSpaceWastedKResizing([]int{10, 20, 30}, 1))
	fmt.Println(MinSpaceWastedKResizing([]int{10, 20, 15, 30, 20}, 2))
}

const INF = 1 << 30

// Time: O(k * n^2), Space: O(k * n)
func MinSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	k++ // k resizes = k+1 segments

	// g[i][j] = wasted space for segment nums[i..j]
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		s, mx := 0, 0
		for j := i; j < n; j++ {
			s += nums[j]
			if nums[j] > mx {
				mx = nums[j]
			}
			g[i][j] = mx*(j-i+1) - s
		}
	}

	// dp[i][j] = min waste for first i elements with j segments
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for h := 0; h < i; h++ {
				val := dp[h][j-1] + g[h][i-1]
				if val < dp[i][j] {
					dp[i][j] = val
				}
			}
		}
	}
	return dp[n][k]
}
```
