# 1688 — Count Of Matches In Tournament

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfMatches(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1688: Count of Matches in Tournament
// https://leetcode.com/problems/count-of-matches-in-tournament/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func NumberOfMatches(n int) int {
	matches := 0
	for n > 1 {
		if n%2 == 0 {
			matches += n / 2
			n /= 2
		} else {
			matches += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return matches
}

// Alternative O(1): return n-1 (each match eliminates one team, champion is last)

func main() {
	fmt.Println(NumberOfMatches(7))
	fmt.Println(NumberOfMatches(14))
}
```
