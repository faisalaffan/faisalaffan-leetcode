# 1782 — Count Pairs Of Nodes

## Deskripsi

**Soal:** [1782. Count Pairs Of Nodes](https://leetcode.com/problems/count-pairs-of-nodes/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func countPairs(n int, edges [][]int, queries []int) []int`

> **Ide Kunci:** Degree sort + Binary Search.

## Solusi Go

```go
package main

// LeetCode #1782: Count Pairs Of Nodes
// https://leetcode.com/problems/count-pairs-of-nodes/
// Difficulty: Hard
//
// Approach: Degree sort + Binary Search.
// 1. Count degree for each node.
// 2. Count edge-incident pairs for each edge (for subtracting overcount).
// 3. For each query, first find all pairs (i,j) with deg[i] + deg[j] > query
//    using two-pointer on sorted degrees.
// 4. Subtract pairs where deg[i] + deg[j] > query but (i,j) is an edge where
//    deg[i] + deg[j] - edgeCount((i,j)) <= query.
// 5. Also handle duplicate edge counts.

import (
	"fmt"
	"sort"
)

func countPairs(n int, edges [][]int, queries []int) []int {
	// Degree of each node
  // Membuat slice untuk menyimpan hasil
	deg := make([]int, n+1)
	// Edge pair counts: key = (min*100000 + max) for uniqueness
  // Membuat map untuk pencarian O(1): key → value
	edgeCount := make(map[int]int)
	for _, e := range edges {
		u, v := e[0], e[1]
		if u > v {
			u, v = v, u
		}
		deg[u]++
		deg[v]++
		edgeCount[u*100000+v]++
	}

	// Sorted degrees (1-indexed)
  // Membuat slice untuk menyimpan hasil
	sortedDeg := make([]int, n)
	copy(sortedDeg, deg[1:])
	sort.Ints(sortedDeg)

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for qi, q := range queries {
		// Two-pointer: count pairs with deg[i] + deg[j] > q
		total := 0
		left, right := 0, n-1
  // Loop two-pointer: kiri vs kanan
		for left < right {
			if sortedDeg[left]+sortedDeg[right] > q {
				total += right - left
				right--
			} else {
				left++
			}
		}

		// Subtract edge pairs that don't satisfy when considering shared edges
  // Membuat map untuk pencarian O(1): key → value
		seen := make(map[int]bool)
		for _, e := range edges {
			u, v := e[0], e[1]
			if u > v {
				u, v = v, u
			}
			key := u*100000 + v
			if seen[key] {
				continue
			}
			seen[key] = true
			cnt := edgeCount[key]
			// deg[u] + deg[v] > q is needed for this edge to be counted,
			// but deg[u] + deg[v] - cnt <= q means it should be subtracted
			if deg[u]+deg[v] > q && deg[u]+deg[v]-cnt <= q {
				total--
			}
		}

		ans[qi] = total
	}

	return ans
}

func main() {
	// Example test case (LeetCode Example 1)
	n := 4
	edges := [][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1}}
	queries := []int{2, 3}
	fmt.Println("n=4,edges=[[1,2],[2,4],[1,3],[2,3],[2,1]],queries=[2,3] →", countPairs(n, edges, queries))
	// Expected: [6, 5]

	// Additional tests
	n2 := 5
	edges2 := [][]int{{1, 5}, {2, 5}, {3, 5}, {4, 5}}
	queries2 := []int{1, 2, 3}
	fmt.Println("n=5,edges=[[1,5],[2,5],[3,5],[4,5]],queries=[1,2,3] →", countPairs(n2, edges2, queries2))

	n3 := 2
	edges3 := [][]int{{1, 2}}
	queries3 := []int{0, 1, 2}
	fmt.Println("n=2,edges=[[1,2]],queries=[0,1,2] →", countPairs(n3, edges3, queries3))
}
```
