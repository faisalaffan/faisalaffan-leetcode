# 0331 — Verify Preorder Serialization Of A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isValidSerialization(preorder string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #331: Verify Preorder Serialization of a Binary Tree
// https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	// Each non-null node consumes 1 slot and creates 2 new slots
	// Each null node consumes 1 slot
	slots := 1
	for _, node := range nodes {
		slots--
		if slots < 0 {
			return false
		}
		if node != "#" {
			slots += 2
		}
	}
	return slots == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", isValidSerialization("1,#"))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", isValidSerialization("9,#,#,1"))
	// Expected: false
}
```
