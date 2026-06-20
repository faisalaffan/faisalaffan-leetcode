# 2708 — Maximum Strength Of A Group

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxStrength(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2708: Maximum Strength of a Group
// https://leetcode.com/problems/maximum-strength-of-a-group/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxStrength(nums []int) int64 {
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}

	positives := []int{}
	negatives := []int{}
	hasZero := false

	for _, v := range nums {
		if v > 0 {
			positives = append(positives, v)
		} else if v < 0 {
			negatives = append(negatives, v)
		} else {
			hasZero = true
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(negatives) // Sort negatives (most negative first)

	ans := int64(1)
	for _, v := range positives {
		ans *= int64(v)
	}

	// Take pairs of negatives (the largest ones, i.e., closest to 0)
	// Sort negatives ascending: [-5, -4, -3, -2, -1]
	// We want to pair them from the least negative end: [-1, -2] * [-3, -4] ...
	if len(negatives)%2 == 1 {
		negatives = negatives[:len(negatives)-1] // Drop the most negative if odd count
	}

	for _, v := range negatives {
		ans *= int64(v)
	}

	if ans == 1 && len(positives) == 0 && len(negatives) == 0 {
		if hasZero {
			return 0
		}
	}

	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxStrength([]int{3, -1, -5, 2, 5, -9}))
	// Expected: 1350

	// Test case 2
	fmt.Println("Test 2:", maxStrength([]int{-4, -5, -6}))
	// Expected: 30

	// Test case 3
	fmt.Println("Test 3:", maxStrength([]int{0, -1}))
	// Expected: 0
}
```
