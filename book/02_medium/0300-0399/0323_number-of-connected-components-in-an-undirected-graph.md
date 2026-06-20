# 0323 — Number Of Connected Components In An Undirected Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func countComponents(n int, edges [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #323: Number of Connected Components in an Undirected Graph
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
// Difficulty: Medium [Paid]
// Time: O(V + E) | Space: O(V)

import "fmt"

func countComponents(n int, edges [][]int) int {
  // Alokasi slice integer
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
			n--
		}
	}

	for _, e := range edges {
		union(e[0], e[1])
	}
	return n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
	// Expected: 1

	// Test case 3: No edges
	fmt.Println("Test 3:", countComponents(3, [][]int{}))
	// Expected: 3
}
```
