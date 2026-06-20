# 2409 — Count Days Spent Together

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func dayOfYear(date string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2409: Count Days Spent Together
// https://leetcode.com/problems/count-days-spent-together/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

var daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	fmt.Println(CountDaysSpentTogether("08-15", "08-18", "08-16", "08-19")) // 3
	fmt.Println(CountDaysSpentTogether("10-01", "10-31", "11-01", "12-31")) // 0
}

func dayOfYear(date string) int {
	mm := int(date[0]-'0')*10 + int(date[1]-'0')
	dd := int(date[3]-'0')*10 + int(date[4]-'0')
	days := 0
	for m := 0; m < mm-1; m++ {
		days += daysInMonth[m]
	}
	return days + dd
}

func CountDaysSpentTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aStart := dayOfYear(arriveAlice)
	aEnd := dayOfYear(leaveAlice)
	bStart := dayOfYear(arriveBob)
	bEnd := dayOfYear(leaveBob)

	start := max(aStart, bStart)
	end := min(aEnd, bEnd)

	if start > end {
		return 0
	}
	return end - start + 1
}
```
