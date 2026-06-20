# 3819 — Rotate Non Negative Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RotateNonNegativeElements(nums []int, k int) []int
```

> **💡 Hint:** Collect non-negative elements, rotate left by k cyclically, place back.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3819: Rotate Non Negative Elements
// https://leetcode.com/problems/rotate-non-negative-elements/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Collect non-negative elements, rotate left by k cyclically, place back.

import "fmt"

func RotateNonNegativeElements(nums []int, k int) []int {
	n := len(nums)
	type pair struct {
		idx int
		val int
	}
	var nonNeg []pair
	for i, v := range nums {
		if v >= 0 {
			nonNeg = append(nonNeg, pair{i, v})
		}
	}

	m := len(nonNeg)
	if m == 0 {
  // Alokasi slice integer
		res := make([]int, n)
		copy(res, nums)
		return res
	}

  // Alokasi slice integer
	result := make([]int, n)
	copy(result, nums)

	// For each position that had a non-negative, put the rotated value
	for i, p := range nonNeg {
		srcIdx := (i + k) % m
		result[p.idx] = nonNeg[srcIdx].val
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(RotateNonNegativeElements([]int{1, -2, 3, -4}, 3)) // Expected: [3 -2 1 -4]

	// Example 2
	fmt.Println(RotateNonNegativeElements([]int{-3, -2, 7}, 1)) // Expected: [-3 -2 7]

	// Example 3
	fmt.Println(RotateNonNegativeElements([]int{5, 4, -9, 6}, 2)) // Expected: [6 5 -9 4]
}
```
