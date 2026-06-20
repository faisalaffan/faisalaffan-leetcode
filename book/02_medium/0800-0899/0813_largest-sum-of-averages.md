# 0813 — Largest Sum Of Averages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestSumOfAverages(nums []int, k int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(k * n^2)  
**Kompleksitas Ruang:** O(k * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #813: Largest Sum of Averages
// https://leetcode.com/problems/largest-sum-of-averages/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LargestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
	fmt.Println(LargestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4))
	fmt.Println(LargestSumOfAverages([]int{4, 1, 7, 5, 6, 2, 3}, 4))
}

// Time: O(k * n^2) | Space: O(k * n)
func LargestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	prefix := make([]float64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + float64(nums[i])
	}

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]float64, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}

	for i := 1; i <= n; i++ {
		dp[i][1] = prefix[i] / float64(i)
	}

	for j := 2; j <= k; j++ {
		for i := j; i <= n; i++ {
			var best float64
			for x := j - 1; x < i; x++ {
				avg := (prefix[i] - prefix[x]) / float64(i-x)
				val := dp[x][j-1] + avg
				if val > best {
					best = val
				}
			}
			dp[i][j] = best
		}
	}

	return dp[n][k]
}
```
