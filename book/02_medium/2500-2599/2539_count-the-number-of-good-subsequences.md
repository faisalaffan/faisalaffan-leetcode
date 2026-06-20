# 2539 — Count The Number Of Good Subsequences

## Deskripsi

**Soal:** [2539. Count The Number Of Good Subsequences](https://leetcode.com/problems/count-the-number-of-good-subsequences/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func countGoodSubsequences(s string) int`

## Solusi Go

```go
package main

// LeetCode #2539: Count the Number of Good Subsequences
// https://leetcode.com/problems/count-the-number-of-good-subsequences/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func countGoodSubsequences(s string) int {
	const mod = 1_000_000_007
  // Membuat map untuk pencarian O(1): key → value
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
  // Membuat slice untuk menyimpan hasil
	fact := make([]int, maxN+1)
	fact[0] = 1
	for i := 1; i <= maxN; i++ {
		fact[i] = fact[i-1] * i % mod
	}

  // Membuat slice untuk menyimpan hasil
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
