# 2368 — Reachable Nodes With Restrictions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func reachableNodes(n int, edges [][]int, restricted []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2368: Reachable Nodes With Restrictions
// https://leetcode.com/problems/reachable-nodes-with-restrictions/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// BFS/DFS from node 0, skip restricted nodes.

import "fmt"

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5})) // 4
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}, {1, 5}, {2, 6}}, []int{1})) // 3
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	restrictedSet := make(map[int]bool, len(restricted))
	for _, r := range restricted {
		restrictedSet[r] = true
	}

  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		visited[u] = true
		count := 1
		for _, v := range graph[u] {
			if !visited[v] && !restrictedSet[v] {
				count += dfs(v)
			}
		}
		return count
	}
	return dfs(0)
}
```
