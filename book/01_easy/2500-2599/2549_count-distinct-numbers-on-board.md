# 2549 — Count Distinct Numbers On Board

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountDistinctNumbersOnBoard(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2549: Count Distinct Numbers on Board
// https://leetcode.com/problems/count-distinct-numbers-on-board/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountDistinctNumbersOnBoard(5)) // 4
	fmt.Println(CountDistinctNumbersOnBoard(2)) // 1
	fmt.Println(CountDistinctNumbersOnBoard(1)) // 1
}

func CountDistinctNumbersOnBoard(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}
```
