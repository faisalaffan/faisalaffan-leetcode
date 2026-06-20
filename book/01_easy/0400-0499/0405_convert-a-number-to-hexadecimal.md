# 0405 — Convert A Number To Hexadecimal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func ConvertANumberToHexadecimal(num int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #405: Convert a Number to Hexadecimal
// https://leetcode.com/problems/convert-a-number-to-hexadecimal/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func ConvertANumberToHexadecimal(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	var result []byte
	// Use uint32 to handle negative numbers via two's complement.
	n := uint32(num)
	for n > 0 {
		result = append([]byte{hex[n&0xf]}, result...)
		n >>= 4
	}
	return string(result)
}

func main() {
	fmt.Println(ConvertANumberToHexadecimal(26))
	fmt.Println(ConvertANumberToHexadecimal(-1))
	fmt.Println(ConvertANumberToHexadecimal(0))
}
```
