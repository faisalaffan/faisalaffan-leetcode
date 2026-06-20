# 2224 — Minimum Number Of Operations To Convert Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfOperationsToConvertTime(current string, correct string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2224: Minimum Number of Operations to Convert Time
// https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfOperationsToConvertTime("02:30", "04:35")) // 3
	fmt.Println(MinimumNumberOfOperationsToConvertTime("11:00", "11:01")) // 1
}

// Time: O(1), Space: O(1)
func MinimumNumberOfOperationsToConvertTime(current string, correct string) int {
	cur := int(current[0]-'0')*600 + int(current[1]-'0')*60 + int(current[3]-'0')*10 + int(current[4]-'0')
	cor := int(correct[0]-'0')*600 + int(correct[1]-'0')*60 + int(correct[3]-'0')*10 + int(correct[4]-'0')

	diff := cor - cur
	ops := 0

	ops += diff / 60
	diff %= 60
	ops += diff / 15
	diff %= 15
	ops += diff / 5
	diff %= 5
	ops += diff

	return ops
}
```
