# 3783 — Mirror Distance Of An Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MirrorDistanceOfAnInteger(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3783: Mirror Distance of an Integer
// https://leetcode.com/problems/mirror-distance-of-an-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MirrorDistanceOfAnInteger(25))
	fmt.Println(MirrorDistanceOfAnInteger(10))
	fmt.Println(MirrorDistanceOfAnInteger(7))
}

// Time: O(log n)
// Space: O(1)
func MirrorDistanceOfAnInteger(n int) int {
	rev := 0
	for x := n; x > 0; x /= 10 {
		rev = rev*10 + x%10
	}
	if n > rev {
		return n - rev
	}
	return rev - n
}
```
