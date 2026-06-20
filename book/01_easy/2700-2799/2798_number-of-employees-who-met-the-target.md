# 2798 — Number Of Employees Who Met The Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfEmployeesWhoMetTheTarget(hours []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2798: Number of Employees Who Met the Target
// https://leetcode.com/problems/number-of-employees-who-met-the-target/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{5, 1, 4, 2, 2}, 6))
}

func NumberOfEmployeesWhoMetTheTarget(hours []int, target int) int {
	count := 0
	for _, h := range hours {
		if h >= target {
			count++
		}
	}
	return count
}
```
