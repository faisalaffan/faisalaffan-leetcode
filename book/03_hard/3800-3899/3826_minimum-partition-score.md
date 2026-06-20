# 3826 — Minimum Partition Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumPartitionScore(nums []int, k int) int64
```

> **💡 Hint:** DP with range queries using sparse table.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3826: Minimum Partition Score
// https://leetcode.com/problems/minimum-partition-score/
// Difficulty: Hard
//
// Partition array into exactly k contiguous subarrays.
// Score = sum of (max - min) for each subarray. Minimize score.
//
// Approach: DP with range queries using sparse table.
// dp[i][j] = min score for first i elements partitioned into j subarrays.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minimumPartitionScore([]int{2, 6, 3, 5}, 2))
	// Example 2
	fmt.Println(minimumPartitionScore([]int{4, 1, 7, 3}, 3))
	// Edge: k=1
	fmt.Println(minimumPartitionScore([]int{5, 2, 8}, 1))
	// Edge: k=n
	fmt.Println(minimumPartitionScore([]int{3, 1, 4}, 3))
}

func minimumPartitionScore(nums []int, k int) int64 {
	n := len(nums)
	if k <= 0 || k > n {
		return 0
	}
	if k == 1 {
		mx, mn := nums[0], nums[0]
		for _, v := range nums {
			if v > mx {
				mx = v
			}
			if v < mn {
				mn = v
			}
		}
		return int64(mx - mn)
	}
	if k == n {
		return 0
	}

	// Build sparse table for range min/max
  // Alokasi slice integer
	log := make([]int, n+1)
	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}
	K := log[n] + 1
  // Membuat matriks/slice 2D untuk DP
	stMax := make([][]int, K)
  // Membuat matriks/slice 2D untuk DP
	stMin := make([][]int, K)
  // Range loop: iterasi dengan indeks + nilai
	for i := range stMax {
		stMax[i] = make([]int, n)
		stMin[i] = make([]int, n)
	}
	copy(stMax[0], nums)
	copy(stMin[0], nums)
	for j := 1; j < K; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			stMax[j][i] = max(stMax[j-1][i], stMax[j-1][i+(1<<(j-1))])
			stMin[j][i] = min(stMin[j-1][i], stMin[j-1][i+(1<<(j-1))])
		}
	}

	rangeMax := func(l, r int) int {
		j := log[r-l+1]
		return max(stMax[j][l], stMax[j][r-(1<<j)+1])
	}
	rangeMin := func(l, r int) int {
		j := log[r-l+1]
		return min(stMin[j][l], stMin[j][r-(1<<j)+1])
	}
	rangeScore := func(l, r int) int64 {
		return int64(rangeMax(l, r) - rangeMin(l, r))
	}

	// DP[i][j] = min score for first i elements, j partitions
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int64, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			for l := j - 1; l < i; l++ {
				if dp[l][j-1] == math.MaxInt64 {
					continue
				}
				score := dp[l][j-1] + rangeScore(l, i-1)
				if score < dp[i][j] {
					dp[i][j] = score
				}
			}
		}
	}

	return dp[n][k]
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
