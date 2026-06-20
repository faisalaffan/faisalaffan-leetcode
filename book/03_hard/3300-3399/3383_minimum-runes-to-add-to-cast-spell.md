# 3383 — Minimum Runes To Add To Cast Spell

## Deskripsi

**Soal:** [3383. Minimum Runes To Add To Cast Spell](https://leetcode.com/problems/minimum-runes-to-add-to-cast-spell/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), BFS (Breadth-First Search / pencarian lebar)

## Solusi Go

```go
package main

// LeetCode #3383: Minimum Runes to Add to Cast Spell
// https://leetcode.com/problems/minimum-runes-to-add-to-cast-spell/
// Difficulty: Hard [Paid]
//
// BFS from crystal nodes + DFS for topological order.
// Count sink components that are not reachable from crystals.

import "fmt"

func main() {
	fmt.Println(MinimumRunesToAddToCastSpell(6, [][]int{{0, 1}, {0, 2}, {3, 4}}, []int{0, 3}))
}

func MinimumRunesToAddToCastSpell(n int, edges [][]int, crystals []int) int {
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
  // Membuat slice 2D untuk DP/tabel
	rev := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		rev[v] = append(rev[v], u)
	}

  // Membuat slice untuk menyimpan hasil
	reachable := make([]bool, n)
  // Membuat slice untuk menyimpan hasil
	q := make([]int, 0, n)
	for _, c := range crystals {
		reachable[c] = true
		q = append(q, c)
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			if !reachable[v] {
				reachable[v] = true
				q = append(q, v)
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	vis := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		vis[u] = true
		for _, v := range rev[u] {
			if !vis[v] {
				dfs(v)
			}
		}
	}
	for _, c := range crystals {
		if !vis[c] {
			dfs(c)
		}
	}

  // Membuat slice untuk menyimpan hasil
	need := make([]bool, n)
	for u := 0; u < n; u++ {
		if !reachable[u] && len(adj[u]) == 0 {
			need[u] = true
		}
	}

	cnt := 0
	for u := 0; u < n; u++ {
		if need[u] {
			cnt++
		}
	}
	return cnt
}
```
