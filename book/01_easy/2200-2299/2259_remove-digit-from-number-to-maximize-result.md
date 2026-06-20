# 2259 — Remove Digit From Number To Maximize Result

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func RemoveDigitFromNumberToMaximizeResult(number string, digit byte) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2259: Remove Digit From Number to Maximize Result
// https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("123", '3'))  // "12"
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("1231", '1')) // "231"
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("551", '5'))  // "51"
}

func RemoveDigitFromNumberToMaximizeResult(number string, digit byte) string {
	best := ""
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			candidate := number[:i] + number[i+1:]
			if candidate > best {
				best = candidate
			}
		}
	}
	return best
}
```
