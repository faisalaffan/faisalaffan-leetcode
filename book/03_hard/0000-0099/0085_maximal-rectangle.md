# 0085 — Maximal Rectangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximalRectangle(matrix [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #85: Maximal Rectangle
// https://leetcode.com/problems/maximal-rectangle/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("85. Maximal Rectangle")
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println("Example 1 ->", maximalRectangle(matrix), "(expected 6)")

	matrix2 := [][]byte{{'0'}}
	fmt.Println("Example 2 ->", maximalRectangle(matrix2), "(expected 0)")

	matrix3 := [][]byte{{'1'}}
	fmt.Println("Example 3 ->", maximalRectangle(matrix3), "(expected 1)")
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
  // Alokasi slice integer
	heights := make([]int, cols)
	maxArea := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == '1' {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}
		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
  // Alokasi slice integer
	stack := make([]int, 0, n)
	maxArea := 0

	for i := 0; i <= n; i++ {
		var h int
		if i == n {
			h = 0
		} else {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}
```
