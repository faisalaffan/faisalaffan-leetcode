# 0864 — Shortest Path To Get All Keys

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestPathAllKeys(grid []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #864: Shortest Path to Get All Keys
// https://leetcode.com/problems/shortest-path-to-get-all-keys/
// Difficulty: Hard
//
// BFS with state (row, col, keyMask). Since there are at most 6 keys,
// the key mask fits in 6 bits. Each state is visited at most once.

import "fmt"

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])

	var startR, startC int
	totalKeys := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ch := grid[i][j]
			if ch == '@' {
				startR, startC = i, j
			} else if ch >= 'a' && ch <= 'f' {
				totalKeys++
			}
		}
	}

	allKeys := (1 << totalKeys) - 1

	// visited[r][c][keyMask]
  // Membuat matriks/slice 2D untuk DP
	visited := make([][][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = make([]bool, 1<<totalKeys)
		}
	}

	type state struct{ r, c, keys int }
	queue := []state{{startR, startC, 0}}
	visited[startR][startC][0] = true
	steps := 0
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		for sz := len(queue); sz > 0; sz-- {
			cur := queue[0]
			queue = queue[1:]
			if cur.keys == allKeys {
				return steps
			}
			for _, d := range dirs {
				nr, nc := cur.r+d[0], cur.c+d[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}
				cell := grid[nr][nc]
				if cell == '#' {
					continue
				}
				nk := cur.keys
				if cell >= 'A' && cell <= 'F' {
					keyBit := 1 << (cell - 'A')
					if cur.keys&keyBit == 0 {
						continue // missing key for this lock
					}
				} else if cell >= 'a' && cell <= 'f' {
					nk = cur.keys | (1 << (cell - 'a'))
				}
				if !visited[nr][nc][nk] {
					visited[nr][nc][nk] = true
					queue = append(queue, state{nr, nc, nk})
				}
			}
		}
		steps++
	}
	return -1
}

func main() {
	// Example 1: ["@.a..","###.#","b.A.B"] -> 8
	grid1 := []string{"@.a..", "###.#", "b.A.B"}
	fmt.Println("Test 1:", shortestPathAllKeys(grid1)) // 8

	// Example 2: ["@..aA","..B#.","....b"] -> 6
	grid2 := []string{"@..aA", "..B#.", "....b"}
	fmt.Println("Test 2:", shortestPathAllKeys(grid2)) // 6

	// Example 3: ["@Aa"] -> -1 (can't reach 'a' behind 'A' without key)
	grid3 := []string{"@Aa"}
	fmt.Println("Test 3:", shortestPathAllKeys(grid3)) // -1

	// Single key: ["@a"] -> 1
	grid4 := []string{"@a"}
	fmt.Println("Test 4:", shortestPathAllKeys(grid4)) // 1

	// No keys: ["@."] -> 0
	grid5 := []string{"@."}
	fmt.Println("Test 5:", shortestPathAllKeys(grid5)) // 0
}
```
