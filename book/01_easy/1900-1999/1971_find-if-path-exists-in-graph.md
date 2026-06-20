# 1971 — Find If Path Exists In Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindIfPathExistsInGraph(n int, edges [][]int, source int, destination int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(V + E), Space: O(V + E)  
**Kompleksitas Ruang:** O(V + E)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1971: Find if Path Exists in Graph
// https://leetcode.com/problems/find-if-path-exists-in-graph/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindIfPathExistsInGraph(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2)) // true
	fmt.Println(FindIfPathExistsInGraph(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5)) // false
}

// Time: O(V + E), Space: O(V + E)
func FindIfPathExistsInGraph(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n)
	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if v == destination {
				return true
			}
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return false
}
```
