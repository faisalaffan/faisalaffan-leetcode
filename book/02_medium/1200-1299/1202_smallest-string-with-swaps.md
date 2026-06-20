# 1202 — Smallest String With Swaps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func newUF(n int) *uf
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1202: Smallest String With Swaps
// https://leetcode.com/problems/smallest-string-with-swaps/
// Difficulty: Medium

// Union-Find on indices. Indices in same connected component can be
// rearranged arbitrarily. Sort chars within each component.

// Time: O(n log n)
// Space: O(n)

type uf struct {
	parent []int
	rank   []int
}

func newUF(n int) *uf {
  // Alokasi slice integer
	p := make([]int, n)
  // Alokasi slice integer
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &uf{p, r}
}

func (u *uf) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *uf) union(x, y int) {
	x, y = u.find(x), u.find(y)
	if x == y {
		return
	}
	if u.rank[x] < u.rank[y] {
		x, y = y, x
	}
	u.parent[y] = x
	if u.rank[x] == u.rank[y] {
		u.rank[x]++
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	u := newUF(n)
	for _, p := range pairs {
		u.union(p[0], p[1])
	}

	// Group indices by root
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := u.find(i)
		groups[root] = append(groups[root], i)
	}

	result := make([]byte, n)
	for _, indices := range groups {
		chars := make([]byte, len(indices))
		for i, idx := range indices {
			chars[i] = s[idx]
		}
  // Custom sort dengan comparator
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

  // Urutkan secara ascending — O(n log n)
		sort.Ints(indices)
		for i, idx := range indices {
			result[idx] = chars[i]
		}
	}

	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n",
		smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}), "bacd")

	fmt.Printf("%q (expected: %q)\n",
		smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}), "abcd")
}
```
