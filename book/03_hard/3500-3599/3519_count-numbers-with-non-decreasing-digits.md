# 3519 — Count Numbers With Non Decreasing Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countNumbers(l string, r string, b int) int
```

> **💡 Hint:** Digit DP. For each position, track the last digit used to ensure

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3519: Count Numbers with Non-Decreasing Digits
// https://leetcode.com/problems/count-numbers-with-non-decreasing-digits/
// Difficulty: Hard
//
// Count numbers in range [l, r] (inclusive) whose digits are non-decreasing
// when represented in base b.
//
// Approach: Digit DP. For each position, track the last digit used to ensure
// non-decreasing property.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countNumbers("10", "20", 10))
	// Example 2
	fmt.Println(countNumbers("1", "9", 10))
	// Example 3: base 2
	fmt.Println(countNumbers("1", "100", 2))
	// Edge: single digit range
	fmt.Println(countNumbers("5", "5", 10))
	// Edge: large base
	fmt.Println(countNumbers("0", "FF", 16))
}

const MOD = 1000000007

func countNumbers(l string, r string, b int) int {
	// Count numbers <= r with non-decreasing digits, subtract those < l
	// plus one if l itself qualifies
	cntR := countUpTo(r, b)
	cntL := countUpTo(l, b)
	// Add 1 if l qualifies
	if hasNonDecreasingDigits(l, b) {
		cntR = (cntR - cntL + 1 + MOD) % MOD
	} else {
		cntR = (cntR - cntL + MOD) % MOD
	}
	return cntR
}

func countUpTo(s string, b int) int {
	// Convert s to digits in base b
	digits := toDigits(s, b)
	n := len(digits)

	// DP[pos][tight][lastDigit][started]
	var memo [101][2][17][2]int
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = [2]int{-1, -1}
			}
		}
	}

	var dp func(pos int, tight bool, last int, started bool) int
	dp = func(pos int, tight bool, last int, started bool) int {
		if pos == n {
			if started {
				return 1
			}
			return 0
		}
		t := 0
		if tight {
			t = 1
		}
		s := 0
		if started {
			s = 1
		}
		if memo[pos][t][last][s] != -1 {
			return memo[pos][t][last][s]
		}

		limit := b - 1
		if tight {
			limit = digits[pos]
		}

		var total int64
		for d := 0; d <= limit; d++ {
			nextTight := tight && (d == limit)
			if !started {
				if d == 0 {
					total = (total + int64(dp(pos+1, nextTight, 0, false))) % MOD
				} else {
					total = (total + int64(dp(pos+1, nextTight, d, true))) % MOD
				}
			} else if d >= last {
				total = (total + int64(dp(pos+1, nextTight, d, true))) % MOD
			}
		}

		memo[pos][t][last][s] = int(total)
		return memo[pos][t][last][s]
	}

	return dp(0, true, 0, false)
}

func toDigits(s string, b int) []int {
	// Convert string representation (any base) to digits in base b
	// For base up to 16, interpret as number
	if b <= 10 {
		return toDigitsDecimal(s)
	}
	// For base > 10, parse from string
	num := 0
	for _, ch := range s {
		var d int
		if ch >= '0' && ch <= '9' {
			d = int(ch - '0')
		} else if ch >= 'A' && ch <= 'F' {
			d = int(ch-'A') + 10
		} else if ch >= 'a' && ch <= 'f' {
			d = int(ch-'a') + 10
		}
		num = num*b + d
	}
	// Convert back to digits
	var result []int
	if num == 0 {
		return []int{0}
	}
	for num > 0 {
		result = append([]int{num % b}, result...)
		num /= b
	}
	return result
}

func toDigitsDecimal(s string) []int {
  // Alokasi slice integer
	result := make([]int, len(s))
	for i, ch := range s {
		result[i] = int(ch - '0')
	}
	return result
}

func hasNonDecreasingDigits(s string, b int) bool {
	digits := toDigits(s, b)
	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			return false
		}
	}
	return true
}
```
