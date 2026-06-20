# 0255 — Verify Preorder Sequence In Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func verifyPreorder(preorder []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Stack

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #255: Verify Preorder Sequence in Binary Search Tree
// https://leetcode.com/problems/verify-preorder-sequence-in-binary-search-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import "fmt"

func verifyPreorder(preorder []int) bool {
	stack := []int{}
	lower := ^int(^uint(0) >> 1) // math.MinInt

	for _, val := range preorder {
		if val < lower {
			return false
		}
		for len(stack) > 0 && val > stack[len(stack)-1] {
			lower = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, val)
	}

	return true
}

func main() {
	fmt.Println(verifyPreorder([]int{5, 2, 1, 3, 6}))
	fmt.Println(verifyPreorder([]int{5, 2, 6, 1, 3}))
	fmt.Println(verifyPreorder([]int{1, 2, 3}))
}
```
