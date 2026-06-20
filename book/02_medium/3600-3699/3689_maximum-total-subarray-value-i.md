# 3689 — Maximum Total Subarray Value I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumTotalSubarrayValueI(nums []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3689: Maximum Total Subarray Value I
// https://leetcode.com/problems/maximum-total-subarray-value-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"
import "math"

func maximumTotalSubarrayValueI(nums []int, k int) int64 {
	minVal := math.MaxInt32
	maxVal := math.MinInt32
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return int64(maxVal-minVal) * int64(k)
}

func main() {
	fmt.Println(maximumTotalSubarrayValueI([]int{1, 3, 2}, 2))
	fmt.Println(maximumTotalSubarrayValueI([]int{4, 2, 5, 1}, 3))
	fmt.Println(maximumTotalSubarrayValueI([]int{1, 1, 1}, 5))
}
```
