# 2320 — Count Number Of Ways To Place Houses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countHousePlacements(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2320: Count Number of Ways to Place Houses
// https://leetcode.com/problems/count-number-of-ways-to-place-houses/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countHousePlacements(n int) int {
	const mod = 1_000_000_007
	// Fibonacci-ish: ways to place on one side
	a, b := 1, 1 // empty, single plot
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%mod
	}
	oneSide := b
	return (oneSide * oneSide) % mod
}

func main() {
	// Test case 1
	fmt.Println(countHousePlacements(1))
	// Expected: 4

	// Test case 2
	fmt.Println(countHousePlacements(2))
	// Expected: 9

	// Test case 3
	fmt.Println(countHousePlacements(3))
	// Expected: 25
}
```
