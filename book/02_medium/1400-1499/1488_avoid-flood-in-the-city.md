# 1488 — Avoid Flood In The City

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AvoidFlood(rains []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(N log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1488: Avoid Flood in The City
// https://leetcode.com/problems/avoid-flood-in-the-city/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AvoidFlood([]int{1, 2, 3, 4}))
	fmt.Println(AvoidFlood([]int{1, 2, 0, 0, 2, 1}))
	fmt.Println(AvoidFlood([]int{1, 2, 0, 1, 2}))
}

func AvoidFlood(rains []int) []int {
	// Time: O(N log N), Space: O(N)
	n := len(rains)
  // Alokasi slice integer
	ans := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range ans {
		ans[i] = -1 // default for rain days
	}

  // Membuat map (HashMap) — pencarian O(1)
	lastRain := make(map[int]int) // lake -> last rain day
  // Alokasi slice integer
	dryDays := make([]int, 0)     // indices of dry days (0s)

	for i, lake := range rains {
		if lake == 0 {
			dryDays = append(dryDays, i)
			ans[i] = 1 // placeholder
			continue
		}

		ans[i] = -1 // rain day, no action

		if prev, exists := lastRain[lake]; exists {
			// Find a dry day after prev to dry this lake
			idx := sort.Search(len(dryDays), func(j int) bool {
				return dryDays[j] > prev
			})
			if idx == len(dryDays) {
				return nil // impossible to prevent flood
			}
			ans[dryDays[idx]] = lake
			// Remove used dry day
			dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
		}
		lastRain[lake] = i
	}

	// Remaining dry days can be any positive number
	for _, idx := range dryDays {
		ans[idx] = 1
	}

	return ans
}
```
