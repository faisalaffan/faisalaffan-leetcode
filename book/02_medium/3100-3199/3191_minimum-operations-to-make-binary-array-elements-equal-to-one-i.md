# 3191 — Minimum Operations To Make Binary Array Elements Equal To One I

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

// LeetCode #3191: Minimum Operations to Make Binary Array Elements Equal to One I
// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i <= n-3; i++ {
		if nums[i] == 0 {
			ans++
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
		}
	}

	for i := n - 2; i < n; i++ {
		if nums[i] == 0 {
			return -1
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{0, 1, 1, 1, 0, 0})) // Expected: 3
	fmt.Println(minOperations([]int{0, 1, 1, 1}))        // Expected: -1
}
```
