# 0675 — Cut Off Trees For Golf Event

## Deskripsi

**Soal:** [0675. Cut Off Trees For Golf Event](https://leetcode.com/problems/cut-off-trees-for-golf-event/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** BFS (Breadth-First Search / pencarian lebar), Heap (priority queue)

## Solusi Go

```go
package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

// LeetCode #675: Cut Off Trees for Golf Event
// https://leetcode.com/problems/cut-off-trees-for-golf-event/
// Difficulty: Hard
//
// Sort all tree positions by height, then BFS from start to each tree
// in ascending order. Sum distances. O((m*n)^2) worst case.

func main() {
	// Example: [[1,2,3],[0,0,4],[7,6,5]] => 6
	fmt.Println(cutOffTree([][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}))
	// Blocked: [[1,2,3],[0,0,0],[7,6,5]] => -1
	fmt.Println(cutOffTree([][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}}))
	// [[2,3,4],[0,0,5],[8,7,6]] => 6
	fmt.Println(cutOffTree([][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}}))
	// Single cell
	fmt.Println(cutOffTree([][]int{{1}}))
	// 1x2
	fmt.Println(cutOffTree([][]int{{1, 3}}))
}

type tree struct {
	h, r, c int
}

type pqItem struct {
	r, c, dist int
}

type minHeap []pqItem

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(pqItem)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	var trees []tree
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, tree{h: forest[i][j], r: i, c: j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].h < trees[j].h })

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	sr, sc := 0, 0
	total := 0

	for _, t := range trees {
		dist := bfsShortest(forest, sr, sc, t.r, t.c, m, n, dirs)
		if dist == -1 {
			return -1
		}
		total += dist
		sr, sc = t.r, t.c
	}
	return total
}

func bfsShortest(forest [][]int, sr, sc, tr, tc, m, n int, dirs [][2]int) int {
	if sr == tr && sc == tc {
		return 0
	}
  // Membuat slice 2D untuk DP/tabel
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[sr][sc] = 0
	h := &minHeap{{sr, sc, 0}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).(pqItem)
		if cur.r == tr && cur.c == tc {
			return cur.dist
		}
		if cur.dist > dist[cur.r][cur.c] {
			continue
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || forest[nr][nc] == 0 {
				continue
			}
			nd := cur.dist + 1
			if nd < dist[nr][nc] {
				dist[nr][nc] = nd
				heap.Push(h, pqItem{nr, nc, nd})
			}
		}
	}
	return -1
}
```
