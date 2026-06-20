# 0499 — The Maze Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findShortestWay(maze [][]int, ball []int, hole []int) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, BFS, Heap, Dijkstra

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #499: The Maze III
// https://leetcode.com/problems/the-maze-iii/
// Difficulty: Hard [Paid]
// Approach: Dijkstra with lexicographic path. The ball rolls until it hits a wall.
// We use a priority queue to find the shortest path. If multiple paths have the
// same length, choose the lexicographically smaller path string.

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println("499 - The Maze III")

	// Test cases
	maze1 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	ball1 := []int{4, 3}
	hole1 := []int{0, 1}
	fmt.Printf("findShortestWay(maze1, ball=%v, hole=%v) = %q (expected: \"lul\")\n",
		ball1, hole1, findShortestWay(maze1, ball1, hole1))

	maze2 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	ball2 := []int{4, 3}
	hole2 := []int{3, 0}
	fmt.Printf("findShortestWay(maze2, ball=%v, hole=%v) = %q (expected: \"impossible\")\n",
		ball2, hole2, findShortestWay(maze2, ball2, hole2))

	maze3 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	ball3 := []int{0, 0}
	hole3 := []int{2, 2}
	fmt.Printf("findShortestWay(maze3, ball=%v, hole=%v) = %q\n",
		ball3, hole3, findShortestWay(maze3, ball3, hole3))

	maze4 := [][]int{
		{0, 0},
		{0, 0},
	}
	ball4 := []int{0, 0}
	hole4 := []int{0, 1}
	fmt.Printf("findShortestWay(maze4, ball=%v, hole=%v) = %q (expected: \"r\")\n",
		ball4, hole4, findShortestWay(maze4, ball4, hole4))
}

// Directions: down, up, right, left (to match expected lexicographic order)
// "d", "l", "r", "u" - but we need sorted lexicographically: "d" < "l" < "r" < "u"
// Let's use: d, l, r, u

type Dir struct {
	dr, dc int
	ch     byte
}

var dirs = []Dir{
	{1, 0, 'd'},  // down
	{-1, 0, 'u'}, // up
	{0, 1, 'r'},  // right
	{0, -1, 'l'}, // left
}

// State represents (row, col) with distance and path
type State struct {
	r, c int
	dist int
	path string
	idx   int // for heap
}

// Priority queue
type PQ []*State

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].dist != pq[j].dist {
		return pq[i].dist < pq[j].dist
	}
	return pq[i].path < pq[j].path
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	state := x.(*State)
	state.idx = n
	*pq = append(*pq, state)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	state.idx = -1
	*pq = old[0 : n-1]
	return state
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	m, n := len(maze), len(maze[0])
	startR, startC := ball[0], ball[1]
	holeR, holeC := hole[0], hole[1]

	// dist[r][c] = minimum distance to reach (r,c)
  // Matriks 2D
	dist := make([][]int, m)
  // Range loop
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	// path[r][c] = lexicographically smallest path to reach (r,c)
  // Matriks 2D
	path := make([][]string, m)
  // Range loop
	for i := range path {
		path[i] = make([]string, n)
	}

	pq := &PQ{}
	heap.Init(pq)

	dist[startR][startC] = 0
	path[startR][startC] = ""
  // Push ke priority queue
	heap.Push(pq, &State{r: startR, c: startC, dist: 0, path: ""})

	for pq.Len() > 0 {
  // Pop dari priority queue
		cur := heap.Pop(pq).(*State)

		// Skip if we already found a better path to this cell
		if cur.dist > dist[cur.r][cur.c] {
			continue
		}
		if cur.path != path[cur.r][cur.c] {
			continue
		}

		// Try all 4 directions
		for _, d := range dirs {
			// Roll the ball in this direction until it hits a wall or goes out of bounds
			nr, nc := cur.r, cur.c
			steps := 0

			for {
				// Check if we hit the hole during rolling
				if nr == holeR && nc == holeC {
					break
				}

				nextR := nr + d.dr
				nextC := nc + d.dc

				// Stop if next position is out of bounds or is a wall
				if nextR < 0 || nextR >= m || nextC < 0 || nextC >= n {
					break
				}
				if maze[nextR][nextC] == 1 {
					break
				}

				nr = nextR
				nc = nextC
				steps++
			}

			if steps == 0 {
				continue
			}

			newDist := cur.dist + steps
			newPath := cur.path + string(d.ch)

			// Check if this is a better path to (nr, nc)
			if newDist < dist[nr][nc] || (newDist == dist[nr][nc] && newPath < path[nr][nc]) {
				dist[nr][nc] = newDist
				path[nr][nc] = newPath
  // Push ke priority queue
				heap.Push(pq, &State{r: nr, c: nc, dist: newDist, path: newPath})
			}
		}
	}

	if dist[holeR][holeC] == math.MaxInt32 {
		return "impossible"
	}
	return path[holeR][holeC]
}
```
