# 3931 — Check Adjacent Digit Differences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckAdjacentDigitDifferences(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3931: Check Adjacent Digit Differences
// https://leetcode.com/problems/check-adjacent-digit-differences/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckAdjacentDigitDifferences("132"))
	fmt.Println(CheckAdjacentDigitDifferences("129"))
}

// Time: O(n)
// Space: O(1)
func CheckAdjacentDigitDifferences(s string) bool {
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s)-1; i++ {
		diff := int(s[i]) - int(s[i+1])
		if diff < 0 {
			diff = -diff
		}
		if diff > 2 {
			return false
		}
	}
	return true
}
```
