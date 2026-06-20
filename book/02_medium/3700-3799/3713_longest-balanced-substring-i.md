# 3713 — Longest Balanced Substring I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestBalancedSubstringI(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3713: Longest Balanced Substring I
// https://leetcode.com/problems/longest-balanced-substring-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func longestBalancedSubstringI(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		var cnt [26]int
		distinct := 0
		maxFreq := 0
		for j := i; j < n; j++ {
			idx := s[j] - 'a'
			cnt[idx]++
			if cnt[idx] == 1 {
				distinct++
			}
			if cnt[idx] > maxFreq {
				maxFreq = cnt[idx]
			}
			if maxFreq*distinct == j-i+1 {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalancedSubstringI("abbac"))
	fmt.Println(longestBalancedSubstringI("zzabccy"))
	fmt.Println(longestBalancedSubstringI("aba"))
}
```
