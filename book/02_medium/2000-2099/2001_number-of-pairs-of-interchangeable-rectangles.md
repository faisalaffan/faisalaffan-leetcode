# 2001 — Number Of Pairs Of Interchangeable Rectangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfPairsOfInterchangeableRectangles(rectangles [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log max(w,h)), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2001: Number of Pairs of Interchangeable Rectangles
// https://leetcode.com/problems/number-of-pairs-of-interchangeable-rectangles/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfPairsOfInterchangeableRectangles([][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}))
	fmt.Println(NumberOfPairsOfInterchangeableRectangles([][]int{{4, 5}, {7, 8}}))
}

// Time: O(n log max(w,h)), Space: O(n)
func NumberOfPairsOfInterchangeableRectangles(rectangles [][]int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[[2]int]int64)
	for _, r := range rectangles {
		w, h := r[0], r[1]
		g := gcd2001(w, h)
		key := [2]int{w / g, h / g}
		cnt[key]++
	}

	var ans int64
	for _, m := range cnt {
		ans += m * (m - 1) / 2
	}
	return ans
}

func gcd2001(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```
