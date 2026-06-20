# 2595 — Number Of Even And Odd Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfEvenAndOddBits(n int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2595: Number of Even and Odd Bits
// https://leetcode.com/problems/number-of-even-and-odd-bits/
// Difficulty: Easy
// Time O(log n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfEvenAndOddBits(17)) // [2,0]
	fmt.Println(NumberOfEvenAndOddBits(2))  // [0,1]
}

func NumberOfEvenAndOddBits(n int) []int {
	even, odd := 0, 0
	idx := 0
	for n > 0 {
		if n&1 == 1 {
			if idx%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		n >>= 1
		idx++
	}
	return []int{even, odd}
}
```
