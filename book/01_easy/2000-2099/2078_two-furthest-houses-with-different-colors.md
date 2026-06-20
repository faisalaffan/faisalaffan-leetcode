# 2078 — Two Furthest Houses With Different Colors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TwoFurthestHousesWithDifferentColors(colors []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2078: Two Furthest Houses With Different Colors
// https://leetcode.com/problems/two-furthest-houses-with-different-colors/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{1, 1, 1, 6, 1, 1, 1})) // 3
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{1, 2, 3, 4, 5}))       // 4
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{0, 1}))                // 1
}

// Time: O(n), Space: O(1)
func TwoFurthestHousesWithDifferentColors(colors []int) int {
	n := len(colors)
	maxDist := 0

	// Check from leftmost with rightmost
	if colors[0] != colors[n-1] {
		return n - 1
	}

	// If ends are same, find furthest different color from either end
	for i := 1; i < n-1; i++ {
		if colors[i] != colors[0] {
			dist := n - 1 - i
			if dist > maxDist {
				maxDist = dist
			}
			dist = i
			if dist > maxDist {
				maxDist = dist
			}
			break
		}
	}
	return maxDist
}
```
