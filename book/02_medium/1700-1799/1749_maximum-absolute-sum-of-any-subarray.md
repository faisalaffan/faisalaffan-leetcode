# 1749 — Maximum Absolute Sum Of Any Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxAbsoluteSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1749: Maximum Absolute Sum of Any Subarray
// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxAbsoluteSum(nums []int) int {
	maxEnding := 0
	minEnding := 0
	maxSoFar := 0

	for _, v := range nums {
		maxEnding = max(0, maxEnding+v)
		minEnding = min(0, minEnding+v)
		maxSoFar = max(maxSoFar, maxEnding, -minEnding)
	}
	return maxSoFar
}

func max(nums ...int) int {
	r := nums[0]
	for _, v := range nums[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4})) // Expected: 5
	fmt.Println(maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2})) // Expected: 8
	fmt.Println(maxAbsoluteSum([]int{-1})) // Expected: 1
}
```
