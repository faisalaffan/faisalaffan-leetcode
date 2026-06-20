# 1967 — Number Of Strings That Appear As Substrings In Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfStringsThatAppearAsSubstringsInWord(patterns []string, word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1967: Number of Strings That Appear as Substrings in Word
// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "abc", "bc", "d"}, "abc")) // 3
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "b", "c"}, "aaaaabbbbb"))  // 2
}

// Time: O(n * m), Space: O(1)
func NumberOfStringsThatAppearAsSubstringsInWord(patterns []string, word string) int {
	count := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			count++
		}
	}
	return count
}
```
