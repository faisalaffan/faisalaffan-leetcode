# 2169 — Count Operations To Obtain Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountOperationsToObtainZero(num1 int, num2 int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log max(num1, num2)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2169: Count Operations to Obtain Zero
// https://leetcode.com/problems/count-operations-to-obtain-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountOperationsToObtainZero(2, 3))  // 3
	fmt.Println(CountOperationsToObtainZero(10, 10)) // 1
}

// Time: O(log max(num1, num2)), Space: O(1)
func CountOperationsToObtainZero(num1 int, num2 int) int {
	ops := 0
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		ops++
	}
	return ops
}
```
