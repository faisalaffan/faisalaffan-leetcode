# 0149 — Max Points On A Line

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxPoints(points [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #149: Max Points on a Line
// https://leetcode.com/problems/max-points-on-a-line/
// Difficulty: Hard

import (
	"fmt"
)

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxCount := 0

	for i := 0; i < n; i++ {
  // Membuat map (HashMap) — pencarian O(1)
		slopes := make(map[[2]int]int)
		duplicate := 1

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 && dy == 0 {
				duplicate++
				continue
			}

			g := gcd(dx, dy)
			dx /= g
			dy /= g

			// Normalize slope sign: ensure dx is positive, or if dx==0, dy positive
			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}

			key := [2]int{dx, dy}
			slopes[key]++
		}

		localMax := 0
		for _, count := range slopes {
			if count > localMax {
				localMax = count
			}
		}

		if localMax+duplicate > maxCount {
			maxCount = localMax + duplicate
		}
	}

	return maxCount
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	result := maxPoints(points)
	expected := 3

	fmt.Printf("maxPoints(%v) = %d\n", points, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
