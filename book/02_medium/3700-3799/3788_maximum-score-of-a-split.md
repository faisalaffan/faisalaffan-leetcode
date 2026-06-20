# 3788 — Maximum Score Of A Split

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumScoreOfASplit(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3788: Maximum Score of a Split
// https://leetcode.com/problems/maximum-score-of-a-split/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func maximumScoreOfASplit(nums []int) int64 {
	n := len(nums)
  // Alokasi slice integer
	suf := make([]int64, n)
	suf[n-1] = int64(nums[n-1])
	for i := n - 2; i >= 0; i-- {
		suf[i] = int64(nums[i])
		if suf[i+1] < suf[i] {
			suf[i] = suf[i+1]
		}
	}

	var pre int64
	var ans int64 = math.MinInt64
	for i := 0; i < n-1; i++ {
		pre += int64(nums[i])
		score := pre - suf[i+1]
		if score > ans {
			ans = score
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumScoreOfASplit([]int{10, -1, 3, -4, -5}))
	fmt.Println(maximumScoreOfASplit([]int{1, 2, 3, 4}))
	fmt.Println(maximumScoreOfASplit([]int{-5, -3, -1}))
}
```
