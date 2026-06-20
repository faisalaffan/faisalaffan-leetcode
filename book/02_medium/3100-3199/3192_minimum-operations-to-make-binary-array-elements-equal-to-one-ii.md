# 3192 — Minimum Operations To Make Binary Array Elements Equal To One Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3192: Minimum Operations to Make Binary Array Elements Equal to One II
// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)
	ans := 0
	flip := 0

	for i := 0; i < n; i++ {
		cur := nums[i] ^ flip
		if cur == 0 {
			ans++
			flip ^= 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{0, 1, 1, 0, 1})) // Expected: 3
	fmt.Println(minOperations([]int{1, 0, 1, 0}))     // Expected: 2
}
```
