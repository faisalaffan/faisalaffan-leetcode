# 1810 — Minimum Path Cost In A Hidden Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func findShortestPath(master GridMaster) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS, Dijkstra

**Kompleksitas Waktu:** O(m * n log(m*n)), Space: O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1810: Minimum Path Cost in a Hidden Grid
// https://leetcode.com/problems/minimum-path-cost-in-a-hidden-grid/
// Difficulty: Medium [Paid]
// This is an interactive problem. We implement the solving algorithm.
// Time: O(m * n log(m*n)), Space: O(m * n)

import "fmt"

// GridMaster is the interface provided by LeetCode for the hidden grid problem.
// Real implementation: GridMaster.canMove(dir), GridMaster.move(dir), GridMaster.isTarget()
// GridMaster.move(dir) returns int cost to move
type GridMaster interface {
	CanMove(dir byte) bool
	Move(dir byte) int
	IsTarget() bool
}

func findShortestPath(master GridMaster) int {
	// Discover grid: visited[r][c] = true means the cell was reached
  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[[2]int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	costs := make(map[[2]int]int)
	targetPos := [2]int{-1, -1}

	dirs := []byte{'U', 'D', 'L', 'R'}
	dr := map[byte]int{'U': -1, 'D': 1, 'L': 0, 'R': 0}
	dc := map[byte]int{'U': 0, 'D': 0, 'L': -1, 'R': 1}
	rev := map[byte]byte{'U': 'D', 'D': 'U', 'L': 'R', 'R': 'L'}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if master.IsTarget() {
			targetPos = [2]int{r, c}
		}
		visited[[2]int{r, c}] = true
		for _, d := range dirs {
			nr, nc := r+dr[d], c+dc[d]
			key := [2]int{nr, nc}
			if visited[key] {
				continue
			}
			if master.CanMove(d) {
				cost := master.Move(d)
				costs[key] = cost
				dfs(nr, nc)
				master.Move(rev[d])
			}
		}
	}

	dfs(0, 0)

	if targetPos == [2]int{-1, -1} {
		return -1
	}

	// Dijkstra for shortest weighted path
  // Membuat map (HashMap) — pencarian O(1)
	dist := make(map[[2]int]int)
	dist[[2]int{0, 0}] = 0
	pq := [][3]int{{0, 0, 0}} // dist, r, c

	for len(pq) > 0 {
		// Extract min
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i][0] < pq[minIdx][0] {
				minIdx = i
			}
		}
		cur := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)
		d, r, c := cur[0], cur[1], cur[2]

		if [2]int{r, c} == targetPos {
			return d
		}

		if d > dist[[2]int{r, c}] {
			continue
		}

		for _, dir := range dirs {
			nr, nc := r+dr[dir], c+dc[dir]
			nk := [2]int{nr, nc}
			if !visited[nk] {
				continue
			}
			nd := d + costs[nk]
			if old, ok := dist[nk]; !ok || nd < old {
				dist[nk] = nd
				pq = append(pq, [3]int{nd, nr, nc})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Interactive problem - test via LeetCode platform")
	fmt.Println("Implementation ready for GridMaster interface with costs")
}
```
