# 1877 — Minimize Maximum Pair Sum In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinPairSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1877: Minimize Maximum Pair Sum in Array
// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinPairSum([]int{3, 5, 2, 3}))
	fmt.Println(MinPairSum([]int{3, 5, 4, 2, 4, 6}))
}

// Time: O(n log n), Space: O(1)
func MinPairSum(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	maxSum := 0
	n := len(nums)
	for i := 0; i < n/2; i++ {
		pairSum := nums[i] + nums[n-1-i]
		if pairSum > maxSum {
			maxSum = pairSum
		}
	}
	return maxSum
}
```
