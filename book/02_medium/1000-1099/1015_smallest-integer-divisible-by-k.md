# 1015 — Smallest Integer Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestRepunitDivByK(k int) int
```

> **💡 Hint:** Build remainder incrementally. Keep dividing.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(K)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1015: Smallest Integer Divisible by K
// https://leetcode.com/problems/smallest-integer-divisible-by-k/
// Difficulty: Medium
//
// Approach: Build remainder incrementally. Keep dividing.
//   N = 1, 11, 111, ... = N*10 + 1 (mod K)
//   If N % K == 0, return length. If remainder repeats, return -1.
// Time: O(K)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(smallestRepunitDivByK(1))  // 1
	fmt.Println(smallestRepunitDivByK(2))  // -1
	fmt.Println(smallestRepunitDivByK(3))  // 3
}

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	remainder := 0
	for length := 1; length <= k; length++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return length
		}
	}

	return -1
}
```
