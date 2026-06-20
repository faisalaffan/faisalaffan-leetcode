# 3331 — Find Subtree Sizes After Changes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func findSubtreeSizes(parent []int, s string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

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
  // Matriks 2D
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		g[parent[i]] = append(g[parent[i]], i)
	}

  // Alokasi slice
	ans := make([]int, n)
  // Alokasi slice
	last := make([]int, 26)
  // Range loop
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
