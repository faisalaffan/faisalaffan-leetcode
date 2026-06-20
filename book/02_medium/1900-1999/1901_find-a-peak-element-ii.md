# 1901 — Find A Peak Element Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindPeakGrid(mat [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(m log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1901: Find a Peak Element II
// https://leetcode.com/problems/find-a-peak-element-ii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindPeakGrid([][]int{{1, 4, 3, 2}, {2, 3, 4, 5}, {5, 4, 3, 6}}))
	fmt.Println(FindPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))
}

// Time: O(m log n), Space: O(1)
func FindPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2
		maxRow := 0
		for i := 0; i < m; i++ {
			if mat[i][mid] > mat[maxRow][mid] {
				maxRow = i
			}
		}

		leftVal := -1
		if mid > 0 {
			leftVal = mat[maxRow][mid-1]
		}
		rightVal := -1
		if mid < n-1 {
			rightVal = mat[maxRow][mid+1]
		}

		if mat[maxRow][mid] > leftVal && mat[maxRow][mid] > rightVal {
			return []int{maxRow, mid}
		} else if mat[maxRow][mid] < leftVal {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}
```
