# 0013 — Roman To Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func RomanToInt(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #13: Roman to Integer
// https://leetcode.com/problems/roman-to-integer/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RomanToInt(s string) int {
	vals := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	sum, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := vals[s[i]]
		if cur < prev {
			sum -= cur
		} else {
			sum += cur
		}
		prev = cur
	}
	return sum
}

func main() {
	fmt.Println(RomanToInt("III"))
	fmt.Println(RomanToInt("LVIII"))
	fmt.Println(RomanToInt("MCMXCIV"))
}
```
