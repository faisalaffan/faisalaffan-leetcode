# 3633 — Earliest Finish Time For Land And Water Rides I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func EarliestFinishTimeForLandAndWaterRidesI(landStartTime, landDuration, waterStartTime, waterDuration []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n*m) where n = len(landStartTime), m = len(waterStartTime)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3633: Earliest Finish Time for Land and Water Rides I
// https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{5}, []int{3}, []int{1}, []int{10}))
}

// Time: O(n*m) where n = len(landStartTime), m = len(waterStartTime)
// Space: O(1)
func EarliestFinishTimeForLandAndWaterRidesI(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	ans := int(^uint(0) >> 1) // max int

	// Land -> Water
  // Range loop: iterasi dengan indeks + nilai
	for i := range landStartTime {
		finishLand := landStartTime[i] + landDuration[i]
		for j := range waterStartTime {
			startWater := max(finishLand, waterStartTime[j])
			finish := startWater + waterDuration[j]
			if finish < ans {
				ans = finish
			}
		}
	}

	// Water -> Land
  // Range loop: iterasi dengan indeks + nilai
	for i := range waterStartTime {
		finishWater := waterStartTime[i] + waterDuration[i]
		for j := range landStartTime {
			startLand := max(finishWater, landStartTime[j])
			finish := startLand + landDuration[j]
			if finish < ans {
				ans = finish
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
