# 2678 — Number Of Senior Citizens

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfSeniorCitizens(details []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2678: Number of Senior Citizens
// https://leetcode.com/problems/number-of-senior-citizens/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(NumberOfSeniorCitizens([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
	fmt.Println(NumberOfSeniorCitizens([]string{"1313579440F2036", "2921522980M5644"}))
}

func NumberOfSeniorCitizens(details []string) int {
	count := 0
	for _, d := range details {
		age, _ := strconv.Atoi(d[11:13])
		if age > 60 {
			count++
		}
	}
	return count
}
```
