# 2002 — Maximum Product Of The Length Of Two Palindromic Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaximumProductOfTheLengthOfTwoPalindromicSubsequences(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Bitmask

**Waktu:** O(3^n), Space: O(2^n)  |  **Ruang:** O(2^n)

> 🎓 **Fresh Grad Tips:** Kuasai **Bitmask** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2002: Maximum Product of the Length of Two Palindromic Subsequences
// https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/
// Difficulty: Medium

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(MaximumProductOfTheLengthOfTwoPalindromicSubsequences("leetcodecom"))
	fmt.Println(MaximumProductOfTheLengthOfTwoPalindromicSubsequences("bb"))
	fmt.Println(MaximumProductOfTheLengthOfTwoPalindromicSubsequences("accbcaxxcxx"))
}

// Time: O(3^n), Space: O(2^n)
func MaximumProductOfTheLengthOfTwoPalindromicSubsequences(s string) int {
	n := len(s)
	p := make([]bool, 1<<n)

	for mask := 1; mask < 1<<n; mask++ {
		p[mask] = true
		i, j := 0, n-1
		for i < j {
			for i < j && (mask>>i&1) == 0 {
				i++
			}
			for i < j && (mask>>j&1) == 0 {
				j--
			}
			if i < j && s[i] != s[j] {
				p[mask] = false
				break
			}
			i++
			j--
		}
	}

	ans := 0
	all := 1<<n - 1
	for mask := 1; mask < 1<<n; mask++ {
		if !p[mask] {
			continue
		}
		a := bits.OnesCount(uint(mask))
		rest := all ^ mask
		for sub := rest; sub > 0; sub = (sub - 1) & rest {
			if p[sub] {
				b := bits.OnesCount(uint(sub))
				if a*b > ans {
					ans = a * b
				}
			}
		}
	}

	return ans
}
```
