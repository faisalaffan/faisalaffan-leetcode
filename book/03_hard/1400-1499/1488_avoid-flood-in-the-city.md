# 1488 — Avoid Flood In The City

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func avoidFlood(rains []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1488: Avoid Flood in The City
// https://leetcode.com/problems/avoid-flood-in-the-city/
// Difficulty: Hard
//
// rains[i] > 0 means lake rains[i] fills up.
// rains[i] == 0 means we can dry one lake.
// Return an array ans where ans[i] = -1 if rains[i] > 0,
// else ans[i] = the lake we dry (any). If flood cannot be avoided, return [].

import (
	"fmt"
	"sort"
)

// avoidFlood returns an array of actions to avoid flooding.
func avoidFlood(rains []int) []int {
	n := len(rains)
  // Alokasi slice integer
	ans := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range ans {
		ans[i] = 1 // default dry value for zero-rain days
	}

	// Map lake -> last day it rained
  // Membuat map (HashMap) — pencarian O(1)
	lastRain := make(map[int]int)

	// Collect dry days (when rains[i] == 0)
	var dryDays []int

	for i, lake := range rains {
		if lake == 0 {
			dryDays = append(dryDays, i)
			continue
		}

		ans[i] = -1 // it's raining on this day

		if prevDay, ok := lastRain[lake]; ok {
			// This lake is already full, we need to dry it before today
			// Find a dry day after prevDay to dry this lake
			idx := sort.Search(len(dryDays), func(j int) bool {
				return dryDays[j] > prevDay
			})

			if idx >= len(dryDays) {
				// Cannot dry this lake in time -> flood unavoidable
				return []int{}
			}

			dryDay := dryDays[idx]
			ans[dryDay] = lake
			// Remove this dry day from the list
			dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
		}

		lastRain[lake] = i
	}

	return ans
}

func main() {
	// Test case 1
	rains1 := []int{1, 2, 3, 4}
	result1 := avoidFlood(rains1)
	fmt.Printf("Test 1: rains=%v => %v (expected [-1,-1,-1,-1])\n", rains1, result1)

	// Test case 2
	rains2 := []int{1, 2, 0, 0, 2, 1}
	result2 := avoidFlood(rains2)
	fmt.Printf("Test 2: rains=%v => %v (expected [-1,-1,2,1,-1,-1])\n", rains2, result2)

	// Test case 3
	rains3 := []int{1, 2, 0, 1, 2}
	result3 := avoidFlood(rains3)
	fmt.Printf("Test 3: rains=%v => %v (expected [])\n", rains3, result3)

	// Test case 4
	rains4 := []int{69, 0, 0, 0, 69}
	result4 := avoidFlood(rains4)
	fmt.Printf("Test 4: rains=%v => %v\n", rains4, result4)

	// Test case 5: LeetCode example
	rains5 := []int{1, 0, 2, 0, 2, 1}
	result5 := avoidFlood(rains5)
	fmt.Printf("Test 5: rains=%v => %v\n", rains5, result5)

	// Test case 6: impossible - two floods without dry day
	rains6 := []int{1, 1}
	result6 := avoidFlood(rains6)
	fmt.Printf("Test 6: rains=%v => %v (expected [])\n", rains6, result6)
}
```
