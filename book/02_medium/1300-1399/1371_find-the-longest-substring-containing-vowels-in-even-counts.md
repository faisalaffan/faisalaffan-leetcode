# 1371 — Find The Longest Substring Containing Vowels In Even Counts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func findTheLongestSubstring(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) where n = length of string  |  **Ruang:** O(1) - fixed array of 32 states (2^5)


## 💻 Solusi Go

```go
package main

// LeetCode #1371: Find the Longest Substring Containing Vowels in Even Counts
// https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(findTheLongestSubstring("eleetminicoworoep")) // 13

	// Test case 2
	fmt.Println(findTheLongestSubstring("leetcodeisgreat")) // 5

	// Test case 3
	fmt.Println(findTheLongestSubstring("bcbcbc")) // 6
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 32 states (2^5)
func findTheLongestSubstring(s string) int {
	// Map bitmask (5 bits for a,e,i,o,u) to first occurrence index
	// State 0 (all vowels even) at position 0
  // Alokasi slice
	firstSeen := make([]int, 32)
  // Range loop
	for i := range firstSeen {
		firstSeen[i] = -1
	}
	firstSeen[0] = 0

	mask := 0
	maxLen := 0

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			mask ^= 1 << 0
		case 'e':
			mask ^= 1 << 1
		case 'i':
			mask ^= 1 << 2
		case 'o':
			mask ^= 1 << 3
		case 'u':
			mask ^= 1 << 4
		}

		if firstSeen[mask] == -1 {
			firstSeen[mask] = i + 1
		} else {
			length := i + 1 - firstSeen[mask]
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}
```
