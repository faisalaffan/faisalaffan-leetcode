# 3483 — Unique 3 Digit Even Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func UniqueThreeDigitEvenNumbers(digits []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^3). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3483: Unique 3-Digit Even Numbers
// https://leetcode.com/problems/unique-3-digit-even-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(UniqueThreeDigitEvenNumbers([]int{1, 2, 3, 4}))
	fmt.Println(UniqueThreeDigitEvenNumbers([]int{0, 2, 2}))
}

// UniqueThreeDigitEvenNumbers counts unique 3-digit even numbers that can be formed from digits (no leading zero).
// Time: O(n^3). Space: O(n).
func UniqueThreeDigitEvenNumbers(digits []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	used := make(map[int]bool)
	n := len(digits)
	for i := 0; i < n; i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				num := digits[i]*100 + digits[j]*10 + digits[k]
				if num%2 == 0 {
					used[num] = true
				}
			}
		}
	}
	return len(used)
}
```
