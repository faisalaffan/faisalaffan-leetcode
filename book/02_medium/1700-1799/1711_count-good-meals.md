# 1711 — Count Good Meals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPairs(deliciousness []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * log(maxVal)), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1711: Count Good Meals
// https://leetcode.com/problems/count-good-meals/
// Difficulty: Medium
// Time: O(n * log(maxVal)), Space: O(n)

import "fmt"

const mod = 1_000_000_007

func countPairs(deliciousness []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	result := 0

	for _, d := range deliciousness {
		// Check each power of 2
		for sum := 1; sum <= (1 << 21); sum <<= 1 {
			complement := sum - d
			if cnt, ok := count[complement]; ok {
				result = (result + cnt) % mod
			}
		}
		count[d]++
	}
	return result
}

func main() {
	fmt.Println(countPairs([]int{1, 3, 5, 7, 9}))    // Expected: 4
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7})) // Expected: 15
	fmt.Println(countPairs([]int{2, 4, 6, 8})) // Expected: 3
}
```
