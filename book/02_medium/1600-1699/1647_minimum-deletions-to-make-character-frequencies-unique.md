# 1647 — Minimum Deletions To Make Character Frequencies Unique

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinDeletions(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1647: Minimum Deletions to Make Character Frequencies Unique
// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinDeletions("aab"))
	fmt.Println(MinDeletions("aaabbbcc"))
	fmt.Println(MinDeletions("ceabaacb"))
}

func MinDeletions(s string) int {
	// Time: O(N), Space: O(1)
  // Alokasi slice
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

  // HashMap: O(1) lookup
	used := make(map[int]bool)
	deletions := 0

	for _, f := range freq {
		for f > 0 && used[f] {
			f--
			deletions++
		}
		if f > 0 {
			used[f] = true
		}
	}

	return deletions
}
```
