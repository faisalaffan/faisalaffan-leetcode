# 0828 — Count Unique Characters Of All Substrings Of A Given String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func uniqueLetterString(s string) int
```

> **💡 Hint:** Contribution per character. For each s[i], count substrings where s[i] is the

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #828: Count Unique Characters of All Substrings of a Given String
// https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/
// Difficulty: Hard
// Approach: Contribution per character. For each s[i], count substrings where s[i] is the
// first occurrence of that character within the substring. Use prev/next occurrence arrays.

import "fmt"

func uniqueLetterString(s string) int {
	n := len(s)
  // Alokasi slice integer
	prev := make([]int, n)
  // Alokasi slice integer
	next := make([]int, n)
  // Alokasi slice integer
	last := make([]int, 26)

  // Range loop: iterasi dengan indeks + nilai
	for i := range last {
		last[i] = -1
	}
	for i := 0; i < n; i++ {
		c := int(s[i] - 'A')
		prev[i] = last[c]
		last[c] = i
	}
  // Range loop: iterasi dengan indeks + nilai
	for i := range last {
		last[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		c := int(s[i] - 'A')
		next[i] = last[c]
		last[c] = i
	}

	ans := 0
	for i := 0; i < n; i++ {
		left := i - prev[i]
		right := next[i] - i
		ans += left * right
	}
	return ans
}

func main() {
	fmt.Println(uniqueLetterString("ABC")) // Expected: 10
	fmt.Println(uniqueLetterString("ABA")) // Expected: 8
	fmt.Println(uniqueLetterString("LEETCODE")) // Additional test
}
```
