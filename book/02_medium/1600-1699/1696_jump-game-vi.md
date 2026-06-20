# 1696 — Jump Game Vi

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxResult(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Monotonic Stack

**Waktu:** O(n), Space: O(k)  |  **Ruang:** O(k)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1696: Jump Game VI
// https://leetcode.com/problems/jump-game-vi/
// Difficulty: Medium
// Time: O(n), Space: O(k)

import "fmt"

func maxResult(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice
	dp := make([]int, n)
	dp[0] = nums[0]

	// Monotonic deque storing indices with decreasing dp values
  // Alokasi slice
	deque := make([]int, 0, n)
	deque = append(deque, 0)

	for i := 1; i < n; i++ {
		// Remove out-of-range elements
		for len(deque) > 0 && deque[0] < i-k {
			deque = deque[1:]
		}

		// dp[i] = nums[i] + max dp from i-k to i-1
		dp[i] = nums[i] + dp[deque[0]]

		// Maintain decreasing order
		for len(deque) > 0 && dp[deque[len(deque)-1]] <= dp[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	return dp[n-1]
}

func main() {
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2)) // Expected: 7
	fmt.Println(maxResult([]int{10, -5, -2, 4, 0, 3}, 3)) // Expected: 17
	fmt.Println(maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2)) // Expected: 0
}
```
