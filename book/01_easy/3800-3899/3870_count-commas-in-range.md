# 3870 — Count Commas In Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountCommasInRange(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3870: Count Commas in Range
// https://leetcode.com/problems/count-commas-in-range/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountCommasInRange(1002))
	fmt.Println(CountCommasInRange(998))
	fmt.Println(CountCommasInRange(1500000))
}

// Time: O(log n)
// Space: O(1)
func CountCommasInRange(n int) int {
	total := 0
	threshold := 1000
	for n >= threshold {
		total += n - threshold + 1
		threshold *= 1000
	}
	return total
}
```
