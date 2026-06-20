# 0453 — Minimum Moves To Equal Array Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minMoves(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #453: Minimum Moves to Equal Array Elements
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func minMoves(nums []int) int {
	minVal := math.MaxInt32
	sum := 0
	for _, n := range nums {
		sum += n
		if n < minVal {
			minVal = n
		}
	}
	return sum - minVal*len(nums)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves([]int{1, 2, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minMoves([]int{1, 1, 1}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minMoves([]int{1, 1000000000}))
	// Expected: 999999999
}
```
