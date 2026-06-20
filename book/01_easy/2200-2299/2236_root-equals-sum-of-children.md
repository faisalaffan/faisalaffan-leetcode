# 2236 — Root Equals Sum Of Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func RootEqualsSumOfChildren(root *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2236: Root Equals Sum of Children
// https://leetcode.com/problems/root-equals-sum-of-children/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{10, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root)) // true

	root2 := &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root2)) // false
}

// Time: O(1), Space: O(1)
func RootEqualsSumOfChildren(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
```
