# 1456 — Maximum Number Of Vowels In A Substring Of Given Length

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isVowel(ch byte) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n) where n = length of string  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1456: Maximum Number of Vowels in a Substring of Given Length
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxVowels("abciiidef", 3)) // 3

	// Test case 2
	fmt.Println(maxVowels("aeiou", 2)) // 2

	// Test case 3
	fmt.Println(maxVowels("leetcode", 3)) // 2

	// Test case 4
	fmt.Println(maxVowels("rhythms", 4)) // 0
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

// Time: O(n) where n = length of string
// Space: O(1)
func maxVowels(s string, k int) int {
	// Count vowels in first window
	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	maxCount := count

	// Sliding window
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			count--
		}
		if isVowel(s[i]) {
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
```
