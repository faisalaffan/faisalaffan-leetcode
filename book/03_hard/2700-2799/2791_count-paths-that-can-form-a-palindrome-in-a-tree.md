# 2791 — Count Paths That Can Form A Palindrome In A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPalindromePaths(parent []int, s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Bitmask, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2791: Count Paths That Can Form a Palindrome in a Tree
// https://leetcode.com/problems/count-paths-that-can-form-a-palindrome-in-a-tree/
// Difficulty: Hard
//
// DFS + bitmask. Characters are on edges (s[i] = edge char from parent[i] to i).
// Compute XOR mask from root for each node. Path(u,v) XOR = mask[u] ^ mask[v].
// A palindrome requires at most 1 bit set. Count pairs by iterating all masks
// and for each, counting prior masks that differ by 0 or 1 bit.
// O(N * 26) time, O(N) space.

import "fmt"

func countPalindromePaths(parent []int, s string) int {
	n := len(parent)
  // Membuat matriks/slice 2D untuk DP
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	// Compute mask for each node (XOR of edge chars from root)
  // Alokasi slice integer
	mask := make([]int, n)
	var dfsMask func(node int, cur int)
	dfsMask = func(node int, cur int) {
		mask[node] = cur
		for _, child := range children[node] {
			edgeMask := 1 << (s[child] - 'a')
			dfsMask(child, cur^edgeMask)
		}
	}
	dfsMask(0, 0)

	// Count pairs: iterate masks linearly, counting prior masks with XOR=0 or XOR=1bit
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	result := 0
	for _, m := range mask {
		// XOR = 0: same mask
		result += count[m]
		// XOR has exactly 1 bit: differ by one bit
		for b := 0; b < 26; b++ {
			result += count[m^(1<<b)]
		}
		count[m]++
	}

	return result
}

func main() {
	// LeetCode Example 1: parent=[-1,0,0,1,1,2], s="acaabc" => 8
	fmt.Println(countPalindromePaths([]int{-1, 0, 0, 1, 1, 2}, "acaabc"))
	// LeetCode Example 2: parent=[-1,0,0,0,0], s="aaaaa" => 10
	fmt.Println(countPalindromePaths([]int{-1, 0, 0, 0, 0}, "aaaaa"))
	// Single node (no edges)
	fmt.Println(countPalindromePaths([]int{-1}, "a"))
	// Two nodes
	fmt.Println(countPalindromePaths([]int{-1, 0}, "aa"))
	fmt.Println(countPalindromePaths([]int{-1, 0}, "ab"))
}
```
