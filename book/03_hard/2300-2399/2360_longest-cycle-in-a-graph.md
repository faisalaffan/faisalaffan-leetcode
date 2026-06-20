# 2360 — Longest Cycle In A Graph

## Deskripsi

**Soal:** [2360. Longest Cycle In A Graph](https://leetcode.com/problems/longest-cycle-in-a-graph/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func longestCycle(edges []int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// 2360. Longest Cycle in a Graph
// ----------------------------------------------------------------
// Directed graph on n nodes.  Each node has exactly one outgoing edge
// (given as array edges[n] where edges[i] = j means i → j, or -1 for no edge).
// Find the length of the longest cycle.  If no cycle exists, return -1.
//
// Use 3‑state DFS: 0 = unvisited, 1 = visiting (on current path),
// 2 = fully processed.  Maintain a distance array for the current DFS walk.
// When we encounter a visiting node, we have found a cycle:
//   length = dist[cur] - dist[node] + 1.

func longestCycle(edges []int) int {
	n := len(edges)
  // Membuat slice untuk menyimpan hasil
	state := make([]int, n) // 0=unvisited, 1=visiting, 2=done
  // Membuat slice untuk menyimpan hasil
	dist := make([]int, n)
	ans := -1

	for start := 0; start < n; start++ {
		if state[start] != 0 {
			continue
		}
		// Walk the current path.
		cur := start
		step := 0
		for cur != -1 && state[cur] == 0 {
			state[cur] = 1
			dist[cur] = step
			step++
			cur = edges[cur]
		}
		// If we hit a node in visiting state, we found a cycle.
		if cur != -1 && state[cur] == 1 {
			cycleLen := step - dist[cur]
			if cycleLen > ans {
				ans = cycleLen
			}
		}
		// Mark all nodes in the current walk as done.
		cur = start
		for cur != -1 && state[cur] == 1 {
			state[cur] = 2
			cur = edges[cur]
		}
	}
	return ans
}

// ---------------------------------------------------------------------------
//  Wrapper

func LongestCycleInAGraph() interface{} {
	return longestCycle([]int{3, 3, 4, 2, 3})
}

func main() {
	fmt.Println(LongestCycleInAGraph())

	tests := []struct {
		edges []int
		want  int
	}{
		{[]int{3, 3, 4, 2, 3}, 3},
		{[]int{2, -1, 3, 1}, -1},
		{[]int{1, 2, 0, 4, 5, 3}, 3}, // 0→1→2→0 (len 3) and 3→4→5→3 (len 3) → max 3
		{[]int{-1, -1, -1}, -1},
		{[]int{1, 2, 3, 4, 0}, 5}, // 0→1→2→3→4→0
	}
	// Recompute test 3 manually: edges[3]=4, edges[4]=5, edges[5]=3 → cycle of 3. edges[0]=1, edges[1]=2, edges[2]=0 → cycle of 3. Longest=3.
	// So the test says want 3, that's fine.

	for _, tc := range tests {
		got := longestCycle(tc.edges)
		if got != tc.want {
			fmt.Printf("FAIL edges=%v: got %d, want %d\n", tc.edges, got, tc.want)
		}
	}
	fmt.Println("Done testing 2360.")
}
```
