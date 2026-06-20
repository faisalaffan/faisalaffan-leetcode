# 1055 — Shortest Way To Form String

## Deskripsi

**Soal:** [1055. Shortest Way To Form String](https://leetcode.com/problems/shortest-way-to-form-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m) worst case, O(n + m) with precomputed indices  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Two Pointer (penunjuk kiri & kanan), Two Pointer (penunjuk kiri & kanan), Greedy (pemilihan optimal lokal)

> **Ide Kunci:** Greedy with two pointers. Count subsequence matches.

## Solusi Go

```go
package main

// LeetCode #1055: Shortest Way to Form String
// https://leetcode.com/problems/shortest-way-to-form-string/
// Difficulty: Medium
//
// Approach: Greedy with two pointers. Count subsequence matches.
// Time: O(n * m) worst case, O(n + m) with precomputed indices
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(shortestWay("abc", "abcbc"))   // 2
	fmt.Println(shortestWay("abc", "acdbc"))   // -1
	fmt.Println(shortestWay("xyz", "xzyxz"))   // 2
}

func shortestWay(source string, target string) int {
	count := 0
	i := 0 // index in target

	// Pre-check: all chars in target must exist in source
  // Membuat map untuk pencarian O(1): key → value
	sourceSet := make(map[byte]bool)
	for k := 0; k < len(source); k++ {
		sourceSet[source[k]] = true
	}
	for k := 0; k < len(target); k++ {
		if !sourceSet[target[k]] {
			return -1
		}
	}

	for i < len(target) {
		count++
		j := 0 // index in source
		for j < len(source) && i < len(target) {
			if source[j] == target[i] {
				i++
			}
			j++
		}
	}

	return count
}
```
