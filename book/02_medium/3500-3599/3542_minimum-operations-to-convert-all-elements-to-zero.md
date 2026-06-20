# 3542 — Minimum Operations To Convert All Elements To Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToConvertAllElementsToZero(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3542: Minimum Operations to Convert All Elements to Zero
// https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumOperationsToConvertAllElementsToZero([]int{1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", MinimumOperationsToConvertAllElementsToZero([]int{0, 0, 0}))
	// Test case 3
	fmt.Println("Test 3:", MinimumOperationsToConvertAllElementsToZero([]int{1, 1, 1}))
}

func MinimumOperationsToConvertAllElementsToZero(nums []int) int {
	ops := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			// flip from i onward (or flip just this element)
			ops++
		}
	}
	return ops
}
```
