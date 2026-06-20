# 3168 — Minimum Number Of Chairs In A Waiting Room

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfChairsInAWaitingRoom(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3168: Minimum Number of Chairs in a Waiting Room
// https://leetcode.com/problems/minimum-number-of-chairs-in-a-waiting-room/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumChairs
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("EEEE"))   // 4
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("ELELEL")) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minimumChairs
func MinimumNumberOfChairsInAWaitingRoom(s string) int {
	current := 0
	maxChairs := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == 'E' {
			current++
			if current > maxChairs {
				maxChairs = current
			}
		} else {
			current--
		}
	}
	return maxChairs
}
```
