# 1344 — Angle Between Hands Of A Clock

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func angleClock(hour int, minutes int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1344: Angle Between Hands of a Clock
// https://leetcode.com/problems/angle-between-hands-of-a-clock/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(angleClock(12, 30)) // 165

	// Test case 2
	fmt.Println(angleClock(3, 30)) // 75

	// Test case 3
	fmt.Println(angleClock(3, 15)) // 7.5

	// Test case 4
	fmt.Println(angleClock(4, 50)) // 155
}

// Time: O(1)
// Space: O(1)
func angleClock(hour int, minutes int) float64 {
	// Minute hand: 360 degrees in 60 minutes = 6 degrees per minute
	minAngle := float64(minutes) * 6.0

	// Hour hand: 360 degrees in 12 hours = 30 degrees per hour
	// Plus 0.5 degrees per minute (30 degrees / 60 minutes)
	hourAngle := float64(hour%12)*30.0 + float64(minutes)*0.5

	diff := math.Abs(hourAngle - minAngle)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}
```
