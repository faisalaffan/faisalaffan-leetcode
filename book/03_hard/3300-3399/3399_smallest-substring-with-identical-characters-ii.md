# 3399 — Smallest Substring With Identical Characters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SmallestSubstringWithIdenticalCharactersIi(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

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
