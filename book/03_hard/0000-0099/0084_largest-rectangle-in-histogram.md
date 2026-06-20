# 0084 — Largest Rectangle In Histogram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func largestRectangleArea(heights []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #84: Largest Rectangle in Histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("84. Largest Rectangle in Histogram")
	fmt.Println("[2,1,5,6,2,3] ->", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), "(expected 10)")
	fmt.Println("[2,4] ->", largestRectangleArea([]int{2, 4}), "(expected 4)")
	fmt.Println("[2,1,2] ->", largestRectangleArea([]int{2, 1, 2}), "(expected 3)")
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
