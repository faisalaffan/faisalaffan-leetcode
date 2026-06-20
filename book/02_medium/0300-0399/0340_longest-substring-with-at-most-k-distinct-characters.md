# 0340 — Longest Substring With At Most K Distinct Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func lengthOfLongestSubstringKDistinct(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #340: Longest Substring with At Most K Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(k)

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for len(freq) > k {
			freq[s[left]]--
			if freq[s[left]] == 0 {
				delete(freq, s[left])
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
	// Test case 1
	fmt.Println("Test 1:", lengthOfLongestSubstringKDistinct("eceba", 2))
	// Expected: 3 ("ece")

	// Test case 2
	fmt.Println("Test 2:", lengthOfLongestSubstringKDistinct("aa", 1))
	// Expected: 2 ("aa")

	// Test case 3
	fmt.Println("Test 3:", lengthOfLongestSubstringKDistinct("a", 0))
	// Expected: 0
}
```
