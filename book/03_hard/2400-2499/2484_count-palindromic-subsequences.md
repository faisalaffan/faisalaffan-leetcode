# 2484 — Count Palindromic Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPalindromicSubsequences(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2484: Count Palindromic Subsequences
// https://leetcode.com/problems/count-palindromic-subsequences/
// Difficulty: Hard
//
// Count distinct palindromic subsequences of length 5 (a b c b a).
// For each middle position j and each pair (a,b):
//   leftCount[a][b] = number of (a,b) ordered pairs in s[0:j] (exclusive)
//   rightCount[a][b] = number of (a,b) ordered pairs in s[j+1:n] (exclusive)
//   result += leftCount[a][b] * rightCount[b][a]

import "fmt"

const MOD = 1000000007

func main() {
	// Example 1: "103301" => 2
	fmt.Println(countPalindromicSubsequences("103301"))
	// Example 2: "0000000" => 21
	fmt.Println(countPalindromicSubsequences("0000000"))
	// Example 3: "9999900000" => 96
	fmt.Println(countPalindromicSubsequences("9999900000"))
	// Edge: length 5 palindrome
	fmt.Println(countPalindromicSubsequences("12321"))
	// Edge: all distinct
	fmt.Println(countPalindromicSubsequences("12345"))
}

func countPalindromicSubsequences(s string) int {
	n := len(s)
	if n < 5 {
		return 0
	}

  // Alokasi slice integer
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i] - '0')
	}

	// prefix counts: prefixCnt[i][d] = count of digit d in s[0:i]
  // Alokasi slice integer
	prefixCnt := make([][10]int, n+1)
	for i := 0; i < n; i++ {
		for d := 0; d < 10; d++ {
			prefixCnt[i+1][d] = prefixCnt[i][d]
		}
		prefixCnt[i+1][nums[i]]++
	}

	// suffix counts: suffixCnt[i][d] = count of digit d in s[i:n]
  // Alokasi slice integer
	suffixCnt := make([][10]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for d := 0; d < 10; d++ {
			suffixCnt[i][d] = suffixCnt[i+1][d]
		}
		suffixCnt[i][nums[i]]++
	}

	// Compute total ordered pairs in the full string
	// rightPairs[a][b] = number of (a,b) pairs where a appears before b
	rightPairs := [10][10]int64{}
	for i := 0; i < n; i++ {
		d := nums[i]
		for a := 0; a < 10; a++ {
			rightPairs[a][d] += int64(prefixCnt[i][a])
		}
	}

	leftPairs := [10][10]int64{}
	var result int64 = 0

	for j := 0; j < n; j++ {
		d := nums[j]

		// Remove pairs where position j is the FIRST element (d at j, any after j)
		// These are no longer after position j.
		for b := 0; b < 10; b++ {
			rightPairs[d][b] -= int64(suffixCnt[j+1][b])
		}

		// For each (a,b), count palindromes with j as middle
		for a := 0; a < 10; a++ {
			for b := 0; b < 10; b++ {
				result = (result + leftPairs[a][b]*rightPairs[b][a]) % MOD
			}
		}

		// Add position j's contribution to leftPairs for NEXT iteration.
		// Pairs (a at p < j, d at j) now become part of the left prefix.
		for a := 0; a < 10; a++ {
			leftPairs[a][d] += int64(prefixCnt[j][a])
		}
	}

	return int(result)
}
```
