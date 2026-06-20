# 2414 — Length Of The Longest Alphabetical Continuous Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestContinuousSubstring(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2414: Length of the Longest Alphabetical Continuous Substring
// https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Scan, count consecutive chars where s[i] == s[i-1] + 1.

import "fmt"

func main() {
	fmt.Println(longestContinuousSubstring("abacaba")) // 2 ("ab")
	fmt.Println(longestContinuousSubstring("abcde"))   // 5
}

func longestContinuousSubstring(s string) int {
	ans, cur := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```
