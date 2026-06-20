# 1055 — Shortest Way To Form String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func shortestWay(source string, target string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m) worst case, O(n + m) with precomputed indices  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
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
