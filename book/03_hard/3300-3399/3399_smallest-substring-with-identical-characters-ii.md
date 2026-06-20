# 3399 — Smallest Substring With Identical Characters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestSubstringWithIdenticalCharactersIi(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3399: Smallest Substring With Identical Characters II
// https://leetcode.com/problems/smallest-substring-with-identical-characters-ii/
// Difficulty: Hard
//
// Binary search on max run length. Greedy check with runLen/(L+1) flips.

import "fmt"

func main() {
	fmt.Println(SmallestSubstringWithIdenticalCharactersIi("000111", 2))
}

func SmallestSubstringWithIdenticalCharactersIi(s string, k int) int {
	if k == 1 {
		// Alternating pattern: need to find min possible max run
		n := len(s)
		if n <= 1 {
			return n
		}
		// With 1 operation, we can break at most one run by flipping
		// The best we can do is make all runs length 1.
		// If there's a run of length 2+, we can flip the middle.
		hasLongRun := false
		i := 0
		for i < n {
			j := i
			for j < n && s[j] == s[i] {
				j++
			}
			if j-i >= 2 {
				hasLongRun = true
				break
			}
			i = j
		}
		if !hasLongRun {
			return 1
		}
		return 2
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
