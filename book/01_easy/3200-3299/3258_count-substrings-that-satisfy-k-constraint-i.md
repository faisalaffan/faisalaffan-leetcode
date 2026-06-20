# 3258 — Count Substrings That Satisfy K Constraint I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountSubstringsThatSatisfyKConstraintI(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3258: Count Substrings That Satisfy K-Constraint I
// https://leetcode.com/problems/count-substrings-that-satisfy-k-constraint-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("10101", 1))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("1010101", 2))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("11111", 1))
}

// CountSubstringsThatSatisfyKConstraintI counts substrings where both number of 0s and 1s <= k.
// Time: O(n^2). Space: O(1).
func CountSubstringsThatSatisfyKConstraintI(s string, k int) int {
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		zeros, ones := 0, 0
		for j := i; j < n; j++ {
			if s[j] == '0' {
				zeros++
			} else {
				ones++
			}
			if zeros <= k || ones <= k {
				count++
			} else {
				break
			}
		}
	}
	return count
}
```
