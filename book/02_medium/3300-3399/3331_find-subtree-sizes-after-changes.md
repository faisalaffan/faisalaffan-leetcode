# 3331 — Find Subtree Sizes After Changes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findSubtreeSizes(parent []int, s string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3331: Find Subtree Sizes After Changes
// https://leetcode.com/problems/find-subtree-sizes-after-changes/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(findSubtreeSizes([]int{-1, 0, 0, 1, 1, 2}, "abacbe")) // [6 3 2 1 1 1]
	fmt.Println(findSubtreeSizes([]int{-1, 0, 0}, "abc"))             // [3 1 1]
}

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		g[parent[i]] = append(g[parent[i]], i)
	}

  // Alokasi slice integer
	ans := make([]int, n)
  // Alokasi slice integer
	last := make([]int, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range last {
		last[i] = -1
	}

	var dfs func(u int)
	dfs = func(u int) {
		old := last[s[u]-'a']
		last[s[u]-'a'] = u
		ans[u] = 1

		for _, v := range g[u] {
			dfs(v)
			p := last[s[v]-'a']
			if p == -1 {
				ans[u] += ans[v]
			} else {
				ans[p] += ans[v]
			}
		}

		last[s[u]-'a'] = old
	}

	dfs(0)
	return ans
}
```
