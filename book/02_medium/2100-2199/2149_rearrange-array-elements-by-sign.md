# 2149 — Rearrange Array Elements By Sign

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rearrangeArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2149: Rearrange Array Elements by Sign
// https://leetcode.com/problems/rearrange-array-elements-by-sign/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func rearrangeArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	result := make([]int, n)
	posIdx, negIdx := 0, 1

	for _, v := range nums {
		if v > 0 {
			result[posIdx] = v
			posIdx += 2
		} else {
			result[negIdx] = v
			negIdx += 2
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	// Expected: [3,-2,1,-5,2,-4]

	// Test case 2
	fmt.Println("Test 2:", rearrangeArray([]int{-1, 1}))
	// Expected: [1,-1]

	// Test case 3
	fmt.Println("Test 3:", rearrangeArray([]int{28, -41, 22, -8, -37, 46, 35, -9}))
	// Expected: [28,-41,22,-8,46,-37,35,-9]
}
```
