# 1722 — Minimize Hamming Distance After Swap Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewUnionFind(n int) *UnionFind
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** O(n + swaps), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1722: Minimize Hamming Distance After Swap Operations
// https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/
// Difficulty: Medium
// Time: O(n + swaps), Space: O(n)

import "fmt"

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
  // Alokasi slice integer
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px != py {
		uf.parent[px] = py
	}
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	uf := NewUnionFind(n)

	for _, swap := range allowedSwaps {
		uf.Union(swap[0], swap[1])
	}

	// Group indices by component
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		groups[root] = append(groups[root], i)
	}

	hamming := 0
	for _, indices := range groups {
  // Membuat map (HashMap) — pencarian O(1)
		counts := make(map[int]int)
		for _, idx := range indices {
			counts[source[idx]]++
		}
		for _, idx := range indices {
			if counts[target[idx]] > 0 {
				counts[target[idx]]--
			} else {
				hamming++
			}
		}
	}
	return hamming
}

func main() {
	fmt.Println(minimumHammingDistance([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}})) // Expected: 1
	fmt.Println(minimumHammingDistance([]int{1, 2, 3, 4}, []int{1, 3, 2, 4}, [][]int{})) // Expected: 2
	fmt.Println(minimumHammingDistance([]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}})) // Expected: 0
}
```
