# 1018 — Binary Prefix Divisible By 5

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func prefixesDivBy5(nums []int) []bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (excluding output)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1018: Binary Prefix Divisible By 5
// https://leetcode.com/problems/binary-prefix-divisible-by-5/
// Difficulty: Easy
// Time: O(n) | Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))          // [true,false,false]
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))          // [false,false,false]
	fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1})) // [true,false,false,false,true,false]
}

// LeetCode submission: prefixesDivBy5
func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	val := 0
	for i, b := range nums {
		val = (val*2 + b) % 5
		ans[i] = val == 0
	}
	return ans
}
```
