# 0012 — Integer To Roman

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func intToRoman(num int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #12: Integer to Roman
// https://leetcode.com/problems/integer-to-roman/
// Difficulty: Medium

import "fmt"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(intToRoman(3749)) // "MMMDCCXLIX"

	// Test case 2
	fmt.Println(intToRoman(58)) // "LVIII"

	// Test case 3
	fmt.Println(intToRoman(1994)) // "MCMXCIV"
}

// Time: O(1) | Space: O(1)
```
