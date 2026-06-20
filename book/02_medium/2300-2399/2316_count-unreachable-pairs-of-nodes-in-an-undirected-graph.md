# 2316 — Count Unreachable Pairs Of Nodes In An Undirected Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPairs(n int, edges [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2316: Count Unreachable Pairs of Nodes in an Undirected Graph
// https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func countPairs(n int, edges [][]int) int64 {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	var totalPairs int64 = 0
	var prevComponents int64 = 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		// BFS/DFS to find component size
		size := int64(0)
		queue := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			size++
			for _, nei := range graph[node] {
				if !visited[nei] {
					visited[nei] = true
					queue = append(queue, nei)
				}
			}
		}
		totalPairs += prevComponents * size
		prevComponents += size
	}
	return totalPairs
}

func main() {
	// Test case 1
	fmt.Println(countPairs(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	// Expected: 0

	// Test case 2
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
	// Expected: 14
}
```
