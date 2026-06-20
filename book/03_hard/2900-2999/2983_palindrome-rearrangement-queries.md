# 2983 — Palindrome Rearrangement Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func canMakePalindromeQueries(s string, queries [][]int) []bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2983: Palindrome Rearrangement Queries
// https://leetcode.com/problems/palindrome-rearrangement-queries/
// Difficulty: Hard
//
// Given string s of even length, queries [a,b,c,d]. For each query, we can
// rearrange characters in s[a..b] and s[c..d]. Can s become a palindrome?
//
// Split into two halves: left = s[:mid], right = reverse(s[mid:]).
// For a palindrome, each position i in left must match position i in right.
// Map query ranges onto the first-half index space and analyze character
// coverage overlaps.

import (
	"fmt"
)

func canMakePalindromeQueries(s string, queries [][]int) []bool {
	n := len(s)
	mid := n / 2

	// First half: s[0..mid-1]
	// Second half reversed: reverse(s[mid..n-1])
	a := s[:mid]
	b := reverseStr(s[mid:])

	// Prefix sums for character counts
  // Alokasi slice integer
	prefA := make([][26]int, mid+1)
  // Alokasi slice integer
	prefB := make([][26]int, mid+1)
	for i := 0; i < mid; i++ {
		prefA[i+1] = prefA[i]
		prefA[i+1][a[i]-'a']++
		prefB[i+1] = prefB[i]
		prefB[i+1][b[i]-'a']++
	}

	// diff[i] = number of mismatched positions in prefix [0, i-1]
  // Alokasi slice integer
	diff := make([]int, mid+1)
	for i := 0; i < mid; i++ {
		diff[i+1] = diff[i]
		if a[i] != b[i] {
			diff[i+1]++
		}
	}

	// Helper: count of each character in [l, r]
	cntRange := func(pref [][26]int, l, r int) [26]int {
		var res [26]int
		if l > r {
			return res
		}
		for i := 0; i < 26; i++ {
			res[i] = pref[r+1][i] - pref[l][i]
		}
		return res
	}

	ans := make([]bool, len(queries))
	for qi, q := range queries {
		qa, qb, qc, qd := q[0], q[1], q[2], q[3]

		// Map query ranges to the first-half index space
		// For positions >= mid, they map to n-1-pos in the reversed second half
		rb := n - 1 - qd
		re := n - 1 - qc
		l1, r1 := qa, qb
		l2, r2 := rb, re

		if l1 < 0 || r1 >= mid || l2 < 0 || r2 >= mid {
			ans[qi] = false
			continue
		}

		// Ensure l1 <= l2 for simpler case analysis
		if l1 > l2 {
			l1, r1, l2, r2 = l2, r2, l1, r1
		}

		// Check that positions outside both intervals match
		if diff[l1] > 0 || diff[mid]-diff[max2(r1, r2)+1] > 0 {
			ans[qi] = false
			continue
		}

		// Check if the intervals overlap
		inBoth := max2(l1, l2) <= min2(r1, r2)

		if !inBoth {
			// Non-overlapping: each range needs to fix its own mismatches
			cntA1 := cntRange(prefA, l1, r1)
			cntB1 := cntRange(prefB, l1, r1)
			cntA2 := cntRange(prefA, l2, r2)
			cntB2 := cntRange(prefB, l2, r2)
			ok := true
			for c := 0; c < 26; c++ {
				if cntA1[c] != cntB1[c] || cntA2[c] != cntB2[c] {
					ok = false
					break
				}
			}
			ans[qi] = ok
		} else {
			// Overlapping: analyze left-only, right-only, and shared regions
			lBoth := max2(l1, l2)
			rBoth := min2(r1, r2)
			leftL, leftR := l1, l2-1
			rightL, rightR := r2+1, r1

			cntAleft := cntRange(prefA, leftL, leftR)
			cntBleft := cntRange(prefB, leftL, leftR)
			cntAright := cntRange(prefA, rightL, rightR)
			cntBright := cntRange(prefB, rightL, rightR)
			cntBothA := cntRange(prefA, lBoth, rBoth)
			cntBothB := cntRange(prefB, lBoth, rBoth)

			ok := true
			// Left-only region: A's exclusive chars must match B's exclusive chars
			for c := 0; c < 26; c++ {
				if cntAleft[c] != cntBleft[c] || cntAright[c] != cntBright[c] {
					ok = false
					break
				}
			}
			if ok {
				// Shared region: must also be fixable
				for c := 0; c < 26; c++ {
					if cntBothA[c] != cntBothB[c] {
						ok = false
						break
					}
				}
			}
			ans[qi] = ok
		}
	}
	return ans
}

func reverseStr(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example
	res := canMakePalindromeQueries("abcabc", [][]int{{1, 1, 3, 5}, {0, 2, 5, 5}})
	fmt.Println(res)

	// Simple
	res2 := canMakePalindromeQueries("ab", [][]int{{0, 0, 1, 1}})
	fmt.Println(res2)

	// More tests
	res3 := canMakePalindromeQueries("abba", [][]int{{0, 0, 2, 3}})
	fmt.Println(res3)

	res4 := canMakePalindromeQueries("abcddcba", [][]int{{0, 0, 1, 1}})
	fmt.Println(res4)
}
```
