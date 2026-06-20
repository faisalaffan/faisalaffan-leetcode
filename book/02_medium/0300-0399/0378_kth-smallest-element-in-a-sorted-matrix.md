# 0378 — Kth Smallest Element In A Sorted Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func kthSmallest(matrix [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O((m+n) * log(max-min))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #378: Kth Smallest Element in a Sorted Matrix
// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
// Difficulty: Medium
// Time: O((m+n) * log(max-min)) | Space: O(1)

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	lo, hi := matrix[0][0], matrix[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		row, col := n-1, 0
		for row >= 0 && col < n {
			if matrix[row][col] <= mid {
				count += row + 1
				col++
			} else {
				row--
			}
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	// Test case 1
	matrix1 := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	fmt.Println("Test 1:", kthSmallest(matrix1, 8))
	// Expected: 13

	// Test case 2
	fmt.Println("Test 2:", kthSmallest([][]int{{-5}}, 1))
	// Expected: -5

	// Test case 3
	matrix3 := [][]int{{1, 2}, {1, 3}}
	fmt.Println("Test 3:", kthSmallest(matrix3, 2))
	// Expected: 1
}
```
