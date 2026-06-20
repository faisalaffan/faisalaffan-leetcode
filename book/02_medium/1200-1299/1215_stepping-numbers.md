# 1215 — Stepping Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countSteppingNumbers(low int, high int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(2^n) where n = number of digits in high  
**Kompleksitas Ruang:** O(2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1215: Stepping Numbers
// https://leetcode.com/problems/stepping-numbers/
// Difficulty: Medium [Paid]

// Stepping numbers are numbers where adjacent digits differ by 1.
// Return all stepping numbers in range [low, high] sorted.

// Time: O(2^n) where n = number of digits in high
// Space: O(2^n)

func countSteppingNumbers(low int, high int) []int {
  // Alokasi slice integer
	result := make([]int, 0)

	var dfs func(num int)
	dfs = func(num int) {
		if num > high {
			return
		}
		if num >= low {
			result = append(result, num)
		}
		lastDigit := num % 10
		if lastDigit > 0 {
			dfs(num*10 + (lastDigit - 1))
		}
		if lastDigit < 9 {
			dfs(num*10 + (lastDigit + 1))
		}
	}

	for i := 0; i <= 9; i++ {
		dfs(i)
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [0 1 2 3 4 5 6 7 8 9 10 12])\n", countSteppingNumbers(0, 12))
	fmt.Printf("%v (expected: [10 12])\n", countSteppingNumbers(10, 14))
}
```
