# 0246 — Strobogrammatic Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func IsStrobogrammatic(num string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #246: Strobogrammatic Number
// https://leetcode.com/problems/strobogrammatic-number/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(1)
func IsStrobogrammatic(num string) bool {
	pairs := map[byte]byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}
	i, j := 0, len(num)-1
	for i <= j {
		if v, ok := pairs[num[i]]; !ok || v != num[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(IsStrobogrammatic("69"))
	fmt.Println(IsStrobogrammatic("88"))
	fmt.Println(IsStrobogrammatic("962"))
}
```
