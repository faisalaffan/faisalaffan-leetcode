# 2743 — Count Substrings Without Repeating Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSubstringsWithoutRepeatingCharacter(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2743: Count Substrings Without Repeating Character
// https://leetcode.com/problems/count-substrings-without-repeating-character/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func CountSubstringsWithoutRepeatingCharacter(s string) int {
  // Membuat map (HashMap) — pencarian O(1)
	lastPos := make(map[byte]int)
	left := 0
	count := 0

	for right := 0; right < len(s); right++ {
		if pos, ok := lastPos[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		lastPos[s[right]] = right
		count += right - left + 1
	}

	return count
}

func main() {
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("abcabc"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("aaaa"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter(""))
}
```
