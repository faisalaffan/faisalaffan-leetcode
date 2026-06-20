# 2427 — Number Of Common Factors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfCommonFactors(a int, b int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2427: Number of Common Factors
// https://leetcode.com/problems/number-of-common-factors/
// Difficulty: Easy
// Time O(min(a,b)) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfCommonFactors(12, 6)) // 4
	fmt.Println(NumberOfCommonFactors(25, 30)) // 2
}

func NumberOfCommonFactors(a int, b int) int {
	count := 0
	n := a
	if b < n {
		n = b
	}
	for i := 1; i <= n; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}
```
