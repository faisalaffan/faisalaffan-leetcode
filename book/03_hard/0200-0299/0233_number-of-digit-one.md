# 0233 — Number Of Digit One

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countDigitOne(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #233: Number of Digit One
// https://leetcode.com/problems/number-of-digit-one/
// Difficulty: Hard

import "fmt"

func countDigitOne(n int) int {
	count := 0
	if n <= 0 {
		return 0
	}

	factor := 1
	for factor <= n {
		lower := n % factor
		cur := (n / factor) % 10
		higher := n / (factor * 10)

		switch cur {
		case 0:
			count += higher * factor
		case 1:
			count += higher*factor + lower + 1
		default:
			count += (higher + 1) * factor
		}

		// Check overflow
		if factor > n/10 {
			break
		}
		factor *= 10
	}

	return count
}

func main() {
	fmt.Println(countDigitOne(13))
	fmt.Println(countDigitOne(0))
}
```
