# 0191 — Number Of 1 Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func HammingWeight(num uint32) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #191: Number of 1 Bits
// https://leetcode.com/problems/number-of-1-bits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func HammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingWeight(11))    // 3
	fmt.Println(HammingWeight(128))   // 1
	fmt.Println(HammingWeight(4294967293)) // 31
}
```
