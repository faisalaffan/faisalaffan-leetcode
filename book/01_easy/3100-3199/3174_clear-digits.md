# 3174 — Clear Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func ClearDigits(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3174: Clear Digits
// https://leetcode.com/problems/clear-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ClearDigits("abc"))
	fmt.Println(ClearDigits("cb34"))
	fmt.Println(ClearDigits("a1b2c3"))
}

// ClearDigits removes all digits and their nearest non-digit character to the left.
// Time: O(n). Space: O(n).
func ClearDigits(s string) string {
	stack := make([]byte, 0, len(s))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
```
