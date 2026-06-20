# 3196 — Maximize Total Cost Of Alternating Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumTotalCost(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3196: Maximize Total Cost of Alternating Subarrays
// https://leetcode.com/problems/maximize-total-cost-of-alternating-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumTotalCost(nums []int) int64 {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}

	add := int64(nums[0])
	sub := int64(nums[0])

	for i := 1; i < len(nums); i++ {
		v := int64(nums[i])
		newAdd := maxInt64(add, sub) + v
		newSub := add - v
		add, sub = newAdd, newSub
	}

	return maxInt64(add, sub)
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumTotalCost([]int{1, -2, 3, 4}))    // Expected: 10
	fmt.Println(maximumTotalCost([]int{1, -1, 1, -1}))   // Expected: 4
}
```
