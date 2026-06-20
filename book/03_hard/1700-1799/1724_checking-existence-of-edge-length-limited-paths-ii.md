# 1724 — Checking Existence Of Edge Length Limited Paths Ii

## Deskripsi

**Soal:** [1724. Checking Existence Of Edge Length Limited Paths Ii](https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist`

## Solusi Go

```go
package main

// LeetCode #1724: Checking Existence of Edge Length Limited Paths II
// https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths-ii/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type DistanceLimitedPathsExist struct {
	n     int
	edges [][]int
}

func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist {
  // Membuat slice 2D untuk DP/tabel
	sortedEdges := make([][]int, len(edgeList))
	copy(sortedEdges, edgeList)
	sort.Slice(sortedEdges, func(i, j int) bool {
		return sortedEdges[i][2] < sortedEdges[j][2]
	})
	return DistanceLimitedPathsExist{n: n, edges: sortedEdges}
}

func (this *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, this.n)
	for i := 0; i < this.n; i++ {
		parent[i] = i
	}
	var find func(int) int
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
		}
	}

	for _, e := range this.edges {
		if e[2] >= limit {
			break
		}
		union(e[0], e[1])
	}
	return find(p) == find(q)
}

func main() {
	// Test 1
	edges1 := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	obj := Constructor(3, edges1)
	fmt.Println("Test 1:")
	fmt.Printf(" Query(0,2,2): %v (Expected: false)\n", obj.Query(0, 2, 2))
	fmt.Printf(" Query(0,2,5): %v (Expected: true)\n\n", obj.Query(0, 2, 5))

	// Test 2
	edges2 := [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 3}, {0, 3, 20}}
	obj2 := Constructor(4, edges2)
	fmt.Println("Test 2:")
	fmt.Printf(" Query(0,3,15): %v (Expected: true)\n", obj2.Query(0, 3, 15))
	fmt.Printf(" Query(0,3,4):  %v (Expected: false)\n", obj2.Query(0, 3, 4))
	fmt.Printf(" Query(0,2,6):  %v (Expected: true)\n", obj2.Query(0, 2, 6))
}
```
