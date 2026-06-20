# 3848 — Check Digitorial Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckDigitorialPermutation(n int) bool
```

> **💡 Hint:** Compute sum of factorials of digits, check if any permutation

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3848: Check Digitorial Permutation
// https://leetcode.com/problems/check-digitorial-permutation/
// Difficulty: Medium
// Time: O(log N) | Space: O(1)
// Approach: Compute sum of factorials of digits, check if any permutation
// of n equals that sum (i.e., they have same digit frequency).

import "fmt"

func CheckDigitorialPermutation(n int) bool {
	// Precompute factorials for digits 0-9
	fact := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

	// Compute sum of factorials of digits
	temp := n
	sum := 0
	for temp > 0 {
		sum += fact[temp%10]
		temp /= 10
	}

	// Check if sum has same digits as n
	// If they have the same digits, one is a permutation of the other
  // Alokasi slice integer
	digitsN := make([]int, 10)
  // Alokasi slice integer
	digitsSum := make([]int, 10)

	temp = n
	for temp > 0 {
		digitsN[temp%10]++
		temp /= 10
	}

	temp = sum
	for temp > 0 {
		digitsSum[temp%10]++
		temp /= 10
	}

	// Handle sum = 0 (if n = 0, but n >= 1 per constraints)
	if sum == 0 {
		return false
	}

	for i := 0; i < 10; i++ {
		if digitsN[i] != digitsSum[i] {
			return false
		}
	}

	return true
}

func main() {
	// Example 1
	fmt.Println(CheckDigitorialPermutation(145)) // Expected: true

	// Example 2
	fmt.Println(CheckDigitorialPermutation(10)) // Expected: false

	// Example 3
	fmt.Println(CheckDigitorialPermutation(40585)) // 4!+0!+5!+8!+5! = 24+1+120+40320+120 = 40585
}
```
