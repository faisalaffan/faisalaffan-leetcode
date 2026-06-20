# 1742 — Maximum Number Of Balls In A Box

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountBalls(lowLimit int, highLimit int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n) - but n <= 10^5, fine  
**Kompleksitas Ruang:** O(n) - but n <= 10^5, fine

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1742: Maximum Number of Balls in a Box
// https://leetcode.com/problems/maximum-number-of-balls-in-a-box/
// Difficulty: Easy

import "fmt"

// Time: O(n log n), Space: O(n) - but n <= 10^5, fine
func CountBalls(lowLimit int, highLimit int) int {
  // Membuat map (HashMap) — pencarian O(1)
	boxes := make(map[int]int)
	maxBalls := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := digitSum(i)
		boxes[sum]++
		if boxes[sum] > maxBalls {
			maxBalls = boxes[sum]
		}
	}
	return maxBalls
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(CountBalls(1, 10))
	fmt.Println(CountBalls(5, 15))
	fmt.Println(CountBalls(19, 28))
}
```
