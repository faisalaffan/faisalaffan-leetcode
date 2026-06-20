# 2851 — String Transformation

## Deskripsi

**Soal:** [2851. String Transformation](https://leetcode.com/problems/string-transformation/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), KMP (Knuth-Morris-Pratt, pencocokan string)

**Fungsi Solusi:** `func numberOfWays(s, t string, k int64) int`

## Solusi Go

```go
package main

// LeetCode #2851: String Transformation
// https://leetcode.com/problems/string-transformation/
// Difficulty: Hard
//
// Given strings s and t of equal length n. Operation: remove a suffix of length
// l (0 < l < n) and prepend it. Count ways to transform s into t in exactly k
// operations, mod 1e9+7.
//
// Reduction: each operation is a rotation. Let g = number of rotations of s
// that equal t (found via KMP in O(n)). Two-state DP: good (==t) / bad (!=t).
// Transition matrix exponentiation in O(log k). Total: O(n + log k).

import "fmt"

const mod2851 = 1000000007

func numberOfWays(s, t string, k int64) int {
	n := len(s)

	// --- KMP: find all rotations of s that equal t ---
	// Pattern = t, Text = s+s (without last char to avoid full wrap)
	pattern := t
	text := s + s[:n-1]

	// Build LPS array for pattern
  // Membuat slice untuk menyimpan hasil
	lps := make([]int, n)
	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		lps[i] = j
	}

	g := 0
	j := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pattern[j] {
			j = lps[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == n {
			g++
			j = lps[j-1]
		}
	}

	if k == 0 {
		if s == t {
			return 1
		}
		return 0
	}

	// --- Matrix exponentiation ---
	// States: 0 = good (equals t), 1 = bad
	// M[0][0] = g-1 (good -> good)
	// M[0][1] = g   (bad -> good)
	// M[1][0] = n-g  (good -> bad)
	// M[1][1] = n-1-g (bad -> bad)
	m := [2][2]int64{
		{int64((g - 1 + mod2851) % mod2851), int64(g % mod2851)},
		{int64((n - g + mod2851) % mod2851), int64((n - 1 - g + mod2851) % mod2851)},
	}

	mk := matPow(m, k)

	startGood := int64(0)
	if s == t {
		startGood = 1
	}
	startBad := int64(1) - startGood

	// result = mk[0][0] * startGood + mk[0][1] * startBad
	result := (mk[0][0]*startGood + mk[0][1]*startBad) % mod2851
	return int(result)
}

func matMul(a, b [2][2]int64) [2][2]int64 {
	return [2][2]int64{
		{(a[0][0]*b[0][0] + a[0][1]*b[1][0]) % mod2851,
			(a[0][0]*b[0][1] + a[0][1]*b[1][1]) % mod2851},
		{(a[1][0]*b[0][0] + a[1][1]*b[1][0]) % mod2851,
			(a[1][0]*b[0][1] + a[1][1]*b[1][1]) % mod2851},
	}
}

func matPow(m [2][2]int64, k int64) [2][2]int64 {
	res := [2][2]int64{{1, 0}, {0, 1}}
	for k > 0 {
		if k&1 == 1 {
			res = matMul(res, m)
		}
		m = matMul(m, m)
		k >>= 1
	}
	return res
}

func main() {
	// Example: "abcd" -> "cdab" in k=2 steps
	// Rotations of "abcd" equal to "cdab": 1 (rotation by 2)
	// n=4, g=1. M = [[0,1],[3,2]]. dp_0 = [0,1].
	// After 2 steps: should get 2
	fmt.Println(numberOfWays("abcd", "cdab", 2))

	// s == t, k=0
	fmt.Println(numberOfWays("abc", "abc", 0))

	// s != t, k=0
	fmt.Println(numberOfWays("abc", "cab", 0))

	// s == t, k=1: from "abc" to "abc" in 1 step
	// Only "abc" equals "abc", so g=1.
	// n=3, g=1. M = [[0,1],[2,1]]. dp_0 = [1,0].
	// After 1 step: dp_1[good] = 0*1 + 1*0 = 0? No...
	// Wait, from "abc", can we reach "abc" in 1 step?
	// l=1: "cab", l=2: "bca". Neither is "abc". So 0 ways. ✓
	fmt.Println(numberOfWays("abc", "abc", 1))

	// All chars same: s="aaa", t="aaa"
	// All 3 rotations match, so g=3. M = [[2,3],[0,0]].
	// dp_0 = [1,0].
	// k=1: from "aaa", any operation (l=1 or l=2) gives "aaa". So 2 ways.
	fmt.Println(numberOfWays("aaa", "aaa", 1))

	// k=2: from "aaa", 2*2=4 ways.
	fmt.Println(numberOfWays("aaa", "aaa", 2))

	// Larger k test
	fmt.Println(numberOfWays("ab", "ba", 3))
}
```
