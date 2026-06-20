# 3248 — Snake In Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func SnakeInMatrix(n int, commands []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3248: Snake in Matrix
// https://leetcode.com/problems/snake-in-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SnakeInMatrix(3, []string{"RIGHT", "DOWN"}))
	fmt.Println(SnakeInMatrix(2, []string{"DOWN", "RIGHT", "UP"}))
}

// SnakeInMatrix returns the final position of the snake in an n x n matrix after following commands.
// Time: O(m). Space: O(1).
func SnakeInMatrix(n int, commands []string) int {
	r, c := 0, 0
	for _, cmd := range commands {
		switch cmd {
		case "UP":
			r--
		case "DOWN":
			r++
		case "LEFT":
			c--
		case "RIGHT":
			c++
		}
	}
	return r*n + c
}
```
