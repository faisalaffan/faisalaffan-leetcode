# 3329 — Count Substrings With K Frequency Characters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubstrings(s string, k int) int64
```

> **💡 Hint:** Sliding window with two pointers. For each right endpoint,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3329: Count Substrings With K-Frequency Characters II
// https://leetcode.com/problems/count-substrings-with-k-frequency-characters-ii/
// Difficulty: Hard [Paid]
//
// Count substrings where at least one character appears at least k times.
//
// Approach: Sliding window with two pointers. For each right endpoint,
// maintain window [left, right] where no character reaches frequency k.
// All substrings starting at [0, left-1] and ending at right are valid.

import "fmt"

func main() {
	// Example 1
	fmt.Println(numberOfSubstrings("abacb", 2))
	// Example 2
	fmt.Println(numberOfSubstrings("abcde", 1))
	// Example 3
	fmt.Println(numberOfSubstrings("aaaaa", 2))
	// Edge: no valid substrings
	fmt.Println(numberOfSubstrings("abc", 5))
	// Single char repeated
	fmt.Println(numberOfSubstrings("aa", 2))
}

func numberOfSubstrings(s string, k int) int64 {
	freq := [26]int{}
	var total int64
	left := 0

	for right := 0; right < len(s); right++ {
		cur := s[right] - 'a'
		freq[cur]++

		for freq[cur] >= k {
			freq[s[left]-'a']--
			left++
		}

		total += int64(left)
	}

	return total
}
```
