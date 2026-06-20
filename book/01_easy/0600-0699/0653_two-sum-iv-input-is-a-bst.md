# 0653 — Two Sum Iv Input Is A Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func findTarget(root *TreeNode, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #653: Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// Difficulty: Easy

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [5,3,6,2,4,null,7], k=9 => true
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findTarget(root, 9))  // true
	fmt.Println(findTarget(root, 28)) // false
}

// findTarget returns true if there exist two elements in the BST that sum to k.
// Time: O(n). Space: O(n).
func findTarget(root *TreeNode, k int) bool {
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	return find(root, k, seen)
}

func find(node *TreeNode, k int, seen map[int]bool) bool {
	if node == nil {
		return false
	}
	if seen[k-node.Val] {
		return true
	}
	seen[node.Val] = true
	return find(node.Left, k, seen) || find(node.Right, k, seen)
}
```
