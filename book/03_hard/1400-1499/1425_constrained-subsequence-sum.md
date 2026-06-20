# 1425 — Constrained Subsequence Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func constrainedSubsetSum(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window, Dynamic Programming, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1425: Constrained Subsequence Sum
// https://leetcode.com/problems/constrained-subsequence-sum/
// Difficulty: Hard

import "fmt"

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice integer
	dp := make([]int, n)
	// Monotonic deque storing indices, decreasing dp values
  // Alokasi slice integer
	deque := make([]int, 0, n)
	ans := nums[0]

	for i := 0; i < n; i++ {
		// Remove indices out of window
		for len(deque) > 0 && deque[0] < i-k {
			deque = deque[1:]
		}

		dp[i] = nums[i]
		if len(deque) > 0 {
			// max dp in window
			if dp[deque[0]] > 0 {
				dp[i] += dp[deque[0]]
			}
		}

		if dp[i] > ans {
			ans = dp[i]
		}

		// Maintain decreasing deque
		for len(deque) > 0 && dp[deque[len(deque)-1]] <= dp[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	return ans
}

func main() {
	// Example: [10,2,-10,5,20], 2 -> 37
	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
}
```
