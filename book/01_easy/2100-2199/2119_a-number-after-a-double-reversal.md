# 2119 — A Number After A Double Reversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func ANumberAfterADoubleReversal(num int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2119: A Number After a Double Reversal
// https://leetcode.com/problems/a-number-after-a-double-reversal/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ANumberAfterADoubleReversal(526))  // true
	fmt.Println(ANumberAfterADoubleReversal(1800)) // false
	fmt.Println(ANumberAfterADoubleReversal(0))    // true
}

// Time: O(1), Space: O(1)
func ANumberAfterADoubleReversal(num int) bool {
	// Reversing twice yields the same iff num has no trailing zeros
	return num == 0 || num%10 != 0
}
```
