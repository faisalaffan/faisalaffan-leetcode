# 3383 — Minimum Runes To Add To Cast Spell

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumRunesToAddToCastSpell(n int, edges [][]int, crystals []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Topological Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
  // Membuat matriks/slice 2D untuk DP
	rev := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		rev[v] = append(rev[v], u)
	}

	reachable := make([]bool, n)
  // Alokasi slice integer
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
