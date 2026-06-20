# 1034 — Coloring A Border

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func colorBorder(grid [][]int, row int, col int, color int) [][]int
```

> **💡 Hint:** DFS to find connected component, then color border cells

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1034: Coloring A Border
// https://leetcode.com/problems/coloring-a-border/
// Difficulty: Medium
//
// Approach: DFS to find connected component, then color border cells
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{{1, 1}, {1, 2}}, 0, 0, 3)) // [[3,3],[3,2]]
	fmt.Println(colorBorder([][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3)) // [[1,3,3],[2,3,3]]
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	origColor := grid[row][col]
	if origColor == color {
		return grid
	}

  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		visited[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] != origColor {
				grid[r][c] = color
			} else if !visited[nr][nc] {
				dfs(nr, nc)
			}
		}
	}

	dfs(row, col)
	return grid
}
```
