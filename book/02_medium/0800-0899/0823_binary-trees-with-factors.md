# 0823 — Binary Trees With Factors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func BinaryTreesWithFactors(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #823: Binary Trees With Factors
// https://leetcode.com/problems/binary-trees-with-factors/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BinaryTreesWithFactors([]int{2, 4}))
	fmt.Println(BinaryTreesWithFactors([]int{2, 4, 5, 10}))
	fmt.Println(BinaryTreesWithFactors([]int{2, 3, 4, 6, 8, 12, 24}))
}

// Time: O(n^2) | Space: O(n)
func BinaryTreesWithFactors(arr []int) int {
	const mod = 1_000_000_007
  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr)

  // Membuat map (HashMap) — pencarian O(1)
	dp := make(map[int]int)
	for _, x := range arr {
		dp[x] = 1
	}

	for i, x := range arr {
		for j := 0; j < i; j++ {
			if x%arr[j] == 0 {
				if val, ok := dp[x/arr[j]]; ok {
					dp[x] = (dp[x] + dp[arr[j]]*val) % mod
				}
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return ans
}
```
