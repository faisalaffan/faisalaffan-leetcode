# 1557 — Minimum Number Of Vertices To Reach All Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindSmallestSetOfVertices(n int, edges [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N + E), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1557: Minimum Number of Vertices to Reach All Nodes
// https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
	fmt.Println(FindSmallestSetOfVertices(3, [][]int{{0, 1}, {2, 1}}))
	fmt.Println(FindSmallestSetOfVertices(5, [][]int{{0, 1}, {2, 1}, {3, 1}, {4, 0}}))
}

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	// Time: O(N + E), Space: O(N)
	// Nodes with indegree 0 must be in the result since they can't be reached
  // Alokasi slice integer
	indegree := make([]int, n)
	for _, e := range edges {
		indegree[e[1]]++
	}

  // Alokasi slice integer
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}
```
