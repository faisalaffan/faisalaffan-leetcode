# 2620 — Counter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func counter(n int) func() int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2620: Counter
// https://leetcode.com/problems/counter/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript closure problem, adapted to Go. Returns a counter function.

import "fmt"

func main() {
	counter := counter(10)
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func counter(n int) func() int {
	return func() int {
		n++
		return n - 1
	}
}
```
