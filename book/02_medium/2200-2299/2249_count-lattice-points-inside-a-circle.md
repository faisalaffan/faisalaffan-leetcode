# 2249 — Count Lattice Points Inside A Circle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countLatticePoints(circles [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * r^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2249: Count Lattice Points Inside a Circle
// https://leetcode.com/problems/count-lattice-points-inside-a-circle/
// Difficulty: Medium
// Time: O(n * r^2) | Space: O(n)

import "fmt"

func countLatticePoints(circles [][]int) int {
  // Membuat map (HashMap) — pencarian O(1)
	points := make(map[[2]int]bool)

	for _, c := range circles {
		x, y, r := c[0], c[1], c[2]
		rr := r * r
		for dx := -r; dx <= r; dx++ {
			for dy := -r; dy <= r; dy++ {
				if dx*dx+dy*dy <= rr {
					points[[2]int{x + dx, y + dy}] = true
				}
			}
		}
	}
	return len(points)
}

func main() {
	// Test case 1
	fmt.Println(countLatticePoints([][]int{{2, 2, 1}}))
	// Expected: 5

	// Test case 2
	fmt.Println(countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
	// Expected: 16
}
```
