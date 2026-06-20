# 1625 — Lexicographically Smallest String After Applying Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindLexSmallestString(s string, a int, b int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(N^2), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1625: Lexicographically Smallest String After Applying Operations
// https://leetcode.com/problems/lexicographically-smallest-string-after-applying-operations/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLexSmallestString("5525", 9, 2))
	fmt.Println(FindLexSmallestString("74", 5, 1))
	fmt.Println(FindLexSmallestString("0011", 4, 2))
}

func FindLexSmallestString(s string, a int, b int) string {
	// Time: O(N^2), Space: O(N)
	// BFS over all possible states
	n := len(s)
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)
	queue := []string{s}
	seen[s] = true
	smallest := s

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr < smallest {
			smallest = curr
		}

		// Operation 1: add a to odd-position digits
		bytes := []byte(curr)
		for i := 1; i < len(bytes); i += 2 {
			val := int(bytes[i] - '0')
			val = (val + a) % 10
			bytes[i] = byte('0' + val)
		}
		next1 := string(bytes)

		// Operation 2: rotate right by b
		next2 := curr[n-b:] + curr[:n-b]

		for _, next := range []string{next1, next2} {
			if !seen[next] {
				seen[next] = true
				queue = append(queue, next)
			}
		}
	}

	return smallest
}
```
