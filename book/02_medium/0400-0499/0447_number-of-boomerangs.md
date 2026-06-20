# 0447 — Number Of Boomerangs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfBoomerangs(points [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #447: Number of Boomerangs
// https://leetcode.com/problems/number-of-boomerangs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	total := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(points); i++ {
  // Membuat map (HashMap) — pencarian O(1)
		distCount := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := dx*dx + dy*dy
			distCount[dist]++
		}
		for _, count := range distCount {
			total += count * (count - 1)
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numberOfBoomerangs([][]int{{0, 0}}))
	// Expected: 0
}
```
