# 0411 — Minimum Unique Word Abbreviation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minAbbreviation(target string, dictionary []string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #411: Minimum Unique Word Abbreviation
// https://leetcode.com/problems/minimum-unique-word-abbreviation/
// Difficulty: Hard [Paid]
//
// Enumerate all 2^n abbreviation patterns via bitmask. A bit = 1 means the
// character is abbreviated (counted as part of a number); bit = 0 means the
// character is kept literally. For each mask, check if ANY dictionary word
// of the same length matches all literal positions. If none matches, the
// abbreviation is valid. Return the shortest (and lexicographically smallest
// among ties).

import (
	"fmt"
	"strconv"
)

func main() {
	// Example 1: target="apple", dictionary=["blade"] -> "a4"
	fmt.Println(minAbbreviation("apple", []string{"blade"}))
	// Example 2: target="apple", dictionary=["plain", "amber", "blade"] -> "1p3"
	fmt.Println(minAbbreviation("apple", []string{"plain", "amber", "blade"}))
	// Example 3: target="usa", dictionary=["usa"] -> "3" (no unique abbr shorter)
	fmt.Println(minAbbreviation("usaandchinaaregreat", []string{"usaandchinaaregreat"}))
	// Edge: no dictionary
	fmt.Println(minAbbreviation("hello", []string{}))
	// Edge: same length but different chars
	fmt.Println(minAbbreviation("abcde", []string{"fghij"}))
}

func minAbbreviation(target string, dictionary []string) string {
	n := len(target)

	// Filter to same-length words only
	sameLen := make([]string, 0, len(dictionary))
	for _, w := range dictionary {
		if len(w) == n {
			sameLen = append(sameLen, w)
		}
	}

	// If no same-length words, just abbreviate entire word
	if len(sameLen) == 0 {
		return strconv.Itoa(n)
	}

	best := target // worst case: no abbreviation
	masks := 1 << uint(n)

	for mask := 0; mask < masks; mask++ {
		// Check if any dict word matches at all literal positions
		conflict := false
		for _, w := range sameLen {
			match := true
			for i := 0; i < n; i++ {
				if mask>>uint(i)&1 == 0 { // literal position
					if target[i] != w[i] {
						match = false
						break
					}
				}
			}
			if match {
				conflict = true
				break
			}
		}
		if conflict {
			continue
		}

		ab := abbreviate(target, mask)
		// Shorter is better; ties: lexicographically smaller
		if len(ab) < len(best) || (len(ab) == len(best) && ab < best) {
			best = ab
		}
	}

	return best
}

func abbreviate(word string, mask int) string {
	n := len(word)
	var res []byte
	count := 0

	for i := 0; i < n; i++ {
		if mask>>uint(i)&1 == 1 {
			count++
		} else {
			if count > 0 {
				res = append(res, strconv.Itoa(count)...)
				count = 0
			}
			res = append(res, word[i])
		}
	}
	if count > 0 {
		res = append(res, strconv.Itoa(count)...)
	}
	return string(res)
}
```
