# 3106 — Lexicographically Smallest String After Operations With Constraint

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func getSmallestString(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3106: Lexicographically Smallest String After Operations With Constraint
// https://leetcode.com/problems/lexicographically-smallest-string-after-operations-with-constraint/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getSmallestString(s string, k int) string {
	if k == 0 {
		return s
	}

	bytes := []byte(s)
	for i, ch := range bytes {
		dist := int(ch - 'a')
		move := min(dist, 26-dist)
		if move <= k {
			k -= move
			bytes[i] = 'a'
		} else {
			bytes[i] = byte(int(ch) - k)
			k = 0
			break
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(getSmallestString("zbbz", 3))  // Expected: "aaaz"
	fmt.Println(getSmallestString("xaxcd", 4)) // Expected: "aawcd"
	fmt.Println(getSmallestString("lol", 0))   // Expected: "lol"
}
```
