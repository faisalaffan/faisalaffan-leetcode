# 1960 — Maximum Product Of The Length Of Two Palindromic Substrings

## Deskripsi

**Soal:** [1960. Maximum Product Of The Length Of Two Palindromic Substrings](https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-substrings/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func maxProduct(s string) int64`

## Solusi Go

```go
package main

// LeetCode #1960: Maximum Product of the Length of Two Palindromic Substrings
// https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-substrings/
// Difficulty: Hard
// Manacher to find all palindrome radii, then compute L[i] = longest palindrome
// ending at i and R[i] = longest palindrome starting at i. Answer = max L[i]*R[i+1].

import "fmt"

func maxProduct(s string) int64 {
	n := len(s)
	// Odd palindrome radii (center is a character)
  // Membuat slice untuk menyimpan hasil
	odd := make([]int, n)
	center, right := 0, 0
	for i := 0; i < n; i++ {
		if i < right {
			odd[i] = min(odd[2*center-i], right-i)
		}
		for i-odd[i] >= 0 && i+odd[i] < n && s[i-odd[i]] == s[i+odd[i]] {
			odd[i]++
		}
		if i+odd[i] > right {
			center, right = i, i+odd[i]
		}
	}

	// Even palindrome radii (centered between s[i-1] and s[i])
	// even[i] = radius for even palindrome centered between i-1 and i
  // Membuat slice untuk menyimpan hasil
	even := make([]int, n+1)
	l, r := 0, 0
	for i := 0; i <= n; i++ {
		if i < r {
			even[i] = min(even[l+r-i], r-i)
		}
		for i-even[i]-1 >= 0 && i+even[i] < n && s[i-even[i]-1] == s[i+even[i]] {
			even[i]++
		}
		if i+even[i] > r {
			l, r = i-even[i], i+even[i]
		}
	}

	// L[i] = longest palindrome ending at i
  // Membuat slice untuk menyimpan hasil
	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = 1
	}
	// Odd palindromes
	for c := 0; c < n; c++ {
		rad := odd[c]
		end := c + rad - 1
		length := 2*rad - 1
		if end < n && length > L[end] {
			L[end] = length
		}
	}
	// Even palindromes
	for i := 0; i <= n; i++ {
		rad := even[i]
		if rad == 0 {
			continue
		}
		end := i + rad - 1
		length := 2 * rad
		if end < n && length > L[end] {
			L[end] = length
		}
	}
	// Propagate: if palindrome of length k ends at i+1, one of length k-2 ends at i
	for i := n - 2; i >= 0; i-- {
		if L[i+1]-2 > L[i] {
			L[i] = L[i+1] - 2
		}
	}

	// R[i] = longest palindrome starting at i
  // Membuat slice untuk menyimpan hasil
	R := make([]int, n)
	for i := 0; i < n; i++ {
		R[i] = 1
	}
	// Odd palindromes
	for c := 0; c < n; c++ {
		rad := odd[c]
		start := c - rad + 1
		length := 2*rad - 1
		if start >= 0 && length > R[start] {
			R[start] = length
		}
	}
	// Even palindromes
	for i := 0; i <= n; i++ {
		rad := even[i]
		if rad == 0 {
			continue
		}
		start := i - rad
		length := 2 * rad
		if start >= 0 && length > R[start] {
			R[start] = length
		}
	}
	// Propagate forward
	for i := 1; i < n; i++ {
		if R[i-1]-2 > R[i] {
			R[i] = R[i-1] - 2
		}
	}

	// Suffix max of R so we can check any split, not just adjacent positions
  // Membuat slice untuk menyimpan hasil
	suffixMaxR := make([]int, n)
	suffixMaxR[n-1] = R[n-1]
	for i := n - 2; i >= 0; i-- {
		if R[i] > suffixMaxR[i+1] {
			suffixMaxR[i] = R[i]
		} else {
			suffixMaxR[i] = suffixMaxR[i+1]
		}
	}

	var ans int64 = 0
	for i := 0; i < n-1; i++ {
		prod := int64(L[i]) * int64(suffixMaxR[i+1])
		if prod > ans {
			ans = prod
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProduct("ababbb"))   // Expected: 9
	fmt.Println(maxProduct("zaaaxbbby")) // Expected: 9 (zaaaz=5, bbby=4 or similar)
	fmt.Println(maxProduct("a"))         // Expected: 0 (only one char, no two substrings)
}
```
