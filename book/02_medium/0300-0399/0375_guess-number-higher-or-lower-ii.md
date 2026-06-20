# 0375 — Guess Number Higher Or Lower Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func getMoneyAmount(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** O(n^3)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #375: Guess Number Higher or Lower II
// https://leetcode.com/problems/guess-number-higher-or-lower-ii/
// Difficulty: Medium
// Time: O(n^3) | Space: O(n^2)

import "fmt"

func getMoneyAmount(n int) int {
  // Matriks 2D
	dp := make([][]int, n+2)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for length := 2; length <= n; length++ {
		for start := 1; start <= n-length+1; start++ {
			end := start + length - 1
			dp[start][end] = 1<<31 - 1
			for pivot := start; pivot <= end; pivot++ {
				// Cost if pivot is wrong: pivot + max(cost of left, cost of right)
				cost := pivot + max(dp[start][pivot-1], dp[pivot+1][end])
				if cost < dp[start][end] {
					dp[start][end] = cost
				}
			}
		}
	}
	return dp[1][n]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getMoneyAmount(10))
	// Expected: 16

	// Test case 2
	fmt.Println("Test 2:", getMoneyAmount(1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", getMoneyAmount(2))
	// Expected: 1
}
```
