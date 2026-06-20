# 0109 — Convert Sorted List To Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func sortedListToBST(head *ListNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #109: Convert Sorted List to Binary Search Tree
// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		node := &TreeNode{Val: nums[mid]}
		node.Left = build(left, mid-1)
		node.Right = build(mid+1, right)
		return node
	}

	return build(0, len(nums)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Test case 1: [-10,-3,0,5,9]
	head := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	root := sortedListToBST(head)
	printInorder(root) // -10 -3 0 5 9
	fmt.Println()

	// Test case 2: [] -> nil
	root = sortedListToBST(nil)
	printInorder(root)
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```
