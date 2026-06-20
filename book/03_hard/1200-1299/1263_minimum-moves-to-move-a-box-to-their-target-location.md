# 1263 — Minimum Moves To Move A Box To Their Target Location

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minPushBox(grid [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1263: Minimum Moves to Move a Box to Their Target Location
// https://leetcode.com/problems/minimum-moves-to-move-a-box-to-their-target-location/
// Difficulty: Hard
//
// In a grid of cells (0=empty, 1=obstacle), there is a player, a box, and a
// target cell. The player can push the box (moving it one cell) if standing
// behind it. Find the minimum number of pushes (box moves) required to move
// the box to the target. Return -1 if impossible.

import (
	"fmt"
)

func main() {
	// Example 1
	grid := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '.', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'},
	}
	// 'S' = player start, 'B' = box, 'T' = target, '.' = empty, '#' = obstacle
	fmt.Println(minPushBox(grid)) // 3

	// Example 2 (immediate push)
	grid2 := [][]byte{
		{'#', '#', '#'},
		{'#', 'B', '#'},
		{'#', 'S', 'T'},
		{'#', '#', '#'},
	}
	fmt.Println(minPushBox(grid2)) // 0? Actually need to check grid indexing. Box at (1,1), player can push it down to target at (2,2)?

	// Wait, let me reconsider the example. Player at S, target at T, box at B.
	// Player must push box toward target.
	_ = grid2

	// Simple: box already at target
	grid3 := [][]byte{
		{'#', '#', '#'},
		{'#', 'T', '#'},
		{'#', 'S', '#'},
		{'#', '#', '#'},
	}
	// No box? Let's just handle properly based on input.
	_ = grid3
}

func minPushBox(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	// Locate player, box, target
	var playerR, playerC, boxR, boxC, targetR, targetC int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case 'S':
				playerR, playerC = r, c
			case 'B':
				boxR, boxC = r, c
			case 'T':
				targetR, targetC = r, c
			}
		}
	}

	// BFS over box position. We track (boxR, boxC, playerR, playerC) state
	// but optimize: for each box position, we only care if the player can
	// reach the pushing position. We'll use a 0-1 BFS (deque) where pushes
	// cost 1 and player movement costs 0.

	// dist[boxR][boxC] = minimum pushes to get box here
	const INF = 1 << 30
  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, rows)
	for r := 0; r < rows; r++ {
		dist[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			dist[r][c] = INF
		}
	}

	// BFS over (boxR, boxC) with player position implicitly tracked.
	// We use a queue and a visited set for (boxR, boxC, playerR, playerC).
	type State struct {
		br, bc, pr, pc int
	}

	queue := make([]State, 0, rows*cols*4)
  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[State]bool)

	start := State{boxR, boxC, playerR, playerC}
	visited[start] = true
	queue = append(queue, start)
	dist[boxR][boxC] = 0

	// BFS
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	head := 0

	for head < len(queue) {
		cur := queue[head]
		head++

		// Check if box reached target
		if cur.br == targetR && cur.bc == targetC {
			return dist[cur.br][cur.bc]
		}

		// Try each push direction
		for _, d := range dirs {
			nbr := cur.br + d[0]
			nbc := cur.bc + d[1]

			// Box destination must be empty
			if nbr < 0 || nbr >= rows || nbc < 0 || nbc >= cols || grid[nbr][nbc] == '#' {
				continue
			}

			// Player must be able to reach the pushing position (behind the box)
			// pushing position = cur.br - d[0], cur.bc - d[1]
			pushR := cur.br - d[0]
			pushC := cur.bc - d[1]

			if pushR < 0 || pushR >= rows || pushC < 0 || pushC >= cols || grid[pushR][pushC] == '#' {
				continue
			}

			// BFS/DFS from player position to pushing position (no box moves)
			if !canReach(grid, cur.pr, cur.pc, pushR, pushC, cur.br, cur.bc) {
				continue
			}

			ns := State{nbr, nbc, cur.br, cur.bc}
			if !visited[ns] {
				visited[ns] = true
				dist[nbr][nbc] = dist[cur.br][cur.bc] + 1
				queue = append(queue, ns)
			}
		}
	}

	return -1
}

// canReach checks if the player can walk from (sr, sc) to (tr, tc) without
// stepping through the box position (boxR, boxC) or obstacles (#).
func canReach(grid [][]byte, sr, sc, tr, tc, boxR, boxC int) bool {
	if sr == tr && sc == tc {
		return true
	}

	rows := len(grid)
	cols := len(grid[0])
  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		visited[r] = make([]bool, cols)
	}

  // Alokasi slice integer
	queue := make([][2]int, 0, rows*cols)
	queue = append(queue, [2]int{sr, sc})
	visited[sr][sc] = true

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur[0] == tr && cur[1] == tc {
			return true
		}

		for _, d := range dirs {
			nr := cur[0] + d[0]
			nc := cur[1] + d[1]

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if grid[nr][nc] == '#' {
				continue
			}
			if nr == boxR && nc == boxC {
				continue // cannot walk through box
			}
			if visited[nr][nc] {
				continue
			}

			visited[nr][nc] = true
			queue = append(queue, [2]int{nr, nc})
		}
	}

	return false
}
```
