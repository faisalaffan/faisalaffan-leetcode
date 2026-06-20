# 3067 — Count Pairs Of Connectable Servers In A Weighted Tree Network

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3067: Count Pairs of Connectable Servers in a Weighted Tree Network
// https://leetcode.com/problems/count-pairs-of-connectable-servers-in-a-weighted-tree-network/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 1, 1}, {1, 2, 5}}, 1))
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}}, 3))
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
  // Membuat matriks/slice 2D untuk DP
	g := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, w})
	}

  // Alokasi slice integer
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		cnt := []int{}
		for _, ne := range g[i] {
			c := dfs(i, ne[0], ne[1], signalSpeed, g)
			if c > 0 {
				cnt = append(cnt, c)
			}
		}
		total := 0
		for j := 0; j < len(cnt); j++ {
			total += cnt[j]
		}
		pairs := 0
		for j := 0; j < len(cnt); j++ {
			total -= cnt[j]
			pairs += cnt[j] * total
		}
		ans[i] = pairs
	}
	return ans
}

func dfs(prev, curr, dist, signalSpeed int, g [][][2]int) int {
	c := 0
	if dist%signalSpeed == 0 {
		c++
	}
	for _, ne := range g[curr] {
		if ne[0] != prev {
			c += dfs(curr, ne[0], dist+ne[1], signalSpeed, g)
		}
	}
	return c
}
```
