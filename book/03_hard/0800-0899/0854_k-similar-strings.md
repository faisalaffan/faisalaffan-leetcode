# 0854 — K Similar Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func kSimilarity(s1 string, s2 string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #854: K-Similar Strings
// https://leetcode.com/problems/k-similar-strings/
// Difficulty: Hard
// Approach: BFS over string permutations. At each state, pick the first mismatching
// position i and swap it with any later position j that would fix the mismatch.

import "fmt"

func kSimilarity(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}

	n := len(s1)
  // HashMap: O(1) lookup
	visited := make(map[string]bool)
	queue := []string{s1}
	visited[s1] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			cur := queue[k]
			if cur == s2 {
				return steps
			}

			// Find first mismatching position
			i := 0
			for i < n && cur[i] == s2[i] {
				i++
			}

			// Try swapping with any position j > i where cur[j] == s2[i]
			// (this guarantees we fix position i)
			for j := i + 1; j < n; j++ {
				if cur[j] == s2[i] && cur[j] != s2[j] {
					next := []byte(cur)
					next[i], next[j] = next[j], next[i]
					str := string(next)
					if !visited[str] {
						visited[str] = true
						queue = append(queue, str)
					}
				}
			}
		}
		queue = queue[size:]
		steps++
	}

	return -1
}

func main() {
	fmt.Println(kSimilarity("ab", "ba"))       // Expected: 1
	fmt.Println(kSimilarity("abc", "bca"))      // Expected: 2
	fmt.Println(kSimilarity("abac", "baca"))    // Expected: 2
}
```
