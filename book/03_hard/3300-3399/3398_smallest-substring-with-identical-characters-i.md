# 3398 — Smallest Substring With Identical Characters I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestSubstringWithIdenticalCharactersI(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3398: Smallest Substring With Identical Characters I
// https://leetcode.com/problems/smallest-substring-with-identical-characters-i/
// Difficulty: Hard
//
// Binary search on max run length. Greedy check with runLen/(L+1) flips.

import "fmt"

func main() {
	fmt.Println(SmallestSubstringWithIdenticalCharactersI("000111", 1))
}

func SmallestSubstringWithIdenticalCharactersI(s string, k int) int {
	if k == 1 {
		return 1
	}

	n := len(s)
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if feasible(s, k, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func feasible(s string, k, limit int) bool {
	cnt := 0
	i := 0
	n := len(s)
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		cnt += runLen / (limit + 1)
		if cnt > k {
			return false
		}
		i = j
	}
	return cnt <= k
}
```
