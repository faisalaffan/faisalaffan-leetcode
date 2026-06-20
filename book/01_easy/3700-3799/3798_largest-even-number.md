# 3798 — Largest Even Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestEvenNumber(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3798: Largest Even Number
// https://leetcode.com/problems/largest-even-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LargestEvenNumber("1112"))
	fmt.Println(LargestEvenNumber("221"))
	fmt.Println(LargestEvenNumber("1"))
}

// Time: O(n)
// Space: O(n)
func LargestEvenNumber(s string) string {
	i := len(s)
	for i > 0 && s[i-1] == '1' {
		i--
	}
	return s[:i]
}
```
