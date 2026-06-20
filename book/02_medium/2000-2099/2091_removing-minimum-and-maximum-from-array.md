# 2091 — Removing Minimum And Maximum From Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumDeletions(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2091: Removing Minimum and Maximum From Array
// https://leetcode.com/problems/removing-minimum-and-maximum-from-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumDeletions(nums []int) int {
	n := len(nums)
	minIdx, maxIdx := 0, 0
	minVal, maxVal := nums[0], nums[0]

	for i, v := range nums {
		if v < minVal {
			minVal = v
			minIdx = i
		}
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}

	// Possible strategies:
	// 1. Delete both from front: max(minIdx, maxIdx) + 1
	// 2. Delete both from back: n - min(minIdx, maxIdx)
	// 3. Delete one from front, one from back: min(minIdx, maxIdx) + 1 + n - max(minIdx, maxIdx)
	front := max(minIdx, maxIdx) + 1
	back := n - min(minIdx, maxIdx)
	mixed := min(minIdx, maxIdx) + 1 + n - max(minIdx, maxIdx)

	return min(front, min(back, mixed))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6}))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", minimumDeletions([]int{0, -4, 19, 1, 8, -2, -3, 5}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", minimumDeletions([]int{101}))
	// Expected: 1
}
```
