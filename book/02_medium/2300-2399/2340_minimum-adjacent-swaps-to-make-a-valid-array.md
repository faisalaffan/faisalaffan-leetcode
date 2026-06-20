# 2340 — Minimum Adjacent Swaps To Make A Valid Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSwaps(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2340: Minimum Adjacent Swaps to Make a Valid Array
// https://leetcode.com/problems/minimum-adjacent-swaps-to-make-a-valid-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumSwaps(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	minIdx, maxIdx := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[i] >= nums[maxIdx] {
			maxIdx = i
		}
	}

	swaps := minIdx + (n - 1 - maxIdx)
	if minIdx > maxIdx {
		swaps--
	}
	return swaps
}

func main() {
	// Test case 1
	fmt.Println(minimumSwaps([]int{3, 4, 5, 5, 3, 1}))
	// Expected: 6

	// Test case 2
	fmt.Println(minimumSwaps([]int{1, 2, 3, 4}))
	// Expected: 0

	// Test case 3
	fmt.Println(minimumSwaps([]int{2, 1}))
	// Expected: 1
}
```
