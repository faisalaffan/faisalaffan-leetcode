# 2949 — Count Beautiful Substrings Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isVowel(ch byte) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2949: Count Beautiful Substrings II
// https://leetcode.com/problems/count-beautiful-substrings-ii/
// Difficulty: Hard
//
// A substring is beautiful if:
//   1. vowels == consonants (balanced)
//   2. (vowels * consonants) % k == 0, i.e., v^2 % k == 0 (since v == c)
//
// For condition 2: if v^2 % k == 0, then v % p == 0 where p is derived
// from the prime factorization of k: for each prime factor q^e of k,
// p gets q^ceil(e/2).
//
// Track (prefix_diff, index_mod_p) as state in hash map.
// Each matching pair forms a beautiful substring.

import "fmt"

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func beautifulSubstrings(s string, k int) int64 {
	// Compute smallest p such that p^2 % (4k) == 0
	// This is equivalent to: for each prime factor q^e of k*4, p gets q^ceil(e/2)
	k4 := k * 4
	p := 1
	for i := 2; i*i <= k4; i++ {
		if k4%i == 0 {
			cnt := 0
			for k4%i == 0 {
				k4 /= i
				cnt++
			}
			for j := 0; j < (cnt+1)/2; j++ {
				p *= i
			}
		}
	}
	if k4 > 1 {
		p *= k4
	}

	type state struct {
		diff int
		mod  int
	}
  // HashMap: O(1) lookup
	counts := make(map[state]int64)
	counts[state{diff: 0, mod: 0}] = 1

	var ans int64
	diff := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			diff++
		} else {
			diff--
		}
		st := state{diff: diff, mod: (i + 1) % p}
		ans += counts[st]
		counts[st]++
	}
	return ans
}

func main() {
	// Example: s="baeyh", k=2 -> 5
	fmt.Println(beautifulSubstrings("baeyh", 2))

	// No vowels
	fmt.Println(beautifulSubstrings("bcdf", 1))

	// Simple
	fmt.Println(beautifulSubstrings("ab", 1))

	// All same
	fmt.Println(beautifulSubstrings("aaabbb", 1))
	fmt.Println(beautifulSubstrings("leetcode", 2))
}
```
