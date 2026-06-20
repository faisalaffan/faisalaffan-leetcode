# 1130 — Minimum Cost Tree From Leaf Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func mctFromLeafValues(arr []int) int
```

> **💡 Hint:** Monotonic decreasing stack (greedy)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1130: Minimum Cost Tree From Leaf Values
// https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/
// Difficulty: Medium
//
// Approach: Monotonic decreasing stack (greedy)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))    // 32
	fmt.Println(mctFromLeafValues([]int{4, 11}))      // 44
}

func mctFromLeafValues(arr []int) int {
  // Alokasi slice integer
	stack := make([]int, 0)
	result := 0

	for _, v := range arr {
		for len(stack) > 0 && stack[len(stack)-1] <= v {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				result += mid * v
			} else {
				if stack[len(stack)-1] < v {
					result += mid * stack[len(stack)-1]
				} else {
					result += mid * v
				}
			}
		}
		stack = append(stack, v)
	}

	for len(stack) > 1 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result += last * stack[len(stack)-1]
	}

	return result
}
```
