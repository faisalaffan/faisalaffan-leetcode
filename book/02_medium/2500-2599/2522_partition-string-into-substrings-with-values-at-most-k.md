# 2522 — Partition String Into Substrings With Values At Most K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumPartition(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2522: Partition String Into Substrings With Values at Most K
// https://leetcode.com/problems/partition-string-into-substrings-with-values-at-most-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: extend substring while value <= k, then start new partition.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumPartition("165462", 60)) // 4 (16|54|6|2)
	fmt.Println(minimumPartition("238182", 5))  // -1
}

func minimumPartition(s string, k int) int {
	ans := 1
	cur := 0
	for _, ch := range s {
		d := int(ch - '0')
		if d > k {
			return -1
		}
		if cur > math.MaxInt32/10 || cur*10+d > k {
			ans++
			cur = d
		} else {
			cur = cur*10 + d
		}
	}
	return ans
}
```
