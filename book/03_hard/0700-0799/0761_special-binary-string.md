# 0761 — Special Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func makeLargestSpecial(s string) string
```

> **💡 Hint:** Recursive

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #761: Special Binary String
// https://leetcode.com/problems/special-binary-string/
// Difficulty: Hard
//
// A special binary string is one that:
//   - Has equal number of 0s and 1s
//   - Every prefix has at least as many 1s as 0s
//
// Operation: take two consecutive special substrings and swap them.
// Goal: return the lexicographically largest string achievable.
//
// Approach: Recursive
// 1. For a given special binary string, split it into top-level
//    special substrings (balanced substrings that are not nested
//    inside another balanced substring).
// 2. Recursively process each substring.
// 3. Sort the processed substrings in reverse order (descending).
// 4. Concatenate and return.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(makeLargestSpecial("11011000")) // "11100100"
	fmt.Println(makeLargestSpecial("10"))        // "10"
	fmt.Println(makeLargestSpecial("1010"))      // "1010" (already maximal)
}

func makeLargestSpecial(s string) string {
	n := len(s)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return ""
	}

	// Split into top-level special substrings
	var subs []string
	count := 0
	start := 0
	for i, ch := range s {
		if ch == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			// s[start:i+1] is a special binary string
			// Recursively process the inner part (between 1 and 0)
			inner := makeLargestSpecial(s[start+1 : i])
			subs = append(subs, "1"+inner+"0")
			start = i + 1
		}
	}

	// Sort in descending order (lexicographically largest first)
  // Custom sort dengan comparator
	sort.Slice(subs, func(i, j int) bool {
		return subs[i] > subs[j]
	})

	// Concatenate
	result := ""
	for _, sub := range subs {
		result += sub
	}
	return result
}
```
