# 2021 — Brightest Position On Street

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func brightestPosition(lights [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2021: Brightest Position on Street
// https://leetcode.com/problems/brightest-position-on-street/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func brightestPosition(lights [][]int) int {
  // Alokasi slice integer
	events := make([][2]int, 0, len(lights)*2)

	for _, l := range lights {
		pos, rng := l[0], l[1]
		events = append(events, [2]int{pos - rng, 1})
		events = append(events, [2]int{pos + rng + 1, -1})
	}

  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	maxBrightness := 0
	currBrightness := 0
	bestPos := events[0][0]

	for _, e := range events {
		currBrightness += e[1]
		if currBrightness > maxBrightness {
			maxBrightness = currBrightness
			bestPos = e[0]
		}
	}

	return bestPos
}

func main() {
	// Test case 1
	lights1 := [][]int{{-3, 2}, {1, 2}, {3, 3}}
	fmt.Println("Test 1:", brightestPosition(lights1))
	// Expected: -1

	// Test case 2
	lights2 := [][]int{{1, 0}, {0, 1}}
	fmt.Println("Test 2:", brightestPosition(lights2))
	// Expected: 1

	// Test case 3
	lights3 := [][]int{{1, 2}}
	fmt.Println("Test 3:", brightestPosition(lights3))
	// Expected: -1
}
```
