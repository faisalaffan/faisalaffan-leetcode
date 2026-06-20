# 3740 — Minimum Distance Between Three Equal Elements I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumDistanceBetweenThreeEqualElementsI(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3740: Minimum Distance Between Three Equal Elements I
// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-i/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1, 2, 1, 1, 3}))
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1, 1, 2, 3, 2, 1, 2}))
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1}))
}

// Time: O(n)
// Space: O(n)
func MinimumDistanceBetweenThreeEqualElementsI(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	ans := math.MaxInt
	for _, indices := range pos {
		if len(indices) < 3 {
			continue
		}
		for i := 0; i <= len(indices)-3; i++ {
			dist := 2 * (indices[i+2] - indices[i])
			if dist < ans {
				ans = dist
			}
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```
