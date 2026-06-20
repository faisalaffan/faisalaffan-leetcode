# 1653 — Minimum Deletions To Make String Balanced

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumDeletions(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1653: Minimum Deletions to Make String Balanced
// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumDeletions("aababbab"))
	fmt.Println(MinimumDeletions("bbaaaaabb"))
	fmt.Println(MinimumDeletions("a"))
}

func MinimumDeletions(s string) int {
	// Time: O(N), Space: O(1)
	// Count 'a's on the right
	aCount := 0
	for _, ch := range s {
		if ch == 'a' {
			aCount++
		}
	}

	bCount := 0
	minDeletions := len(s)

	for _, ch := range s {
		if ch == 'a' {
			aCount--
		}

		// Deletions needed: remove all 'b's before this point + remove all 'a's after
		deletions := bCount + aCount
		if deletions < minDeletions {
			minDeletions = deletions
		}

		if ch == 'b' {
			bCount++
		}
	}

	return minDeletions
}
```
