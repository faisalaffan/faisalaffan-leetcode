# 0365 — Water And Jug Problem

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canMeasureWater(x int, y int, target int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(log min(x,y))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #365: Water and Jug Problem
// https://leetcode.com/problems/water-and-jug-problem/
// Difficulty: Medium
// Time: O(log min(x,y)) | Space: O(1)

import "fmt"

func canMeasureWater(x int, y int, target int) bool {
	if target > x+y {
		return false
	}
	return target%gcd(x, y) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canMeasureWater(3, 5, 4))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canMeasureWater(2, 6, 5))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", canMeasureWater(1, 2, 3))
	// Expected: true
}
```
