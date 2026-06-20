# 0834 — Sum Of Distances In Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfDistancesInTree(n int, edges [][]int) []int
```

> **💡 Hint:** Rerooting DP. First DFS from root to get subtree sizes and sum of distances

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #834: Sum of Distances in Tree
// https://leetcode.com/problems/sum-of-distances-in-tree/
// Difficulty: Hard
// Approach: Rerooting DP. First DFS from root to get subtree sizes and sum of distances
// from root. Second DFS to compute answers for all nodes using reroot formula:
// ans[child] = ans[parent] + n - 2*subtree[child]

import "fmt"

func sumOfDistancesInTree(n int, edges [][]int) []int {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Alokasi slice integer
	subtree := make([]int, n)
  // Alokasi slice integer
	ans := make([]int, n)

	var dfs1 func(u, parent int)
	dfs1 = func(u, parent int) {
		subtree[u] = 1
		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			dfs1(v, u)
			subtree[u] += subtree[v]
			ans[0] += subtree[v]
		}
	}
	dfs1(0, -1)

	var dfs2 func(u, parent int)
	dfs2 = func(u, parent int) {
		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			ans[v] = ans[u] + n - 2*subtree[v]
			dfs2(v, u)
		}
	}
	dfs2(0, -1)

	return ans
}

func main() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	// Expected: [8 12 6 10 10 10]

	fmt.Println(sumOfDistancesInTree(1, [][]int{}))
	// Expected: [0]
}
```
