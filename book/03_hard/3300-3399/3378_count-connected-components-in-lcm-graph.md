# 3378 — Count Connected Components In Lcm Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func NewDSU(n int) *DSU
```

> **💡 Hint:** Union-Find with multiples. For each num <= threshold, connect

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3378: Count Connected Components in LCM Graph
// https://leetcode.com/problems/count-connected-components-in-lcm-graph/
// Difficulty: Hard
//
// Given an array nums and a threshold, construct a graph where nodes i and j
// are connected iff lcm(nums[i], nums[j]) <= threshold. Count connected
// components.
//
// Approach: Union-Find with multiples. For each num <= threshold, connect
// it to all multiples up to threshold. Numbers > threshold are isolated.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countComponents([]int{6, 12, 10}, 20))
	// Example 2
	fmt.Println(countComponents([]int{2, 4, 8, 16}, 10))
	// Example 3: all > threshold
	fmt.Println(countComponents([]int{100, 200}, 50))
	// Edge: single element
	fmt.Println(countComponents([]int{5}, 10))
	// Edge: empty
	fmt.Println(countComponents([]int{}, 10))
}

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
  // Alokasi slice integer
	p := make([]int, n)
  // Alokasi slice integer
	r := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	x, y = d.Find(x), d.Find(y)
	if x == y {
		return
	}
	if d.rank[x] < d.rank[y] {
		d.parent[x] = y
	} else if d.rank[x] > d.rank[y] {
		d.parent[y] = x
	} else {
		d.parent[y] = x
		d.rank[x]++
	}
}

func countComponents(nums []int, threshold int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}

	dsu := NewDSU(threshold + 1)
	seen := make([]bool, threshold+1)
	isolated := 0

	for _, num := range nums {
		if num > threshold {
			isolated++
			continue
		}
		// Connect num to all its multiples up to threshold
		for multiple := num; multiple <= threshold; multiple += num {
			seen[multiple] = true
			dsu.Union(num, multiple)
		}
	}

	// Count unique roots among numbers <= threshold
  // Membuat map (HashMap) — pencarian O(1)
	roots := make(map[int]bool)
	for _, num := range nums {
		if num <= threshold {
			roots[dsu.Find(num)] = true
		}
	}

	return isolated + len(roots)
}
```
