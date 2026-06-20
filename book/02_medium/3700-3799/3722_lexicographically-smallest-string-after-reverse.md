# 3722 — Lexicographically Smallest String After Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func reverseSubstring(s string, start int, end int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3722: Lexicographically Smallest String After Reverse
// https://leetcode.com/problems/lexicographically-smallest-string-after-reverse/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func reverseSubstring(s string, start int, end int) string {
	b := []byte(s)
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func lexicographicallySmallestStringAfterReverse(s string) string {
	n := len(s)
	best := s

	// Reverse first k
	for k := 1; k <= n; k++ {
		candidate := reverseSubstring(s, 0, k-1)
		if candidate < best {
			best = candidate
		}
	}

	// Reverse last k
	for k := 1; k <= n; k++ {
		candidate := reverseSubstring(s, n-k, n-1)
		if candidate < best {
			best = candidate
		}
	}

	return best
}

func main() {
	fmt.Println(lexicographicallySmallestStringAfterReverse("dcab"))
	fmt.Println(lexicographicallySmallestStringAfterReverse("abba"))
	fmt.Println(lexicographicallySmallestStringAfterReverse("zxy"))
}
```
