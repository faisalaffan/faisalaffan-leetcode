# 3774 — Absolute Difference Between Maximum And Minimum K Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AbsoluteDifferenceBetweenMaximumAndMinimumKElements(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3774: Absolute Difference Between Maximum and Minimum K Elements
// https://leetcode.com/problems/absolute-difference-between-maximum-and-minimum-k-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AbsoluteDifferenceBetweenMaximumAndMinimumKElements([]int{5, 2, 2, 4}, 2))
	fmt.Println(AbsoluteDifferenceBetweenMaximumAndMinimumKElements([]int{100}, 1))
}

// Time: O(n log n)
// Space: O(1)
func AbsoluteDifferenceBetweenMaximumAndMinimumKElements(nums []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
	maxSum := 0
	minSum := 0
	for i := 0; i < k; i++ {
		minSum += nums[i]
		maxSum += nums[n-1-i]
	}
	return maxSum - minSum
}
```
