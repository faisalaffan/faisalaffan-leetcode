# 3936 — Minimum Swaps To Move Zeros To End

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumSwapsToMoveZerosToEnd(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3936: Minimum Swaps to Move Zeros to End
// https://leetcode.com/problems/minimum-swaps-to-move-zeros-to-end/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumSwapsToMoveZerosToEnd([]int{0, 1, 0, 3, 12}))
	fmt.Println(MinimumSwapsToMoveZerosToEnd([]int{0, 1, 0, 2}))
	fmt.Println(MinimumSwapsToMoveZerosToEnd([]int{1, 2, 0}))
}

// Time: O(n)
// Space: O(1)
func MinimumSwapsToMoveZerosToEnd(nums []int) int {
	zeroCount := 0
	for _, v := range nums {
		if v == 0 {
			zeroCount++
		}
	}
	boundary := len(nums) - zeroCount
	swaps := 0
	for i := 0; i < boundary; i++ {
		if nums[i] == 0 {
			swaps++
		}
	}
	return swaps
}
```
