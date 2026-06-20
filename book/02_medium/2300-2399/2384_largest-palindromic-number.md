# 2384 — Largest Palindromic Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func largestPalindromic(num string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2384: Largest Palindromic Number
// https://leetcode.com/problems/largest-palindromic-number/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Count digits. Build palindrome: place pairs from 9 down to 0, then single middle digit.

import "fmt"

func main() {
	fmt.Println(largestPalindromic("444947137")) // "7449447"
	fmt.Println(largestPalindromic("00009"))     // "9"
	fmt.Println(largestPalindromic("0000"))      // "0"
}

func largestPalindromic(num string) string {
  // Alokasi slice integer
	cnt := make([]int, 10)
	for _, ch := range num {
		cnt[ch-'0']++
	}

	left := make([]byte, 0)
	middle := ""

	for d := 9; d >= 0; d-- {
		pairs := cnt[d] / 2
		if d == 0 && len(left) == 0 {
			// skip leading zeros
			if middle == "" && cnt[0]%2 == 1 {
				middle = "0"
			}
			break
		}
		for i := 0; i < pairs; i++ {
			left = append(left, byte('0'+d))
		}
		if middle == "" && cnt[d]%2 == 1 {
			middle = string(rune('0' + d))
		}
	}

	if len(left) == 0 && middle == "" {
		return "0"
	}

	// mirror: left + middle + reverse(left)
	right := make([]byte, len(left))
	for i, b := range left {
		right[len(left)-1-i] = b
	}
	return string(left) + middle + string(right)
}
```
