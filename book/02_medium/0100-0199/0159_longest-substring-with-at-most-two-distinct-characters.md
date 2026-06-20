# 0159 — Longest Substring With At Most Two Distinct Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func lengthOfLongestSubstringTwoDistinct(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #159: Longest Substring with At Most Two Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
  // Membuat map (HashMap) — pencarian O(1)
	charCount := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		for len(charCount) > 2 {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abc"))
}
```
