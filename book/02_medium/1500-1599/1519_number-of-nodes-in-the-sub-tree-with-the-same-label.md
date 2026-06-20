# 1519 — Number Of Nodes In The Sub Tree With The Same Label

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSubTrees(n int, edges [][]int, labels string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1519: Number of Nodes in the Sub-Tree With the Same Label
// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	fmt.Println(CountSubTrees(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb"))
	fmt.Println(CountSubTrees(5, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, "aabab"))
}

func CountSubTrees(n int, edges [][]int, labels string) []int {
	// Time: O(N), Space: O(N)
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Alokasi slice integer
	result := make([]int, n)
	visited := make([]bool, n)

	var dfs func(node int) []int
	dfs = func(node int) []int {
		visited[node] = true
		// Count array for 26 lowercase letters
  // Alokasi slice integer
		count := make([]int, 26)
		count[labels[node]-'a'] = 1

		for _, nei := range graph[node] {
			if visited[nei] {
				continue
			}
			childCount := dfs(nei)
			for i := 0; i < 26; i++ {
				count[i] += childCount[i]
			}
		}

		result[node] = count[labels[node]-'a']
		return count
	}

	dfs(0)
	return result
}
```
