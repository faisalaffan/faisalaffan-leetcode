# 0970 — Powerful Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func powerfulIntegers(x int, y int, bound int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(log_x(bound) * log_y(bound))  
**Kompleksitas Ruang:** O(log_x(bound) * log_y(bound))

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #970: Powerful Integers
// https://leetcode.com/problems/powerful-integers/
// Difficulty: Medium

import "fmt"

// Time: O(log_x(bound) * log_y(bound)) | Space: O(log_x(bound) * log_y(bound))
func powerfulIntegers(x int, y int, bound int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)

	for a := 1; a <= bound; a *= x {
		for b := 1; a+b <= bound; b *= y {
			seen[a+b] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

  // Alokasi slice integer
	ans := make([]int, 0, len(seen))
	for v := range seen {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(3, 5, 15))
	fmt.Println(powerfulIntegers(2, 1, 10))
}
```
