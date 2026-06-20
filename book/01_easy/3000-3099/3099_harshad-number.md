# 3099 — Harshad Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func HarshadNumber(x int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3099: Harshad Number
// https://leetcode.com/problems/harshad-number/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumOfTheDigitsOfHarshadNumber
	fmt.Println(HarshadNumber(18)) // 9
	fmt.Println(HarshadNumber(23)) // -1
}

// Time: O(log n) | Space: O(1)
// LeetCode submission name: sumOfTheDigitsOfHarshadNumber
func HarshadNumber(x int) int {
	sum := 0
	n := x
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
```
