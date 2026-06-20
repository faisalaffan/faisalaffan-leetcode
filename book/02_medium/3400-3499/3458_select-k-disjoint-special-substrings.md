# 3458 — Select K Disjoint Special Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSubstringLength(s string, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n + 26^3) Space: O(26)  
**Kompleksitas Ruang:** O(26)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3458: Select K Disjoint Special Substrings
// https://leetcode.com/problems/select-k-disjoint-special-substrings/
// Difficulty: Medium
// Time: O(n + 26^3) Space: O(26)

import (
	"fmt"
	"slices"
)

func maxSubstringLength(s string, k int) bool {
	if k == 0 {
		return true
	}

	pos := [26][]int{}
	for i, b := range s {
		b -= 'a'
		pos[b] = append(pos[b], i)
	}

	g := [26][]int{}
	for i, p := range pos {
		if p == nil {
			continue
		}
		l, r := p[0], p[len(p)-1]
		for j, q := range pos {
			if j == i || q == nil {
				continue
			}
			idx := lowerBound(q, l)
			if idx < len(q) && q[idx] <= r {
				g[i] = append(g[i], j)
			}
		}
	}

	visited := make([]bool, 26)
	var intervals [][2]int

	for i, p := range pos {
		if p == nil {
			continue
		}
		for j := range visited {
			visited[j] = false
		}
		curL, curR := len(s), 0

		var dfs func(x int)
		dfs = func(x int) {
			visited[x] = true
			pp := pos[x]
			if pp[0] < curL {
				curL = pp[0]
			}
			if pp[len(pp)-1] > curR {
				curR = pp[len(pp)-1]
			}
			for _, y := range g[x] {
				if !visited[y] {
					dfs(y)
				}
			}
		}
		dfs(i)

		if curL > 0 || curR < len(s)-1 {
			intervals = append(intervals, [2]int{curL, curR})
		}
	}

	slices.SortFunc(intervals, func(a, b [2]int) int { return a[1] - b[1] })
	ans := 0
	preR := -1
	for _, p := range intervals {
		if p[0] > preR {
			ans++
			preR = p[1]
		}
	}
	return ans >= k
}

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	fmt.Println(maxSubstringLength("abcdbaefab", 2)) // true
	fmt.Println(maxSubstringLength("cbc", 1)) // true
}
```
