# 1067 — Digit Count In Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func digitCountInRange(d int, low int, high int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1067: Digit Count in Range
// https://leetcode.com/problems/digit-count-in-range/
// Difficulty: Hard [Paid]
//
// Count digit occurrences using mathematical decomposition per position
// (ones, tens, hundreds, ...). For a number n, countDigit(d, n) returns
// occurrences of digit d in [0, n]. Result for range [low, high] is
// countDigit(d, high) - countDigit(d, low-1).

import "fmt"

func main() {
	fmt.Println(digitCountInRange(1, 1, 13))
}

func digitCountInRange(d int, low int, high int) int {
	return countDigits(d, high) - countDigits(d, low-1)
}

func countDigits(d int, n int) int {
	if n < 0 {
		return 0
	}
	count := 0
	for pos := 1; pos <= n; pos *= 10 {
		left := n / (pos * 10)
		cur := (n / pos) % 10
		right := n % pos

		if d != 0 {
			if cur > d {
				count += (left + 1) * pos
			} else if cur == d {
				count += left*pos + right + 1
			} else {
				count += left * pos
			}
		} else {
			// digit 0: skip leading zeros
			if left > 0 {
				if cur > 0 {
					count += left * pos
				} else {
					count += (left-1)*pos + right + 1
				}
			}
		}
	}
	return count
}
```
