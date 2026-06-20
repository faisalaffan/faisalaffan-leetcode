# 1180 — Count Substrings With Only One Distinct Letter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countLetters(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1180: Count Substrings with Only One Distinct Letter
// https://leetcode.com/problems/count-substrings-with-only-one-distinct-letter/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countLetters("aaaba")) // 8
	fmt.Println(countLetters("aaaaaaaaaa")) // 55
}

// LeetCode submission: countLetters
func countLetters(s string) int {
	ans := 0
	for i, n := 0, len(s); i < n; {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		L := j - i
		ans += L * (L + 1) / 2
		i = j
	}
	return ans
}
```
