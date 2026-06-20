# 1591 — Strange Printer Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isPrintable(targetGrid [][]int) bool
```

> **💡 Hint:** // 1. For each color (1..60), find its bounding box (min/max row and col).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS, Topological Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1591: Strange Printer II
// https://leetcode.com/problems/strange-printer-ii/
// Difficulty: Hard
//
// The printer paints axis-aligned rectangles of a single color,
// one at a time, in some order. Later rectangles may cover earlier ones.
// Determine if the target grid can be produced.
//
// Approach:
// 1. For each color (1..60), find its bounding box (min/max row and col).
// 2. For each cell inside the bounding box of color c, if the cell's color
//    is different from c, that other color MUST have been painted AFTER c
//    (to override c's paint). This defines a directed edge: c -> other.
// 3. Build a directed graph and check if there is a cycle. If no cycle,
//    a topological order exists, meaning the grid is producible.

func isPrintable(targetGrid [][]int) bool {
	m := len(targetGrid)
	if m == 0 {
		return true
	}
	n := len(targetGrid[0])

	// Bounding boxes for each color (1-based, max 60 colors)
	type bbox struct {
		minR, maxR int
		minC, maxC int
	}
  // Membuat map (HashMap) — pencarian O(1)
	boxes := make(map[int]*bbox)
  // Membuat map (HashMap) — pencarian O(1)
	hasColor := make(map[int]bool)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			color := targetGrid[r][c]
			hasColor[color] = true
			if _, ok := boxes[color]; !ok {
				boxes[color] = &bbox{minR: r, maxR: r, minC: c, maxC: c}
			} else {
				b := boxes[color]
				if r < b.minR {
					b.minR = r
				}
				if r > b.maxR {
					b.maxR = r
				}
				if c < b.minC {
					b.minC = c
				}
				if c > b.maxC {
					b.maxC = c
				}
			}
		}
	}

	// Build adjacency: if color a's bounding box contains a cell of color b (b != a),
	// then a must be painted before b (edge a -> b).
  // Membuat map (HashMap) — pencarian O(1)
	graph := make(map[int]map[int]bool)
	for color, b := range boxes {
		if graph[color] == nil {
			graph[color] = make(map[int]bool)
		}
		for r := b.minR; r <= b.maxR; r++ {
			for c := b.minC; c <= b.maxC; c++ {
				other := targetGrid[r][c]
				if other != color {
					// color was painted first, then other painted over it
					if !graph[color][other] {
						graph[color][other] = true
					}
				}
			}
		}
	}

	// Detect cycle via DFS (topological sort / Kahn's algorithm)
	// Use three-color DFS: 0=unvisited, 1=visiting, 2=visited
  // Membuat map (HashMap) — pencarian O(1)
	state := make(map[int]int)
	var colors []int
	for c := range hasColor {
		colors = append(colors, c)
	}

	var dfs func(int) bool
	dfs = func(node int) bool {
		state[node] = 1 // visiting
		for next := range graph[node] {
			if state[next] == 1 {
				return true // cycle found
			}
			if state[next] == 0 {
				if dfs(next) {
					return true
				}
			}
		}
		state[node] = 2 // visited
		return false
	}

	for _, c := range colors {
		if state[c] == 0 {
			if dfs(c) {
				return false
			}
		}
	}

	return true
}

func main() {
	// Example 1:
	// Input: targetGrid = [[1,1,1,1],[1,2,2,1],[1,2,2,1],[1,1,1,1]]
	// Output: true
	fmt.Println(isPrintable([][]int{
		{1, 1, 1, 1},
		{1, 2, 2, 1},
		{1, 2, 2, 1},
		{1, 1, 1, 1},
	}))

	// Example 2:
	// Input: targetGrid = [[1,1,1,1],[1,1,3,3],[1,1,3,3],[1,1,1,1]]
	// Output: true
	fmt.Println(isPrintable([][]int{
		{1, 1, 1, 1},
		{1, 1, 3, 3},
		{1, 1, 3, 3},
		{1, 1, 1, 1},
	}))

	// Example 3:
	// Input: targetGrid = [[1,1,1],[3,1,3]]
	// Output: false
	// Color 1's bounding box covers [0,0]..[1,2]. Cell (1,0)=3 and (1,2)=3
	// mean 1->3. Color 3's bounding box covers [1,0]..[1,2]. Cell (1,1)=1
	// means 3->1. Cycle: 1->3->1.
	fmt.Println(isPrintable([][]int{
		{1, 1, 1},
		{3, 1, 3},
	}))
}
```
