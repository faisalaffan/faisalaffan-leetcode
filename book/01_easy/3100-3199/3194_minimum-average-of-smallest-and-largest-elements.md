# 3194 — Minimum Average Of Smallest And Largest Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumAverageOfSmallestAndLargestElements(nums []int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n). Space: O(1) (or O(n) due to sorting).  
**Kompleksitas Ruang:** O(1) (or O(n) due to sorting).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3194: Minimum Average of Smallest and Largest Elements
// https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/
// Difficulty: Easy

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(MinimumAverageOfSmallestAndLargestElements([]int{1, 9, 8, 3, 10, 5}))
	fmt.Println(MinimumAverageOfSmallestAndLargestElements([]int{1, 2, 3, 7, 8, 9}))
}

// MinimumAverageOfSmallestAndLargestElements returns the minimum average of the smallest and largest elements.
// Time: O(n log n). Space: O(1) (or O(n) due to sorting).
func MinimumAverageOfSmallestAndLargestElements(nums []int) float64 {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
	minAvg := math.MaxFloat64
	for i := 0; i < n/2; i++ {
		avg := float64(nums[i]+nums[n-1-i]) / 2.0
		if avg < minAvg {
			minAvg = avg
		}
	}
	return minAvg
}
```
