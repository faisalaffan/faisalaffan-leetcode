# 1256 — Encode Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func encode(num int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Fenwick Tree (BIT)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Fenwick Tree (BIT)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1256: Encode Number
// https://leetcode.com/problems/encode-number/
// Difficulty: Medium [Paid]

// Encode n as binary of n+1, then remove first bit.
// n=0 -> "0" (binary of 1 -> "1", remove first -> "")
// Wait: n=0 -> "". Let me check the pattern.
// 0: "" (1->"1", drop first->"")
// 1: "0" (2->"10", drop first->"0")
// 2: "1" (3->"11", drop first->"1")
// 3: "00" (4->"100", drop first->"00")
// 4: "01" (5->"101", drop first->"01")

// Time: O(log n)
// Space: O(log n)

func encode(num int) string {
	if num == 0 {
		return ""
	}

	// num+1 in binary, then drop first bit
	n := num + 1
	result := ""

	for n > 1 {
		if n%2 == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}
		n /= 2
	}

	return result
}

func main() {
	fmt.Printf("%q (expected: %q)\n", encode(0), "")
	fmt.Printf("%q (expected: %q)\n", encode(1), "0")
	fmt.Printf("%q (expected: %q)\n", encode(2), "1")
	fmt.Printf("%q (expected: %q)\n", encode(3), "00")
}
```
