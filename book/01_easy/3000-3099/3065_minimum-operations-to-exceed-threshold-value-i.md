# 3065 — Minimum Operations To Exceed Threshold Value I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToExceedThresholdValueI(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3065: Minimum Operations to Exceed Threshold Value I
// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minOperations
	fmt.Println(MinimumOperationsToExceedThresholdValueI([]int{2, 11, 10, 1, 3}, 10)) // 3
	fmt.Println(MinimumOperationsToExceedThresholdValueI([]int{1, 1, 2, 4, 9}, 9))    // 4
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minOperations
func MinimumOperationsToExceedThresholdValueI(nums []int, k int) int {
	count := 0
	for _, v := range nums {
		if v < k {
			count++
		}
	}
	return count
}
```
