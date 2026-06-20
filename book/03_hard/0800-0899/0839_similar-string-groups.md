# 0839 — Similar String Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numSimilarGroups(strs []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Union-Find

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #839: Similar String Groups
// https://leetcode.com/problems/similar-string-groups/
// Difficulty: Hard
// Approach: Union-Find. Two strings are similar if they differ by exactly 0 or 2 characters
// (i.e., swapping two positions makes them equal). Group connected components.

import "fmt"

func numSimilarGroups(strs []string) int {
	n := len(strs)
  // Alokasi slice
	parent := make([]int, n)
  // Range loop
	for i := range parent {
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

	isSimilar := func(a, b string) bool {
		diff := 0
  // Linear scan O(n)
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
				if diff > 2 {
					return false
				}
			}
		}
		return diff == 0 || diff == 2
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}

  // HashMap: O(1) lookup
	groups := make(map[int]bool)
	for i := 0; i < n; i++ {
		groups[find(i)] = true
	}
	return len(groups)
}

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"})) // Expected: 2
	fmt.Println(numSimilarGroups([]string{"abc", "abc"}))                   // Expected: 1
	fmt.Println(numSimilarGroups([]string{"omv", "ovm"}))                   // Expected: 1
}
```
