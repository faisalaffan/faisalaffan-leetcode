# 1910 — Remove All Occurrences Of A Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func RemoveOccurrences(s string, part string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n*m) where n = len(s), m = len(part), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1910: Remove All Occurrences of a Substring
// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(RemoveOccurrences("daabcbaabcbc", "abc"))
	fmt.Println(RemoveOccurrences("axxxxyyyyb", "xy"))
	fmt.Println(RemoveOccurrences("aabababa", "aba"))
}

// Time: O(n*m) where n = len(s), m = len(part), Space: O(n)
func RemoveOccurrences(s string, part string) string {
	stack := make([]byte, 0)
	m := len(part)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= m && string(stack[len(stack)-m:]) == part {
			stack = stack[:len(stack)-m]
		}
	}
	return string(stack)
}
```
