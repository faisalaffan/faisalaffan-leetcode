# 2189 — Number Of Ways To Build House Of Cards

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func houseOfCards(remainingCards int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2189: Number of Ways to Build House of Cards
// https://leetcode.com/problems/number-of-ways-to-build-house-of-cards/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n)

import "fmt"

func houseOfCards(remainingCards int) int {
	// DP: dp[c] = number of ways to use c cards
  // Alokasi slice integer
	dp := make([]int, remainingCards+1)
	dp[0] = 1

	// Each row (level) uses 3k - 1 cards where k >= 2
	for k := 2; 3*k-1 <= remainingCards; k++ {
		cards := 3*k - 1
		for c := remainingCards; c >= cards; c-- {
			dp[c] += dp[c-cards]
		}
	}
	return dp[remainingCards]
}

func main() {
	// Test case 1
	fmt.Println(houseOfCards(16))
	// Expected: 2

	// Test case 2
	fmt.Println(houseOfCards(2))
	// Expected: 0

	// Test case 3
	fmt.Println(houseOfCards(6))
	// Expected: 1
}
```
