# 1431 — Kids With The Greatest Number Of Candies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func kidsWithCandies(candies []int, extraCandies int) []bool

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1431: Kids With the Greatest Number of Candies
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Difficulty: Easy
//
// LeetCode submission: func kidsWithCandies(candies []int, extraCandies int) []bool

import "fmt"

func main() {
	fmt.Println(KidsWithTheGreatestNumberOfCandies([]int{2, 3, 5, 1, 3}, 3)) // [true true true false true]
	fmt.Println(KidsWithTheGreatestNumberOfCandies([]int{4, 2, 1, 1, 2}, 1)) // [true false false false false]
}

// Time: O(n), Space: O(n)
func KidsWithTheGreatestNumberOfCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	for _, c := range candies {
		if c > maxCandy {
			maxCandy = c
		}
	}
	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= maxCandy
	}
	return res
}
```
