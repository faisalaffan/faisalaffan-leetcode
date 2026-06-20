# 2624 — Snail Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func snailTraversal(arr []int, rowsCount int, colsCount int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n*m)  
**Kompleksitas Ruang:** O(n*m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2624: Snail Traversal
// https://leetcode.com/problems/snail-traversal/
// Difficulty: Medium
// Time: O(n*m) | Space: O(n*m)

import "fmt"

func snailTraversal(arr []int, rowsCount int, colsCount int) [][]int {
	n := len(arr)
	if rowsCount*colsCount != n {
		return [][]int{}
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, rowsCount)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = make([]int, colsCount)
	}

	idx := 0
	for col := 0; col < colsCount; col++ {
		if col%2 == 0 {
			// Top to bottom
			for row := 0; row < rowsCount; row++ {
				result[row][col] = arr[idx]
				idx++
			}
		} else {
			// Bottom to top
			for row := rowsCount - 1; row >= 0; row-- {
				result[row][col] = arr[idx]
				idx++
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", snailTraversal([]int{1, 2, 3, 4}, 2, 2))
	// Expected: [[1,4],[2,3]] (snail fill)

	// Test case 2
	fmt.Println("Test 2:", snailTraversal([]int{1, 2, 3, 4, 5, 6}, 2, 3))
	// Expected: [[1,4,5],[2,3,6]]

	// Test case 3: invalid dimensions
	fmt.Println("Test 3:", snailTraversal([]int{1, 2, 3}, 2, 2))
	// Expected: []
}
```
