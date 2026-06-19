package main

// LeetCode #2503: Maximum Number of Points From Grid Queries
// https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/
// Difficulty: Hard
//
// Sort queries ascending while preserving original indices. Use a min-heap
// starting from (0,0). For each query q, pop all cells with value < q from
// the heap and BFS to their unvisited neighbors. Count visited cells.

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// Example 1: grid=[[1,2,3],[2,5,7],[3,5,1]], queries=[5,6,2] => [5,8,1]
	fmt.Println(maxPoints([][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, []int{5, 6, 2}))
	// Example 2: grid=[[5,2,1],[1,1,2]], queries=[3] => [0]
	fmt.Println(maxPoints([][]int{{5, 2, 1}, {1, 1, 2}}, []int{3}))
	// Edge: single cell
	fmt.Println(maxPoints([][]int{{1}}, []int{1, 2}))
	// Edge: all cells reachable
	fmt.Println(maxPoints([][]int{{1, 2}, {3, 4}}, []int{5}))
}

type Cell struct {
	val, r, c int
}

type MinHeap []Cell

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	k := len(queries)

	// Sort queries with original indices
	type Query struct {
		val, idx int
	}
	sorted := make([]Query, k)
	for i, v := range queries {
		sorted[i] = Query{v, i}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].val < sorted[j].val
	})

	result := make([]int, k)
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Cell{grid[0][0], 0, 0})

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	count := 0

	for _, q := range sorted {
		// Pop all cells with value < query value
		for h.Len() > 0 && (*h)[0].val < q.val {
			cell := heap.Pop(h).(Cell)
			count++
			for _, d := range dirs {
				nr, nc := cell.r+d[0], cell.c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] {
					visited[nr][nc] = true
					heap.Push(h, Cell{grid[nr][nc], nr, nc})
				}
			}
		}
		result[q.idx] = count
	}

	return result
}
