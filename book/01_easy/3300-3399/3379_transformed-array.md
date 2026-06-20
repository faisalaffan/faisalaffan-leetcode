# 3379 — Transformed Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TransformedArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3379: Transformed Array
// https://leetcode.com/problems/transformed-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TransformedArray([]int{3, -2, 1, 1}))
	fmt.Println(TransformedArray([]int{-1, 4, -1}))
}

// TransformedArray constructs a new array where result[i] = nums[(i + nums[i]) mod n], handling negative wrap-around.
// Time: O(n). Space: O(n).
func TransformedArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	result := make([]int, n)
	for i, val := range nums {
		idx := (i + val) % n
		if idx < 0 {
			idx += n
		}
		result[i] = nums[idx]
	}
	return result
}
```
