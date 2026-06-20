# 2876 — Count Visited Nodes In A Directed Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func countVisitedNodesInADirectedGraph(edges []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2876: Count Visited Nodes in a Directed Graph
// https://leetcode.com/problems/count-visited-nodes-in-a-directed-graph/
// Difficulty: Hard
//
// Each node has exactly one outgoing edge (functional graph). For each node,
// count the number of distinct nodes visited starting from that node.
//
// Approach: Use DFS with three states (0=unvisited, 1=in current path, 2=processed).
// For each unvisited node, follow edges until we encounter a visited node.
// If we encounter a node in the current path, we found a cycle. Assign cycle
// length to all nodes in the cycle. For nodes outside cycles, result = result[child] + 1.

import "fmt"

func countVisitedNodesInADirectedGraph(edges []int) []int {
	n := len(edges)
  // Alokasi slice
	res := make([]int, n)
  // Alokasi slice
	state := make([]int, n) // 0=unvisited, 1=visiting, 2=done

	var dfs func(u int)
	dfs = func(u int) {
		if state[u] == 2 {
			return
		}
		if state[u] == 1 {
			// Found a cycle: compute cycle length
			cycleLen := 1
			v := edges[u]
			for v != u {
				cycleLen++
				v = edges[v]
			}
			// Assign cycle length to all nodes in the cycle
			res[u] = cycleLen
			v = edges[u]
			for v != u {
				res[v] = cycleLen
				v = edges[v]
			}
			return
		}

		state[u] = 1
		dfs(edges[u])
		state[u] = 2

		if res[u] == 0 {
			res[u] = res[edges[u]] + 1
		}
	}

	for i := 0; i < n; i++ {
		if state[i] == 0 {
			dfs(i)
		}
	}

	return res
}

func main() {
	// Example 1: edges=[1,2,0,0] -> [3,3,3,4]
	// 0->1->2->0 (cycle 0,1,2), 3->0 (part of cycle)
	fmt.Println(countVisitedNodesInADirectedGraph([]int{1, 2, 0, 0}))

	// Example 2: edges=[1,2,3,4,0] -> [5,5,5,5,5] (single cycle)
	fmt.Println(countVisitedNodesInADirectedGraph([]int{1, 2, 3, 4, 0}))

	// Line ending in cycle
	// 0->1->2->3->4->3 (cycle at 3,4)
	fmt.Println(countVisitedNodesInADirectedGraph([]int{1, 2, 3, 4, 3}))

	// Self loop
	fmt.Println(countVisitedNodesInADirectedGraph([]int{0, 0, 0}))

	// All point to same node
	// 0->1, 2->1, 1->1 (self loop)
	fmt.Println(countVisitedNodesInADirectedGraph([]int{1, 1, 1}))

	// Two separate components
	fmt.Println(countVisitedNodesInADirectedGraph([]int{1, 0, 3, 2}))
}
```
