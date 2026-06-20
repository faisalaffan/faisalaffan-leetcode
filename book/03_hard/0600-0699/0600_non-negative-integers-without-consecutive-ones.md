# 0600 — Non Negative Integers Without Consecutive Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func findIntegers(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #600: Non-negative Integers without Consecutive Ones
// https://leetcode.com/problems/non-negative-integers-without-consecutive-ones/
// Difficulty: Hard

import "fmt"

func main() {
	// Test cases
	testCases := []struct {
		input int
		want  int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 5},
		{7, 5},
		{8, 6},
		{9, 7},
		{10, 8},
		{100, 34},
		{1000, 144},
	}

	for _, tc := range testCases {
		got := findIntegers(tc.input)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: findIntegers(%d) = %d (want %d)\n", status, tc.input, got, tc.want)
	}
}

func findIntegers(n int) int {
	// DP approach: count numbers with consecutive ones constraint
	// Convert n to binary representation
	binary := fmt.Sprintf("%b", n)

	// fib[i] = number of valid numbers with i bits (no consecutive ones)
  // Alokasi slice integer
	fib := make([]int, len(binary)+2)
	fib[0] = 1 // 0 bits
	fib[1] = 2 // 1 bit: 0, 1
	for i := 2; i <= len(binary)+1; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	ans := 0
	prevBit := false // whether previous bit was 1

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			// If we set this bit to 0, remaining bits can be anything valid
			remaining := len(binary) - i - 1
			ans += fib[remaining]

			// If previous bit was also 1, we have consecutive ones, stop
			if prevBit {
				return ans
			}
			prevBit = true
		} else {
			prevBit = false
		}
	}

	// Count n itself (had no consecutive ones)
	return ans + 1
}
```
