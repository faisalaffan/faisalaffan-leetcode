# 2358 — Maximum Number Of Groups Entering A Competition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumGroups(grades []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2358: Maximum Number of Groups Entering a Competition
// https://leetcode.com/problems/maximum-number-of-groups-entering-a-competition/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func maximumGroups(grades []int) int {
	n := len(grades)
	// We need k such that 1 + 2 + ... + k <= n
	// k(k+1)/2 <= n
	// Solve: k^2 + k - 2n <= 0
	// k = (-1 + sqrt(1 + 8n)) / 2
	k := 0
	for (k+1)*(k+2)/2 <= n {
		k++
	}
	return k
}

func main() {
	// Test case 1
	fmt.Println(maximumGroups([]int{10, 6, 12, 7, 3, 5}))
	// Expected: 3

	// Test case 2
	fmt.Println(maximumGroups([]int{8, 8}))
	// Expected: 1

	// Test case 3
	fmt.Println(maximumGroups([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// Expected: 4
}
```
