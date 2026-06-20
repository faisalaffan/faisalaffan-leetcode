# 0947 — Most Stones Removed With Same Row Or Column

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func removeStones(stones [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Union-Find

**Waktu:** O(n * α(n))  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #947: Most Stones Removed with Same Row or Column
// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
// Difficulty: Medium

import "fmt"

// Time: O(n * α(n)) | Space: O(n)
func removeStones(stones [][]int) int {
	n := len(stones)
  // Alokasi slice
	parent := make([]int, n)
  // Range loop
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = ra
		}
	}

	// Union stones sharing row or column
  // HashMap: O(1) lookup
	rowMap := make(map[int]int)
  // HashMap: O(1) lookup
	colMap := make(map[int]int)
	for i, s := range stones {
		if v, ok := rowMap[s[0]]; ok {
			union(i, v)
		} else {
			rowMap[s[0]] = i
		}
		if v, ok := colMap[s[1]]; ok {
			union(i, v)
		} else {
			colMap[s[1]] = i
		}
	}

	// Count connected components
	components := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			components++
		}
	}

	return n - components
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}}))
}
```
