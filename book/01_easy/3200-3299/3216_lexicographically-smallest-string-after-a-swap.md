# 3216 — Lexicographically Smallest String After A Swap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func LexicographicallySmallestStringAfterASwap(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3216: Lexicographically Smallest String After a Swap
// https://leetcode.com/problems/lexicographically-smallest-string-after-a-swap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestStringAfterASwap("45320"))
	fmt.Println(LexicographicallySmallestStringAfterASwap("001"))
}

// LexicographicallySmallestStringAfterASwap makes the smallest string by swapping one pair of adjacent same-parity digits where left > right.
// Time: O(n). Space: O(n).
func LexicographicallySmallestStringAfterASwap(s string) string {
	b := []byte(s)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(b)-1; i++ {
		if b[i] > b[i+1] && (b[i]%2 == b[i+1]%2) {
			b[i], b[i+1] = b[i+1], b[i]
			break
		}
	}
	return string(b)
}
```
