# 2598 — Smallest Missing Non Negative Integer After Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func findSmallestInteger(nums []int, value int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2598: Smallest Missing Non-negative Integer After Operations
// https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findSmallestInteger(nums []int, value int) int {
  // Alokasi slice integer
	freq := make([]int, value)
	for _, v := range nums {
		// Map to [0, value-1] range
		rem := ((v % value) + value) % value
		freq[rem]++
	}

	for i := 0; ; i++ {
		if freq[i%value] == 0 {
			return i
		}
		freq[i%value]--
	}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 5))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", findSmallestInteger([]int{1, 2, 3, 4, 5}, 2))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", findSmallestInteger([]int{0, 0, 0, 0}, 1))
	// Expected: 4
}
```
