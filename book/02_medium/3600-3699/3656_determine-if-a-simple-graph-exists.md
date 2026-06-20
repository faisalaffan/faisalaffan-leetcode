# 3656 — Determine If A Simple Graph Exists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func determineIfASimpleGraphExists(degrees []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3656: Determine if a Simple Graph Exists
// https://leetcode.com/problems/determine-if-a-simple-graph-exists/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func determineIfASimpleGraphExists(degrees []int) bool {
	n := len(degrees)
  // Alokasi slice integer
	arr := make([]int, n)
	copy(arr, degrees)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			break
		}
		if arr[i] > n-i-1 {
			return false
		}
		for j := i + 1; j <= i+arr[i]; j++ {
			arr[j]--
			if arr[j] < 0 {
				return false
			}
		}
		arr[i] = 0
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	}

	return true
}

func main() {
	fmt.Println(determineIfASimpleGraphExists([]int{3, 3, 3, 3}))
	fmt.Println(determineIfASimpleGraphExists([]int{1, 1, 0}))
	fmt.Println(determineIfASimpleGraphExists([]int{1, 1, 1}))
}
```
