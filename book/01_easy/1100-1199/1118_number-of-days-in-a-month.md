# 1118 — Number Of Days In A Month

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfDays(year, month int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1118: Number of Days in a Month
// https://leetcode.com/problems/number-of-days-in-a-month/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(numberOfDays(1992, 7))  // 31
	fmt.Println(numberOfDays(2000, 2))  // 29
	fmt.Println(numberOfDays(1900, 2))  // 28
}

// LeetCode submission: numberOfDays
func numberOfDays(year, month int) int {
	leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leap {
		days[2] = 29
	}
	return days[month]
}
```
