# 1377 — Frog Position After T Seconds

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func frogPosition(n int, edges [][]int, t int, target int) float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #1377: Frog Position After T Seconds
// https://leetcode.com/problems/frog-position-after-t-seconds/
// Difficulty: Hard
//
// Approach: DFS on tree with probability propagation.
// The frog starts at node 1 at time 0 with probability 1.0.
// At each second, if the current node has unvisited neighbors, the frog
// picks one uniformly and jumps. If no unvisited neighbors remain, the
// frog stays at the current node until time runs out.
// Return the probability that the frog is at `target` at time `t`.

import "fmt"

func frogPosition(n int, edges [][]int, t int, target int) float64 {
  // Matriks 2D
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n+1)

	var dfs func(node, time int, prob float64) float64
	dfs = func(node, time int, prob float64) float64 {
		if time == t {
			if node == target {
				return prob
			}
			return 0
		}
		visited[node] = true

		var children []int
		for _, next := range adj[node] {
			if !visited[next] {
				children = append(children, next)
			}
		}

		var res float64
		if len(children) > 0 {
			childProb := prob / float64(len(children))
			for _, child := range children {
				res += dfs(child, time+1, childProb)
			}
		} else if node == target {
			// Stuck here for remaining time
			res = prob
		}
		visited[node] = false
		return res
	}

	return dfs(1, 0, 1.0)
}

func main() {
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 5}, {3, 6}}, 2, 4)) // 0.166666...
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 5}, {3, 6}}, 1, 7)) // 0.333333...
	fmt.Println(frogPosition(3, [][]int{{1, 2}, {2, 3}}, 1, 2))                                  // 1.0
	fmt.Println(frogPosition(3, [][]int{{2, 1}, {3, 2}}, 1, 3))                                  // 0.0
}
```
