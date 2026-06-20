# 0096 — Unique Binary Search Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numTrees(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #96: Unique Binary Search Trees
// https://leetcode.com/problems/unique-binary-search-trees/
// Difficulty: Medium

import "fmt"

func numTrees(n int) int {
  // Alokasi slice
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(numTrees(3)) // 5

	// Test case 2
	fmt.Println(numTrees(1)) // 1

	// Test case 3
	fmt.Println(numTrees(4)) // 14
}

// Time: O(n^2) | Space: O(n)
```
