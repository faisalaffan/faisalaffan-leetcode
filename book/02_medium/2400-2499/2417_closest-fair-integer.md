# 2417 — Closest Fair Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func closestFair(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n * 2^d)  
**Kompleksitas Ruang:** O(d)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2417: Closest Fair Integer
// https://leetcode.com/problems/closest-fair-integer/
// Difficulty: Medium
// Time: O(log n * 2^d) | Space: O(d)
// Find smallest fair integer >= n (equal even and odd digit count).

import "fmt"

func main() {
	fmt.Println(closestFair(2))    // 10
	fmt.Println(closestFair(403))  // 440
	fmt.Println(closestFair(10))   // 10
}

func closestFair(n int) int {
	for {
		if isFair(n) {
			return n
		}
		n++
	}
}

func isFair(n int) bool {
	s := fmt.Sprint(n)
	even, odd := 0, 0
	for _, ch := range s {
		d := int(ch - '0')
		if d%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even == odd
}
```
