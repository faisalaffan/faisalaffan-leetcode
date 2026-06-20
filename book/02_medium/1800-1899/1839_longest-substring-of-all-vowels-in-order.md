# 1839 — Longest Substring Of All Vowels In Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestBeautifulSubstring(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1839: Longest Substring Of All Vowels in Order
// https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func longestBeautifulSubstring(word string) int {
	vowels := "aeiou"
	maxLen := 0
	i := 0
	n := len(word)

	for i < n {
		// Start of a new substring
		vowelIdx := 0
		start := i

		// Check if starts with 'a'
		if word[i] != 'a' {
			i++
			continue
		}

		for i < n && vowelIdx < 5 {
			if word[i] == vowels[vowelIdx] {
				i++
			} else if vowelIdx+1 < 5 && word[i] == vowels[vowelIdx+1] {
				vowelIdx++
				i++
			} else {
				break
			}
		}

		if vowelIdx == 4 {
			length := i - start
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu")) // Expected: 13
	fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou")) // Expected: 5
	fmt.Println(longestBeautifulSubstring("aaaa")) // Expected: 0 (no 'e')
}
```
