# 1202 — Smallest String With Swaps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func newUF(n int) *uf`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting, Union-Find

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // Alokasi slice
	p := make([]int, n)
  // Alokasi slice
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
  // HashMap: O(1) lookup
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
  // Custom sort
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

  // Sort O(n log n)
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
