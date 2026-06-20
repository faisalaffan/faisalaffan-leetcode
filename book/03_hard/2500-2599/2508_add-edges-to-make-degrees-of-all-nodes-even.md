# 2508 — Add Edges To Make Degrees Of All Nodes Even

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func isPossible(n int, edges [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2508: Add Edges to Make Degrees of All Nodes Even
// https://leetcode.com/problems/add-edges-to-make-degrees-of-all-nodes-even/
// Difficulty: Hard
//
// Find nodes with odd degree. We can add at most 2 edges.
// Cases:
//  0 odd nodes => true
//  2 odd nodes => connect them directly, or connect both to a third node (with even degree)
//  4 odd nodes => try all 3 pairings
//  > 4 odd nodes => false (2 edges fix at most 4 odd nodes)

import "fmt"

func main() {
	// Example 1: n=5, edges=[[1,2],[2,3],[3,4],[4,2],[1,4],[2,5]] => true
	fmt.Println(isPossible(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 4}, {2, 5}}))
	// Example 2: n=4, edges=[[1,2],[3,4]] => true
	fmt.Println(isPossible(4, [][]int{{1, 2}, {3, 4}}))
	// Example 3: n=4, edges=[[1,2],[1,3],[1,4]] => false
	fmt.Println(isPossible(4, [][]int{{1, 2}, {1, 3}, {1, 4}}))
	// Edge: already all even
	fmt.Println(isPossible(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Edge: 2 odd nodes but already connected directly
	fmt.Println(isPossible(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}, {1, 3}}))
}

func isPossible(n int, edges [][]int) bool {
	// Track which edges exist (dense graph, use adjacency matrix/set)
  // Membuat map (HashMap) — pencarian O(1)
	adjSet := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		adjSet[i] = make(map[int]bool)
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		adjSet[a][b] = true
		adjSet[b][a] = true
	}

  // Alokasi slice integer
	deg := make([]int, n+1)
	for _, e := range edges {
		deg[e[0]]++
		deg[e[1]]++
	}

	odd := []int{}
	for i := 1; i <= n; i++ {
		if deg[i]%2 == 1 {
			odd = append(odd, i)
		}
	}

	if len(odd) == 0 {
		return true
	}

	if len(odd) == 2 {
		a, b := odd[0], odd[1]
		// Can connect a and b directly
		if !adjSet[a][b] {
			return true
		}
		// Connect both to a third node
		for c := 1; c <= n; c++ {
			if c != a && c != b {
				if !adjSet[a][c] && !adjSet[b][c] {
					return true
				}
			}
		}
		return false
	}

	if len(odd) == 4 {
		// Try all 3 pairings: (0-1,2-3), (0-2,1-3), (0-3,1-2)
		pairings := [][2][2]int{
			{{odd[0], odd[1]}, {odd[2], odd[3]}},
			{{odd[0], odd[2]}, {odd[1], odd[3]}},
			{{odd[0], odd[3]}, {odd[1], odd[2]}},
		}
		for _, p := range pairings {
			a1, b1 := p[0][0], p[0][1]
			a2, b2 := p[1][0], p[1][1]
			if !adjSet[a1][b1] && !adjSet[a2][b2] {
				return true
			}
		}
		return false
	}

	return false
}
```
