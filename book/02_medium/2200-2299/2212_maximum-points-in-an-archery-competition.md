# 2212 — Maximum Points In An Archery Competition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumBobPoints(numArrows int, aliceArrows []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(n * 2^n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2212: Maximum Points in an Archery Competition
// https://leetcode.com/problems/maximum-points-in-an-archery-competition/
// Difficulty: Medium
// Time: O(n * 2^n) | Space: O(n)

import "fmt"

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	bestScore := 0
	var bestMask int

	for mask := 1; mask < (1 << 12); mask++ {
		arrows := 0
		score := 0
		for i := 0; i < 12; i++ {
			if mask&(1<<i) != 0 {
				arrows += aliceArrows[i] + 1
				score += i
			}
		}
		if arrows <= numArrows && score > bestScore {
			bestScore = score
			bestMask = mask
		}
	}

  // Alokasi slice integer
	result := make([]int, 12)
	used := 0
	for i := 0; i < 12; i++ {
		if bestMask&(1<<i) != 0 {
			result[i] = aliceArrows[i] + 1
			used += result[i]
		}
	}
	// Put remaining arrows in first section (index 0)
	if used < numArrows {
		result[0] += numArrows - used
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumBobPoints(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}))
	// Expected: [0,0,0,0,0,0,0,0,1,1,1,0] or similar valid

	// Test case 2
	fmt.Println(maximumBobPoints(3, []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}))
	// Expected: [0,1,1,0,0,0,0,0,0,0,0,0] or similar valid
}
```
