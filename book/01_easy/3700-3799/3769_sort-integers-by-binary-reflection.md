# 3769 — Sort Integers By Binary Reflection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func SortIntegersByBinaryReflection(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3769: Sort Integers by Binary Reflection
// https://leetcode.com/problems/sort-integers-by-binary-reflection/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortIntegersByBinaryReflection([]int{3, 6, 5}))
	fmt.Println(SortIntegersByBinaryReflection([]int{1, 2, 3}))
}

// Time: O(n log n)
// Space: O(n)
func SortIntegersByBinaryReflection(nums []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	reflections := make(map[int]int)
	for _, v := range nums {
		reflections[v] = reverseBits(v)
	}

  // Custom sort dengan comparator
	sort.Slice(nums, func(i, j int) bool {
		if reflections[nums[i]] != reflections[nums[j]] {
			return reflections[nums[i]] < reflections[nums[j]]
		}
		return nums[i] < nums[j]
	})
	return nums
}

func reverseBits(n int) int {
	rev := 0
	for n > 0 {
		rev = (rev << 1) | (n & 1)
		n >>= 1
	}
	return rev
}
```
