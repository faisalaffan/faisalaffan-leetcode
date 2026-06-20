# 3309 — Maximum Possible Number By Binary Concatenation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxGoodNumber(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3309: Maximum Possible Number by Binary Concatenation
// https://leetcode.com/problems/maximum-possible-number-by-binary-concatenation/
// Difficulty: Medium
// Time: O(1) Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxGoodNumber([]int{1, 2, 3}))    // 30
	fmt.Println(maxGoodNumber([]int{2, 8, 16}))   // 1296
	fmt.Println(maxGoodNumber([]int{1, 1, 1}))    // 7
}

func maxGoodNumber(nums []int) int {
	// Try all 6 permutations
	perms := [][]int{
		{nums[0], nums[1], nums[2]},
		{nums[0], nums[2], nums[1]},
		{nums[1], nums[0], nums[2]},
		{nums[1], nums[2], nums[0]},
		{nums[2], nums[0], nums[1]},
		{nums[2], nums[1], nums[0]},
	}

	maxVal := 0
	for _, p := range perms {
		val := 0
		for _, x := range p {
			bits := 0
			temp := x
			for temp > 0 {
				bits++
				temp >>= 1
			}
			if x == 0 {
				bits = 1
			}
			val = (val << bits) | x
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
```
