# 0476 — Number Complement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberComplement(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #476: Number Complement
// https://leetcode.com/problems/number-complement/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func NumberComplement(num int) int {
	mask := ^0
	for num&mask != 0 {
		mask <<= 1
	}
	return ^num & ^mask
}

func main() {
	fmt.Println(NumberComplement(5))
	fmt.Println(NumberComplement(1))
	fmt.Println(NumberComplement(2))
}
```
