# 0813 — Largest Sum Of Averages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LargestSumOfAverages(nums []int, k int) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum

**Waktu:** O(k * n^2)  |  **Ruang:** O(k * n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

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

  // Matriks 2D
	dp := make([][]float64, n+1)
  // Range loop
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
