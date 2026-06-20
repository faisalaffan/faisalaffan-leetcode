# 0754 — Reach A Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func reachNumber(target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(target))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #754: Reach a Number
// https://leetcode.com/problems/reach-a-number/
// Difficulty: Medium
// Time: O(sqrt(target))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(reachNumber(3))
	fmt.Println(reachNumber(2))
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	sum := 0
	steps := 0

	for sum < target || (sum-target)%2 != 0 {
		steps++
		sum += steps
	}

	return steps
}
```
