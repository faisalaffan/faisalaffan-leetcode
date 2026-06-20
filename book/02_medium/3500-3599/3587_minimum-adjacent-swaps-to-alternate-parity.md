# 3587 — Minimum Adjacent Swaps To Alternate Parity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumAdjacentSwapsToAlternateParity(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3587: Minimum Adjacent Swaps to Alternate Parity
// https://leetcode.com/problems/minimum-adjacent-swaps-to-alternate-parity/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumAdjacentSwapsToAlternateParity([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", MinimumAdjacentSwapsToAlternateParity([]int{1, 3, 2, 4}))
	// Test case 3
	fmt.Println("Test 3:", MinimumAdjacentSwapsToAlternateParity([]int{2, 4, 6, 8}))
}

func MinimumAdjacentSwapsToAlternateParity(nums []int) int {
	n := len(nums)
	// Separate even and odd indices
	var evens, odds []int
	for i, v := range nums {
		if v%2 == 0 {
			evens = append(evens, i)
		} else {
			odds = append(odds, i)
		}
	}
	if abs(len(evens)-len(odds)) > 1 {
		return -1
	}

	// Try starting with even or odd
	minSwaps := -1
	if len(evens) >= len(odds) {
		swaps := 0
		ei := 0
		for i := 0; i < n; i += 2 {
			swaps += abs(evens[ei] - i)
			ei++
		}
		minSwaps = swaps
	}
	if len(odds) >= len(evens) {
		swaps := 0
		oi := 0
		for i := 0; i < n; i += 2 {
			swaps += abs(odds[oi] - i)
			oi++
		}
		if minSwaps == -1 || swaps < minSwaps {
			minSwaps = swaps
		}
	}
	return minSwaps
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
