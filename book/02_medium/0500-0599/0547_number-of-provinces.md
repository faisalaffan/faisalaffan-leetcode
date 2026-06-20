# 0547 — Number Of Provinces

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindCircleNum(isConnected [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #547: Number of Provinces
// https://leetcode.com/problems/number-of-provinces/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	isConnected1 := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(FindCircleNum(isConnected1))
	isConnected2 := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(FindCircleNum(isConnected2))
}

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0

	var dfs func(city int)
	dfs = func(city int) {
		visited[city] = true
		for neighbor := 0; neighbor < n; neighbor++ {
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}

	return count
}
```
