# 3419 — Minimize The Maximum Edge Weight Of Graph

## Deskripsi

**Soal:** [3419. Minimize The Maximum Edge Weight Of Graph](https://leetcode.com/problems/minimize-the-maximum-edge-weight-of-graph/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O((n+m) log n) Space: O(n+m)  
**Kompleksitas Ruang:** O(n+m)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func minMaxWeight(n int, edges [][]int, _ int) int`

## Solusi Go

```go
package main

// LeetCode #3419: Minimize the Maximum Edge Weight of Graph
// https://leetcode.com/problems/minimize-the-maximum-edge-weight-of-graph/
// Difficulty: Medium
// Time: O((n+m) log n) Space: O(n+m)

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
)

type edge struct{ to, w int }
type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minMaxWeight(n int, edges [][]int, _ int) int {
	if len(edges) < n-1 {
		return -1
	}

  // Membuat slice 2D untuk DP/tabel
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[y] = append(g[y], edge{x, w}) // reverse graph
	}

  // Membuat slice untuk menyimpan hasil
	dis := make([]int, n)
  // Iterasi seluruh elemen
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0
	h := &hp{{}}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		x := p.x
		d := p.dis
		if d > dis[x] {
			continue
		}
		for _, e := range g[x] {
			nd := max(d, e.w)
			if nd < dis[e.to] {
				dis[e.to] = nd
				heap.Push(h, pair{nd, e.to})
			}
		}
	}

	ans := slices.Max(dis)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minMaxWeight(6, [][]int{{0, 1, 4}, {0, 2, 3}, {1, 3, 2}, {1, 4, 1}, {2, 5, 5}, {3, 0, 6}, {4, 0, 7}, {5, 0, 8}}, 3)) // 5
	fmt.Println(minMaxWeight(3, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 0, 4}}, 2)) // 4
}
```
