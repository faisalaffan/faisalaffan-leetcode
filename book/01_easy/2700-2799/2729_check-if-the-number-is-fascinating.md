# 2729 — Check If The Number Is Fascinating

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfTheNumberIsFascinating(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2729: Check if The Number is Fascinating
// https://leetcode.com/problems/check-if-the-number-is-fascinating/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(CheckIfTheNumberIsFascinating(192))
	fmt.Println(CheckIfTheNumberIsFascinating(100))
}

func CheckIfTheNumberIsFascinating(n int) bool {
	concat := strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3)
	if len(concat) != 9 {
		return false
	}

	digits := []byte(concat)
  // Custom sort dengan comparator
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
	return string(digits) == "123456789"
}
```
