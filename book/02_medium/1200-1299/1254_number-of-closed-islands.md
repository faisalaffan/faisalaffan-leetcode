# 1254 — Number Of Closed Islands

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func closedIsland(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n) recursion depth

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1254: Number of Closed Islands
// https://leetcode.com/problems/number-of-closed-islands/
// Difficulty: Medium

// A closed island is a group of 0s not connected to the border.
// First mark all border-connected 0s, then count remaining islands.

// Time: O(m*n)
// Space: O(m*n) recursion depth

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != 0 {
			return
		}
		grid[r][c] = 1 // mark as water/visited
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// Mark border-connected 0s
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	count := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		closedIsland([][]int{
			{1, 1, 1, 1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0, 1, 1, 0},
			{1, 0, 1, 0, 1, 1, 1, 0},
			{1, 0, 0, 0, 0, 1, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 0},
		}))

	fmt.Printf("%d (expected: 0)\n",
		closedIsland([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
```
