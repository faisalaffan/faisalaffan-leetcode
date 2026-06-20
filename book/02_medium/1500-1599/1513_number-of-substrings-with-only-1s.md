# 1513 — Number Of Substrings With Only 1S

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumSub(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1513: Number of Substrings With Only 1s
// https://leetcode.com/problems/number-of-substrings-with-only-1s/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSub("0110111"))
	fmt.Println(NumSub("101"))
	fmt.Println(NumSub("111111"))
}

func NumSub(s string) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	count := 0
	consecutive := 0

	for _, ch := range s {
		if ch == '1' {
			consecutive++
			// Each new 1 adds 'consecutive' new substrings ending at this position
			count = (count + consecutive) % mod
		} else {
			consecutive = 0
		}
	}

	return count
}
```
