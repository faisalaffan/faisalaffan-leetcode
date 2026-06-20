# 2539 — Count The Number Of Good Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countGoodSubsequences(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2539: Count the Number of Good Subsequences
// https://leetcode.com/problems/count-the-number-of-good-subsequences/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func countGoodSubsequences(s string) int {
	const mod = 1_000_000_007
  // HashMap: O(1) lookup
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	// Precompute factorials and inverse factorials
	maxN := maxFreq
  // Alokasi slice
	fact := make([]int, maxN+1)
	fact[0] = 1
	for i := 1; i <= maxN; i++ {
		fact[i] = fact[i-1] * i % mod
	}

  // Alokasi slice
	invFact := make([]int, maxN+1)
	invFact[maxN] = powMod(fact[maxN], mod-2, mod)
	for i := maxN; i > 0; i-- {
		invFact[i-1] = invFact[i] * i % mod
	}

	nCr := func(n, r int) int {
		if r < 0 || r > n {
			return 0
		}
		return fact[n] * invFact[r] % mod * invFact[n-r] % mod
	}

	var ans int
	for maxLen := 1; maxLen <= maxFreq; maxLen++ {
		ways := 1
		for _, f := range freq {
			if f >= maxLen {
				ways = ways * (nCr(f, maxLen) + 1) % mod
			}
		}
		ans = (ans + ways - 1 + mod) % mod
	}

	return ans
}

func powMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countGoodSubsequences("aabb"))
	// Expected: 11

	// Test case 2
	fmt.Println("Test 2:", countGoodSubsequences("leet"))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", countGoodSubsequences("abcd"))
	// Expected: 4
}
```
