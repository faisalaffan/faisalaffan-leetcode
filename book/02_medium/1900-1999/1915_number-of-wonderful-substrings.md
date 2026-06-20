# 1915 — Number Of Wonderful Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func WonderfulSubstrings(word string) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(n * 10) = O(n), Space: O(2^10) = O(1)  
**Kompleksitas Ruang:** O(2^10) = O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1915: Number of Wonderful Substrings
// https://leetcode.com/problems/number-of-wonderful-substrings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WonderfulSubstrings("aba"))
	fmt.Println(WonderfulSubstrings("aabb"))
	fmt.Println(WonderfulSubstrings("he"))
}

// Time: O(n * 10) = O(n), Space: O(2^10) = O(1)
func WonderfulSubstrings(word string) int64 {
	// mask represents parity of each of 10 letters
  // Alokasi slice integer
	count := make([]int64, 1024) // 2^10 possible masks
	count[0] = 1
	mask := 0
	var result int64 = 0

	for _, c := range word {
		mask ^= 1 << (c - 'a')

		// Count substrings where all letters have even count
		result += count[mask]

		// Count substrings where exactly one letter has odd count
		for i := 0; i < 10; i++ {
			result += count[mask^(1<<i)]
		}

		count[mask]++
	}
	return result
}
```
