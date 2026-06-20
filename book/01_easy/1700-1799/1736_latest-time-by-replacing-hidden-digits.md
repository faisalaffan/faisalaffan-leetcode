# 1736 — Latest Time By Replacing Hidden Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumTime(time string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1736: Latest Time by Replacing Hidden Digits
// https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func MaximumTime(time string) string {
	t := []byte(time)
	if t[0] == '?' {
		if t[1] == '?' || t[1] <= '3' {
			t[0] = '2'
		} else {
			t[0] = '1'
		}
	}
	if t[1] == '?' {
		if t[0] == '2' {
			t[1] = '3'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}

func main() {
	fmt.Println(MaximumTime("2?:?0"))
	fmt.Println(MaximumTime("0?:3?"))
	fmt.Println(MaximumTime("1?:22"))
}
```
