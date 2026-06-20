# 3929 — Minimum Partition Score Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumPartitionScore(nums []int, k int) int64
```

> **💡 Hint:** DP with sorted array. After sorting, each partition

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3929: Minimum Partition Score II
// https://leetcode.com/problems/minimum-partition-score-ii/
// Difficulty: Hard [Paid]
//
// Partition array into exactly k subarrays. Score of subarray =
// (max - min)^2. Minimize total score.
//
// Approach: DP with sorted array. After sorting, each partition
// is a contiguous segment. Score of segment = (last - first)^2.
// dp[i][j] = min score for first i elements into j partitions.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minimumPartitionScore([]int{3, 1, 4, 1, 5}, 2))
	// Example 2
	fmt.Println(minimumPartitionScore([]int{10, 20, 30}, 3))
	// Edge: k = 1
	fmt.Println(minimumPartitionScore([]int{5, 2, 8}, 1))
}

func minimumPartitionScore(nums []int, k int) int64 {
	n := len(nums)
	if k <= 0 || k > n {
		return 0
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)

	// dp[i][j] = min score for first i elements into j partitions
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			for l := j - 1; l < i; l++ {
				if dp[l][j-1] == math.MaxInt32 {
					continue
				}
				score := (nums[i-1] - nums[l]) * (nums[i-1] - nums[l])
				if dp[l][j-1]+score < dp[i][j] {
					dp[i][j] = dp[l][j-1] + score
				}
			}
		}
	}

	return int64(dp[n][k])
}
```
