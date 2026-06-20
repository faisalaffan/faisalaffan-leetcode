# 2380 — Time Needed To Rearrange A Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func secondsToRemoveOccurrences(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2380: Time Needed to Rearrange a Binary String
// https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Each "01" becomes "10" per second. Equivalent to: each 1 moves right past 0s,
// and the total time is max over each 1 of (position - index_in_final_position).

import "fmt"

func main() {
	fmt.Println(secondsToRemoveOccurrences("0110101")) // 4
	fmt.Println(secondsToRemoveOccurrences("11100"))   // 0
	fmt.Println(secondsToRemoveOccurrences("001011"))  // 3
}

func secondsToRemoveOccurrences(s string) int {
	ans, zeros := 0, 0
	for _, ch := range s {
		if ch == '0' {
			zeros++
		} else if zeros > 0 {
			ans = max(ans+1, zeros)
		}
	}
	return ans
}
```
