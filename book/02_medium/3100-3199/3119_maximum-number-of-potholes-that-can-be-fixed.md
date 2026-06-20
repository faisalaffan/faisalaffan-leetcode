# 3119 — Maximum Number Of Potholes That Can Be Fixed

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxPotholes(road string, budget int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3119: Maximum Number of Potholes That Can Be Fixed
// https://leetcode.com/problems/maximum-number-of-potholes-that-can-be-fixed/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxPotholes(road string, budget int) int {
	var segs []int
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(road); i++ {
		if road[i] == 'x' {
			count++
		} else {
			if count > 0 {
				segs = append(segs, count)
				count = 0
			}
		}
	}
	if count > 0 {
		segs = append(segs, count)
	}

  // Custom sort dengan comparator
	sort.Slice(segs, func(i, j int) bool {
		return segs[i] > segs[j]
	})

	ans := 0
	for _, seg := range segs {
		cost := seg + 1
		if budget >= cost {
			budget -= cost
			ans += seg
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPotholes("...xxx..xx", 7))  // Expected: 5
	fmt.Println(maxPotholes("..xxxxx", 4))     // Expected: 3
	fmt.Println(maxPotholes("x.x.x.x", 10))    // Expected: 4
}
```
