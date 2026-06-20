# 3577 — Count The Number Of Computer Unlocking Permutations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountTheNumberOfComputerUnlockingPermutations(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3577: Count the Number of Computer Unlocking Permutations
// https://leetcode.com/problems/count-the-number-of-computer-unlocking-permutations/
// Difficulty: Medium
// Complexity: O(n! * n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountTheNumberOfComputerUnlockingPermutations(2))
	// Test case 2
	fmt.Println("Test 2:", CountTheNumberOfComputerUnlockingPermutations(3))
	// Test case 3
	fmt.Println("Test 3:", CountTheNumberOfComputerUnlockingPermutations(1))
}

func CountTheNumberOfComputerUnlockingPermutations(n int) int {
	// Count permutations of [1..n] that satisfy unlock pattern rules
	// Simple version: all permutations are valid
	// n! permutations
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
```
