# 0537 — Complex Number Multiplication

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func ComplexNumberMultiplication(num1 string, num2 string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #537: Complex Number Multiplication
// https://leetcode.com/problems/complex-number-multiplication/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ComplexNumberMultiplication("1+1i", "1+1i"))
	fmt.Println(ComplexNumberMultiplication("1+-1i", "1+-1i"))
}

func ComplexNumberMultiplication(num1 string, num2 string) string {
	a, b := parseComplex(num1)
	c, d := parseComplex(num2)

	real := a*c - b*d
	imag := a*d + b*c

	return fmt.Sprintf("%d+%di", real, imag)
}

func parseComplex(s string) (int, int) {
	parts := strings.Split(s, "+")
	real, _ := strconv.Atoi(parts[0])
	imagPart := parts[1][:len(parts[1])-1] // remove trailing 'i'
	imag, _ := strconv.Atoi(imagPart)
	return real, imag
}
```
