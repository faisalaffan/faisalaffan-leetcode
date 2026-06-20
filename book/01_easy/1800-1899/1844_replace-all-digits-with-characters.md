# 1844 — Replace All Digits With Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReplaceDigits(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1844: Replace All Digits with Characters
// https://leetcode.com/problems/replace-all-digits-with-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReplaceDigits(s string) string {
	result := []byte(s)
	for i := 1; i < len(s); i += 2 {
		result[i] = s[i-1] + (s[i] - '0')
	}
	return string(result)
}

func main() {
	fmt.Println(ReplaceDigits("a1c1e1"))
	fmt.Println(ReplaceDigits("a1b2c3d4e"))
}
```
