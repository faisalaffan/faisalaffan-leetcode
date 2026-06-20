# 0171 — Excel Sheet Column Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func TitleToNumber(columnTitle string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #171: Excel Sheet Column Number
// https://leetcode.com/problems/excel-sheet-column-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func TitleToNumber(columnTitle string) int {
	result := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(columnTitle); i++ {
		result = result*26 + int(columnTitle[i]-'A'+1)
	}
	return result
}

func main() {
	fmt.Println(TitleToNumber("A"))
	fmt.Println(TitleToNumber("AB"))
	fmt.Println(TitleToNumber("ZY"))
}
```
