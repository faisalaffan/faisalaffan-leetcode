# 2592 — Maximize Greatness Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximizeGreatness(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2592: Maximize Greatness of an Array
// https://leetcode.com/problems/maximize-greatness-of-an-array/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximizeGreatness(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	j := 0
	for _, v := range nums {
		if v > nums[j] {
			j++
		}
	}
	return j
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maximizeGreatness([]int{1, 2, 3, 4}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", maximizeGreatness([]int{1, 1, 1}))
	// Expected: 0
}
```
