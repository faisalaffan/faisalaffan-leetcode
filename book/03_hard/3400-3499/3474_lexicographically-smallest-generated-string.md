# 3474 — Lexicographically Smallest Generated String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func generateString(s string, t string) string
```

> **💡 Hint:** For each position in s, determine the best character

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3474: Lexicographically Smallest Generated String
// https://leetcode.com/problems/lexicographically-smallest-generated-string/
// Difficulty: Hard
//
// Given string s and target t, insert characters into s to make t
// appear as a subsequence. Minimize the resulting string
// lexicographically.
//
// Approach: For each position in s, determine the best character
// to prepend/append that allows t to still be a subsequence.
// Greedy matching from both ends.

import "fmt"

func main() {
	// Example 1
	fmt.Println(generateString("abc", "abc"))
	// Example 2
	fmt.Println(generateString("ab", "ba"))
	// Edge: empty
	fmt.Println(generateString("", "a"))
}

func generateString(s string, t string) string {
	m, n := len(s), len(t)
	// pref[i] = longest prefix of t that is subsequence of s[:i]
  // Alokasi slice integer
	pref := make([]int, m+1)
	ti := 0
	for i := 0; i < m; i++ {
		if ti < n && s[i] == t[ti] {
			ti++
		}
		pref[i+1] = ti
	}

	// suff[i] = longest suffix of t that is subsequence of s[i:]
  // Alokasi slice integer
	suff := make([]int, m+1)
	ti = n - 1
	for i := m - 1; i >= 0; i-- {
		if ti >= 0 && s[i] == t[ti] {
			ti--
		}
		suff[i] = n - 1 - ti
	}

	// Find position where we can insert to complete subsequence
	// The result will be s with chars added to make t a subsequence
	// Always possible by prepending and/or appending characters

	// Build result by prepending missing prefix and appending missing suffix
	prefLen := pref[m]
	suffLen := suff[0]
	res := ""

	// Determine prefix to add
	if prefLen < n {
		res = t[:n-prefLen] + res
	}
	res += s
	// Determine suffix to add
	if suffLen < n && prefLen < n {
		// Already handled by prefix
	}

	return res
}
```
