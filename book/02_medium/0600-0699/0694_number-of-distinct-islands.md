# 0694 — Number Of Distinct Islands

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numDistinctIslands(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(R * C)  
**Kompleksitas Ruang:** O(R * C)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #694: Number of Distinct Islands
// https://leetcode.com/problems/number-of-distinct-islands/
// Difficulty: Medium [Paid]
// Time: O(R * C)
// Space: O(R * C)

import (
	"fmt"
	"strings"
)

func main() {
	grid1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(numDistinctIslands(grid1))

	grid2 := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}
	fmt.Println(numDistinctIslands(grid2))
}

func numDistinctIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
  // Membuat map (HashMap) — pencarian O(1)
	shapes := make(map[string]bool)

	var dfs func(r, c int, dir byte, sb *strings.Builder)
	dfs = func(r, c int, dir byte, sb *strings.Builder) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
			return
		}
		grid[r][c] = 0
		sb.WriteByte(dir)
		dfs(r-1, c, 'U', sb)
		dfs(r+1, c, 'D', sb)
		dfs(r, c-1, 'L', sb)
		dfs(r, c+1, 'R', sb)
		sb.WriteByte('B')
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				var sb strings.Builder
				dfs(r, c, 'S', &sb)
				shapes[sb.String()] = true
			}
		}
	}

	return len(shapes)
}
```
