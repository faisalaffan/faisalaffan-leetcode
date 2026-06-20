# 2857 — Count Pairs Of Points With Distance K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPairsOfPointsWithDistanceK(coordinates [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2857: Count Pairs of Points With Distance k
// https://leetcode.com/problems/count-pairs-of-points-with-distance-k/
// Difficulty: Medium
// Time: O(n * k) | Space: O(n)

import "fmt"

func CountPairsOfPointsWithDistanceK(coordinates [][]int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	cache := make(map[[2]int]int)
	count := 0

	for _, coord := range coordinates {
		x, y := coord[0], coord[1]
		// XOR property: (x1 ^ x2) + (y1 ^ y2) = k
		// For each possible a, b such that a + b = k:
		for a := 0; a <= k; a++ {
			b := k - a
			px := x ^ a
			py := y ^ b
			count += cache[[2]int{px, py}]
		}
		cache[[2]int{x, y}]++
	}

	return count
}

func main() {
	fmt.Println(CountPairsOfPointsWithDistanceK([][]int{{1, 2}, {4, 2}, {1, 3}, {5, 2}}, 2))
	fmt.Println(CountPairsOfPointsWithDistanceK([][]int{{0, 0}, {1, 1}, {2, 2}}, 2))
}
```
