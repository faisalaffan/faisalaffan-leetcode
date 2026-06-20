# 2564 — Substring Xor Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func substringXorQueries(s string, queries [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * 31 + q)  
**Kompleksitas Ruang:** O(n * 31)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2564: Substring XOR Queries
// https://leetcode.com/problems/substring-xor-queries/
// Difficulty: Medium
// Time: O(n * 31 + q) | Space: O(n * 31)

import "fmt"

func substringXorQueries(s string, queries [][]int) [][]int {
	// For each possible value, store earliest [l, r]
	n := len(s)
  // Membuat map (HashMap) — pencarian O(1)
	posMap := make(map[int][2]int)

	// For each starting position, compute values up to 31 bits (since val <= 10^9 < 2^30)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			if _, ok := posMap[0]; !ok {
				posMap[0] = [2]int{i, i}
			}
			continue
		}
		val := 0
		for j := i; j < n && j-i < 31; j++ {
			val = (val << 1) | int(s[j]-'0')
			if _, ok := posMap[val]; !ok {
				posMap[val] = [2]int{i, j}
			}
		}
	}

  // Membuat matriks/slice 2D untuk DP
	ans := make([][]int, len(queries))
	for idx, q := range queries {
		first, second := q[0], q[1]
		target := first ^ second
		if pos, ok := posMap[target]; ok {
			ans[idx] = []int{pos[0], pos[1]}
		} else {
			ans[idx] = []int{-1, -1}
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", substringXorQueries("101101", [][]int{{0, 5}, {1, 2}}))
	// Expected: [[0,2],[2,3]]

	// Test case 2
	fmt.Println("Test 2:", substringXorQueries("0101", [][]int{{12, 8}}))
	// value=12^8=4 (100), need substring "100"

	// Test case 3
	fmt.Println("Test 3:", substringXorQueries("1", [][]int{{0, 0}}))
	// value=0^0=0, need substring "0"
}
```
