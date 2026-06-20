# 2843 — Count Symmetric Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSymmetricIntegers(low int, high int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(high - low)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2843: Count Symmetric Integers
// https://leetcode.com/problems/count-symmetric-integers/
// Difficulty: Easy
// Time: O(high - low) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CountSymmetricIntegers(1, 100))
	fmt.Println(CountSymmetricIntegers(1200, 1230))
}

func CountSymmetricIntegers(low int, high int) int {
	count := 0
	for n := low; n <= high; n++ {
		if isSymmetric(n) {
			count++
		}
	}
	return count
}

func isSymmetric(n int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s)%2 == 1 {
		return false
	}
	half := len(s) / 2
	sum1, sum2 := 0, 0
	for i := 0; i < half; i++ {
		sum1 += int(s[i] - '0')
		sum2 += int(s[i+half] - '0')
	}
	return sum1 == sum2
}
```
