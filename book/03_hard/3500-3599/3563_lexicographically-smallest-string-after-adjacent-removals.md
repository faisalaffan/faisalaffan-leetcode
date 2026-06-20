# 3563 — Lexicographically Smallest String After Adjacent Removals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func lexicographicallySmallestString(s string) string
```

> **💡 Hint:** Interval DP to compute which substrings can be fully removed,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3563: Lexicographically Smallest String After Adjacent Removals
// https://leetcode.com/problems/lexicographically-smallest-string-after-adjacent-removals/
// Difficulty: Hard
//
// Given a string s, repeatedly remove any adjacent pair of characters that are
// consecutive in the alphabet (circularly, 'a' and 'z' also count). Return the
// lexicographically smallest string achievable after any number of operations.
//
// Approach: Interval DP to compute which substrings can be fully removed,
// then suffix DP to find the lexicographically smallest result.

import "fmt"

func main() {
	// Example 1
	fmt.Println(lexicographicallySmallestString("abc"))
	// Example 2
	fmt.Println(lexicographicallySmallestString("bcda"))
	// Example 3
	fmt.Println(lexicographicallySmallestString("zdce"))
	// Edge: single char
	fmt.Println(lexicographicallySmallestString("a"))
	// Edge: no removable pairs
	fmt.Println(lexicographicallySmallestString("ac"))
}

func lexicographicallySmallestString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// removable[i][j] = can substring s[i..j] be completely removed
  // Membuat matriks/slice 2D untuk DP
	removable := make([][]bool, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range removable {
		removable[i] = make([]bool, n)
	}

	// Length 2: consecutive pair
	for i := 0; i+1 < n; i++ {
		if isConsecutive(s[i], s[i+1]) {
			removable[i][i+1] = true
		}
	}

	// Longer even-length substrings
	for length := 4; length <= n; length += 2 {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			// Case 1: ends are consecutive and interior is removable
			if isConsecutive(s[i], s[j]) && (i+1 > j-1 || removable[i+1][j-1]) {
				removable[i][j] = true
				continue
			}
			// Case 2: split into two removable parts
			for k := i + 1; k < j; k += 2 {
				if removable[i][k] && removable[k+1][j] {
					removable[i][j] = true
					break
				}
			}
		}
	}

	// dp[i] = lexicographically smallest string from suffix i
	dp := make([]string, n+1)
	dp[n] = ""
	for i := n - 1; i >= 0; i-- {
		// Option 1: keep s[i]
		best := string(s[i]) + dp[i+1]

		// Option 2: try to remove s[i..j] entirely
		for j := i + 1; j < n; j++ {
			if removable[i][j] {
				candidate := dp[j+1]
				if candidate < best {
					best = candidate
				}
			}
		}
		dp[i] = best
	}

	return dp[0]
}

func isConsecutive(a, b byte) bool {
	d := int(a - b)
	if d < 0 {
		d = -d
	}
	return d == 1 || d == 25
}
```
