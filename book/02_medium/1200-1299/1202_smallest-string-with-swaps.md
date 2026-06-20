# 1202 — Smallest String With Swaps

## Deskripsi

**Soal:** [1202. Smallest String With Swaps](https://leetcode.com/problems/smallest-string-with-swaps/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func newUF(n int) *uf`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	p := make([]int, n)
  // Membuat slice untuk menyimpan hasil
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
  // Membuat map untuk pencarian O(1): key → value
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := u.find(i)
		groups[root] = append(groups[root], i)
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]byte, n)
	for _, indices := range groups {
  // Membuat slice untuk menyimpan hasil
		chars := make([]byte, len(indices))
		for i, idx := range indices {
			chars[i] = s[idx]
		}
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

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
