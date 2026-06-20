# 3511 — Make A Positive Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MakeAPositiveArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3511: Make a Positive Array
// https://leetcode.com/problems/make-a-positive-array/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MakeAPositiveArray([]int{-1, 2, -3, 4}))
	// Test case 2
	fmt.Println("Test 2:", MakeAPositiveArray([]int{1, 2, 3}))
	// Test case 3
	fmt.Println("Test 3:", MakeAPositiveArray([]int{-5, -10}))
}

func MakeAPositiveArray(nums []int) int {
	// Minimum operations to make all elements positive
	// Each operation can increment an element by 1
	ops := 0
	for _, v := range nums {
		if v <= 0 {
			ops += -v + 1
		}
	}
	return ops
}
```
