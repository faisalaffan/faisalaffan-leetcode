# 3248 — Snake In Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func SnakeInMatrix(n int, commands []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(m). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

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
