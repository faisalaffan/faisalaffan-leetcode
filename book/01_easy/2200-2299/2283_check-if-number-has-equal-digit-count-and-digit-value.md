# 2283 — Check If Number Has Equal Digit Count And Digit Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfNumberHasEqualDigitCountAndDigitValue(num string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2283: Check if Number Has Equal Digit Count and Digit Value
// https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("1210")) // true
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("030"))  // false
}

func CheckIfNumberHasEqualDigitCountAndDigitValue(num string) bool {
	count := [10]int{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(num); i++ {
		count[num[i]-'0']++
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(num); i++ {
		if count[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
```
