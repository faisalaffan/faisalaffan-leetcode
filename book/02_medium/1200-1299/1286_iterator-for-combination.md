# 1286 — Iterator For Combination

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor(characters string, combinationLength int) CombinationIterator`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(C(n,k)) init, O(1) next/hasNext  |  **Ruang:** O(C(n,k))

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1286: Iterator for Combination
// https://leetcode.com/problems/iterator-for-combination/
// Difficulty: Medium

// Generate all combinations of length combinationLength from characters
// in lexicographic order using next() and hasNext().

// Time: O(C(n,k)) init, O(1) next/hasNext
// Space: O(C(n,k))

type CombinationIterator struct {
	combinations []string
	idx          int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combs := make([]string, 0)
	n := len(characters)

	var backtrack func(start int, curr []byte)
	backtrack = func(start int, curr []byte) {
		if len(curr) == combinationLength {
			combs = append(combs, string(curr))
			return
		}
		for i := start; i < n; i++ {
			curr = append(curr, characters[i])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []byte{})
	return CombinationIterator{combinations: combs, idx: 0}
}

func (this *CombinationIterator) Next() string {
	res := this.combinations[this.idx]
	this.idx++
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return this.idx < len(this.combinations)
}

func main() {
	iter := Constructor("abc", 2)
	fmt.Printf("next: %q (expected: %q)\n", iter.Next(), "ab")
	fmt.Printf("hasNext: %t (expected: true)\n", iter.HasNext())
	fmt.Printf("next: %q (expected: %q)\n", iter.Next(), "ac")
	fmt.Printf("hasNext: %t (expected: true)\n", iter.HasNext())
	fmt.Printf("next: %q (expected: %q)\n", iter.Next(), "bc")
	fmt.Printf("hasNext: %t (expected: false)\n", iter.HasNext())
}
```
