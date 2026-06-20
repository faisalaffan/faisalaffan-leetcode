# 1134 — Armstrong Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func isArmstrong(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1134: Armstrong Number
// https://leetcode.com/problems/armstrong-number/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isArmstrong(153))  // true
	fmt.Println(isArmstrong(123))  // false
	fmt.Println(isArmstrong(1))    // true
}

// LeetCode submission: isArmstrong
func isArmstrong(n int) bool {
	digits := 0
	for x := n; x > 0; x /= 10 {
		digits++
	}
	sum := 0
	for x := n; x > 0; x /= 10 {
		d := x % 10
		p := 1
		for i := 0; i < digits; i++ {
			p *= d
		}
		sum += p
	}
	return sum == n
}
```
