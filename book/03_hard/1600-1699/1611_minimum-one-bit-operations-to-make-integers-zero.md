# 1611 — Minimum One Bit Operations To Make Integers Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOneBitOperations(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1611: Minimum One Bit Operations to Make Integers Zero
// https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/
// Difficulty: Hard
//
// Solution: The minimum operations = inverse Gray code of n.
// Gray code: g(n) = n ^ (n>>1)
// Inverse Gray code (find position p such that g(p) = n):
//   ans = n ^ (n>>1) ^ (n>>2) ^ (n>>4) ^ ... (until zero)

import "fmt"

func minimumOneBitOperations(n int) int {
	ans := 0
	for n > 0 {
		ans ^= n
		n >>= 1
	}
	return ans
}

func main() {
	// Test case 1: 3 -> 2
	result := minimumOneBitOperations(3)
	fmt.Printf("minimumOneBitOperations(3) = %d (expected 2)\n", result)

	// Test case 2: 6 -> 4
	result2 := minimumOneBitOperations(6)
	fmt.Printf("minimumOneBitOperations(6) = %d (expected 4)\n", result2)

	// Test case 3: 0 -> 0
	result3 := minimumOneBitOperations(0)
	fmt.Printf("minimumOneBitOperations(0) = %d (expected 0)\n", result3)

	// Test case 4: 2 -> 3
	result4 := minimumOneBitOperations(2)
	fmt.Printf("minimumOneBitOperations(2) = %d (expected 3)\n", result4)

	// Test case 5: 9 -> 14
	result5 := minimumOneBitOperations(9)
	fmt.Printf("minimumOneBitOperations(9) = %d (expected 14)\n", result5)
}
```
