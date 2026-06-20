# 2566 — Maximum Difference By Remapping A Digit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumDifferenceByRemappingADigit(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2566: Maximum Difference by Remapping a Digit
// https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MaximumDifferenceByRemappingADigit(11891)) // 99009
	fmt.Println(MaximumDifferenceByRemappingADigit(90))    // 99
}

func MaximumDifferenceByRemappingADigit(num int) int {
	s := strconv.Itoa(num)

	// Find max: replace first non-9 digit with 9
	maxStr := []byte(s)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(maxStr); i++ {
		if maxStr[i] != '9' {
			replaceWith := maxStr[i]
			for j := i; j < len(maxStr); j++ {
				if maxStr[j] == replaceWith {
					maxStr[j] = '9'
				}
			}
			break
		}
	}
	maxVal, _ := strconv.Atoi(string(maxStr))

	// Find min: replace first non-0 digit (or non-1) with 0
	minStr := []byte(s)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(minStr); i++ {
		if minStr[i] != '0' && minStr[i] != '1' {
			replaceWith := minStr[i]
			for j := i; j < len(minStr); j++ {
				if minStr[j] == replaceWith {
					if i == 0 {
						minStr[j] = '1'
					} else {
						minStr[j] = '0'
					}
				}
			}
			break
		}
	}
	minVal, _ := strconv.Atoi(string(minStr))

	return maxVal - minVal
}
```
