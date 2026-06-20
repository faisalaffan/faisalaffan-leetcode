# 2489 — Number Of Substrings With Fixed Ratio

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func fixedRatio(s string, num1 int, num2 int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2489: Number of Substrings With Fixed Ratio
// https://leetcode.com/problems/number-of-substrings-with-fixed-ratio/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Count substrings where count('0') : count('1') = num1 : num2.
// Transform: track (num2 * cnt0 - num1 * cnt1), count equal values.

import "fmt"

func main() {
	fmt.Println(fixedRatio("01001", 2, 3)) // 2
	fmt.Println(fixedRatio("0000", 1, 1))  // 0
}

func fixedRatio(s string, num1 int, num2 int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	prefix := make(map[int]int64)
	prefix[0] = 1
	var cnt0, cnt1 int64
	var ans int64

	for _, ch := range s {
		if ch == '0' {
			cnt0++
		} else {
			cnt1++
		}
		key := num2*int(cnt0) - num1*int(cnt1)
		ans += prefix[key]
		prefix[key]++
	}
	return ans
}
```
