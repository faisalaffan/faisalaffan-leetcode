# 1806 — Minimum Number Of Operations To Reinitialize A Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func reinitializePermutation(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1806: Minimum Number of Operations to Reinitialize a Permutation
// https://leetcode.com/problems/minimum-number-of-operations-to-reinitialize-a-permutation/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func reinitializePermutation(n int) int {
	ops := 0
	i := 1

	for {
		ops++
		if i%2 == 0 {
			i /= 2
		} else {
			i = n/2 + (i-1)/2
		}
		if i == 1 {
			break
		}
	}
	return ops
}

func main() {
	fmt.Println(reinitializePermutation(2))  // Expected: 1
	fmt.Println(reinitializePermutation(4))  // Expected: 2
	fmt.Println(reinitializePermutation(6))  // Expected: 4
}
```
