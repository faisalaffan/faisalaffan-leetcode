# 3467 — Transform Array By Parity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TransformArrayByParity(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3467: Transform Array by Parity
// https://leetcode.com/problems/transform-array-by-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TransformArrayByParity([]int{4, 3, 2, 1}))
	fmt.Println(TransformArrayByParity([]int{1, 5, 2, 8, 3}))
}

// TransformArrayByParity transforms array: even numbers -> 0 (sorted first), odd numbers -> 1.
// Time: O(n log n). Space: O(1).
func TransformArrayByParity(nums []int) []int {
	// Count evens
	evenCount := 0
	for _, v := range nums {
		if v%2 == 0 {
			evenCount++
		}
	}
  // Alokasi slice integer
	result := make([]int, len(nums))
	for i := 0; i < evenCount; i++ {
		result[i] = 0
	}
	for i := evenCount; i < len(nums); i++ {
		result[i] = 1
	}
	return result
}
```
