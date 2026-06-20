# 0893 — Groups Of Special Equivalent Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func GroupsOfSpecialEquivalentStrings(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m) where n = len(words), m = avg word length  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #893: Groups of Special-Equivalent Strings
// https://leetcode.com/problems/groups-of-special-equivalent-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"a"}))
}

// Time: O(n * m) where n = len(words), m = avg word length | Space: O(n)
func GroupsOfSpecialEquivalentStrings(words []string) int {
  // HashMap: O(1) lookup
	groups := make(map[[52]int]bool)

	for _, word := range words {
		var key [52]int
		for i, c := range word {
			// Even indices: 0-25, Odd indices: 26-51
			key[int(c-'a')+26*(i%2)]++
		}
		groups[key] = true
	}

	return len(groups)
}
```
