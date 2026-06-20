# 3032 — Count Numbers With Unique Digits Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountNumbersWithUniqueDigitsIi(a int, b int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O((b-a) * log(b))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3032: Count Numbers With Unique Digits II
// https://leetcode.com/problems/count-numbers-with-unique-digits-ii/
// Difficulty: Easy [Paid]
//
// Note: This is a premium problem. In Go, we implement the equivalent logic.

import "fmt"

func main() {
	// LeetCode name: numberCount
	fmt.Println(CountNumbersWithUniqueDigitsIi(1, 20))  // 19
	fmt.Println(CountNumbersWithUniqueDigitsIi(9, 19))  // 10
}

// Time: O((b-a) * log(b)) | Space: O(1)
// LeetCode submission name: numberCount
func CountNumbersWithUniqueDigitsIi(a int, b int) int {
	count := 0
	for i := a; i <= b; i++ {
		if hasUniqueDigits(i) {
			count++
		}
	}
	return count
}

func hasUniqueDigits(n int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for n > 0 {
		digit := n % 10
		if seen[digit] {
			return false
		}
		seen[digit] = true
		n /= 10
	}
	return true
}
```
