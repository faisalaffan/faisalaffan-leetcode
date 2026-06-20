# 1585 — Check If String Is Transformable With Substring Sort Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isTransformable(s string, t string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1585: Check If String Is Transformable With Substring Sort Operations
// https://leetcode.com/problems/check-if-string-is-transformable-with-substring-sort-operations/
// Difficulty: Hard
//
// Key insight: Sorting a substring in ascending order moves smaller digits left
// and larger digits right. Therefore a digit can only move right (not left) in
// the string relative to smaller digits.
//
// Algorithm:
// - Maintain queues of positions for each digit (0-9) in the source string s.
// - Scan t left-to-right. For each digit d:
//   1. Pop the earliest available position pos in s for digit d.
//   2. For every smaller digit sd < d, check that ALL remaining positions
//      of sd in s are AFTER pos. If any remaining sd position is before pos,
//      that smaller digit would block d's movement (since smaller digits move
//      left and would overtake d if they start before it).
// - If all checks pass, return true.

func isTransformable(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Queue of positions for each digit in source string s
  // Membuat matriks/slice 2D untuk DP
	pos := make([][]int, 10)
	for i, ch := range s {
		d := ch - '0'
		pos[d] = append(pos[d], i)
	}

	// Also track counts for quick character validation
	countS := [10]int{}
	countT := [10]int{}
	for _, ch := range s {
		countS[ch-'0']++
	}
	for _, ch := range t {
		countT[ch-'0']++
	}
	if countS != countT {
		return false
	}

	// Scan target string left to right
	for _, ch := range t {
		d := ch - '0'

		// Pop the first (earliest) occurrence of this digit in s
		if len(pos[d]) == 0 {
			return false
		}
		j := pos[d][0]
		pos[d] = pos[d][1:]

		// Check all smaller digits: any remaining position before j is a blocker.
		// A smaller digit before d in the source would, when sorted, end up before d,
		// making it impossible for d to be at the current position.
		for sd := 0; sd < int(d); sd++ {
			if len(pos[sd]) > 0 && pos[sd][0] < j {
				return false
			}
		}
	}

	return true
}

func main() {
	// Example 1:
	// Input: s = "84532", t = "34852"
	// Output: true
	fmt.Println(isTransformable("84532", "34852")) // true

	// Example 2:
	// Input: s = "34521", t = "23415"
	// Output: true
	fmt.Println(isTransformable("34521", "23415")) // true

	// Example 3:
	// Input: s = "12345", t = "12435"
	// Output: false — '4' cannot move right past '5' (4 < 5, and 5 is after 4)
	fmt.Println(isTransformable("12345", "12435")) // false

	// Additional tests:
	// Same strings
	fmt.Println(isTransformable("12345", "12345")) // true

	// "21" -> "12" works (sort entire substring)
	fmt.Println(isTransformable("21", "12")) // true

	// "12" -> "21" fails (smaller can't move right, larger can't move left)
	fmt.Println(isTransformable("12", "21")) // false

	// Length mismatch
	fmt.Println(isTransformable("123", "1234")) // false
}
```
