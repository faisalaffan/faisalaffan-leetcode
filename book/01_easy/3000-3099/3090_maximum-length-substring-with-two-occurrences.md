# 3090 — Maximum Length Substring With Two Occurrences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumLengthSubstringWithTwoOccurrences(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3090: Maximum Length Substring With Two Occurrences
// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maximumLengthSubstring
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("bcbbbcba")) // 4
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("aaaa"))      // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: maximumLengthSubstring
func MaximumLengthSubstringWithTwoOccurrences(s string) int {
	left := 0
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[byte]int)
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for freq[s[right]] > 2 {
			freq[s[left]]--
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```
