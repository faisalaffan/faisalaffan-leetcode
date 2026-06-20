# 3272 — Find The Count Of Good Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countGoodIntegers(n int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3272: Find the Count of Good Integers
// https://leetcode.com/problems/find-the-count-of-good-integers/
// Difficulty: Hard
//
// Combinatorics approach:
// 1. Generate all n-digit palindromes by iterating through left halves.
// 2. For each palindrome divisible by k, record its sorted digit multiset.
// 3. For each unique digit multiset, count the number of n-digit permutations
//    (no leading zero) using the formula:
//      (n - freq[0]) * (n-1)! / prod(freq[d]! for d in 0..9)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: n=2,k=2 => 4
	fmt.Println(countGoodIntegers(2, 2))
	// Example 2: n=3,k=5 => 27
	fmt.Println(countGoodIntegers(3, 5))
	// Example 3: n=1,k=1 => 9
	fmt.Println(countGoodIntegers(1, 1))
	// Example 4: n=4,k=7 => 189
	fmt.Println(countGoodIntegers(4, 7))
	// Example 5: n=5,k=6 => 1755
	fmt.Println(countGoodIntegers(5, 6))
}

func countGoodIntegers(n int, k int) int64 {
	// Precompute factorials
  // Alokasi slice
	fact := make([]int64, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i)
	}

  // HashMap: O(1) lookup
	seen := make(map[string]bool)
	var ans int64

	// Start of the left half (1..9 for the first digit)
	start := 1
	for i := 1; i < (n+1)/2; i++ {
		start *= 10
	}
	end := start * 10

	for half := start; half < end; half++ {
		// Build the full palindrome string
		s := fmt.Sprintf("%d", half)
		// Mirror the left half (drop last char for odd length)
		rev := reverse(s)
		if n%2 == 1 {
			rev = rev[1:] // drop the middle character
		}
		palStr := s + rev

		// Check divisibility
		if !divisibleBy(palStr, k) {
			continue
		}

		// Sort digits to get canonical form
		digits := []byte(palStr)
  // Custom sort
		sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
		key := string(digits)

		if seen[key] {
			continue
		}
		seen[key] = true

		// Count frequencies
  // Alokasi slice
		freq := make([]int, 10)
		for _, ch := range palStr {
			freq[ch-'0']++
		}

		// Count permutations without leading zeros
		// Formula: (n - freq[0]) * (n-1)! / prod(freq[d]!)
		perm := int64(n-freq[0]) * fact[n-1]
		for _, f := range freq {
			perm /= fact[f]
		}
		ans += perm
	}

	return ans
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func divisibleBy(s string, k int) bool {
	var rem int
	for _, ch := range s {
		rem = (rem*10 + int(ch-'0')) % k
	}
	return rem == 0
}
```
