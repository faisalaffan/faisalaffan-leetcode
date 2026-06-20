# 3694 — Distinct Points Reachable After Substring Removal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func distinctPointsReachableAfterSubstringRemoval(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3694: Distinct Points Reachable After Substring Removal
// https://leetcode.com/problems/distinct-points-reachable-after-substring-removal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func distinctPointsReachableAfterSubstringRemoval(s string, k int) int {
	n := len(s)
	// prefix sums for x and y
  // Alokasi slice integer
	fx := make([]int, n+1)
  // Alokasi slice integer
	fy := make([]int, n+1)
	for i, ch := range s {
		fx[i+1] = fx[i]
		fy[i+1] = fy[i]
		switch ch {
		case 'U':
			fy[i+1]++
		case 'D':
			fy[i+1]--
		case 'L':
			fx[i+1]--
		case 'R':
			fx[i+1]++
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[[2]int]bool)
	for i := k; i <= n; i++ {
		x := fx[n] - (fx[i] - fx[i-k])
		y := fy[n] - (fy[i] - fy[i-k])
		seen[[2]int{x, y}] = true
	}

	return len(seen)
}

func main() {
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("LUL", 1))
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("UDLR", 4))
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("UU", 1))
}
```
