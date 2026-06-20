# 2520 — Count The Digits That Divide A Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountTheDigitsThatDivideANumber(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2520: Count the Digits That Divide a Number
// https://leetcode.com/problems/count-the-digits-that-divide-a-number/
// Difficulty: Easy
// Time O(log n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountTheDigitsThatDivideANumber(7))    // 1
	fmt.Println(CountTheDigitsThatDivideANumber(121))  // 2
	fmt.Println(CountTheDigitsThatDivideANumber(1248)) // 4
}

func CountTheDigitsThatDivideANumber(num int) int {
	count := 0
	n := num
	for n > 0 {
		digit := n % 10
		if digit != 0 && num%digit == 0 {
			count++
		}
		n /= 10
	}
	return count
}
```
