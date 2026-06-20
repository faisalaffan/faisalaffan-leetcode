# 1870 — Minimum Speed To Arrive On Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinSpeedOnTime(dist []int, hour float64) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log maxDist), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1870: Minimum Speed to Arrive on Time
// https://leetcode.com/problems/minimum-speed-to-arrive-on-time/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 6.0))
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 2.7))
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 1.9))
}

// Time: O(n log maxDist), Space: O(1)
func MinSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour <= float64(n-1) {
		return -1
	}

	lo, hi := 1, 10000000
	ans := -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canReach(dist, mid, hour) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func canReach(dist []int, speed int, hour float64) bool {
	time := 0.0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(dist)-1; i++ {
		time += float64((dist[i] + speed - 1) / speed) // ceil division
	}
	time += float64(dist[len(dist)-1]) / float64(speed)
	return time <= hour
}
```
