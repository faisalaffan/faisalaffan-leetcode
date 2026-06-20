# 1984 — Minimum Difference Between Highest And Lowest Of K Scores

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumDifferenceBetweenHighestAndLowestOfKScores(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) ignoring sort  
**Kompleksitas Ruang:** O(1) ignoring sort

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1984: Minimum Difference Between Highest and Lowest of K Scores
// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
// Difficulty: Easy

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{90}, 1))            // 0
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{9, 4, 1, 7}, 2))    // 2
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{9, 4, 1, 7, 5, 3}, 3)) // 3
}

// Time: O(n log n), Space: O(1) ignoring sort
func MinimumDifferenceBetweenHighestAndLowestOfKScores(nums []int, k int) int {
	if k == 1 {
		return 0
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	minDiff := math.MaxInt32
	for i := 0; i <= len(nums)-k; i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}
```
