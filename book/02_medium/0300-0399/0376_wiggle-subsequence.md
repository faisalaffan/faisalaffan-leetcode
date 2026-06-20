# 0376 — Wiggle Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func wiggleMaxLength(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #376: Wiggle Subsequence
// https://leetcode.com/problems/wiggle-subsequence/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// Expected: 2
}
```
