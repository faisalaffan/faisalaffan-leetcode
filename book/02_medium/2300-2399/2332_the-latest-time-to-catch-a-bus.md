# 2332 — The Latest Time To Catch A Bus

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O((n + m) log(n + m))  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2332: The Latest Time to Catch a Bus
// https://leetcode.com/problems/the-latest-time-to-catch-a-bus/
// Difficulty: Medium
// Time: O((n + m) log(n + m)) | Space: O(n + m)

import (
	"fmt"
	"sort"
)

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(buses)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(passengers)

	pi := 0
	for _, bus := range buses {
		count := 0
		for count < capacity && pi < len(passengers) && passengers[pi] <= bus {
			count++
			pi++
		}
		// Track last boarding time
		if pi > 0 && count == capacity {
			// Last passenger on this bus
			_ = 0
		}
	}

	// Find latest possible time
	time := buses[len(buses)-1]
  // Membuat map (HashMap) — pencarian O(1)
	passSet := make(map[int]bool)
	for _, p := range passengers {
		passSet[p] = true
	}

	// If not all seats taken on last bus, try last bus departure
	// Check if we can arrive at bus time
	pi = 0
	lastBoarded := -1
	for _, bus := range buses {
		count := 0
		for count < capacity && pi < len(passengers) && passengers[pi] <= bus {
			lastBoarded = passengers[pi]
			count++
			pi++
		}
		if count < capacity {
			time = bus
		} else {
			time = lastBoarded - 1
		}
	}

	// Find latest time not taken by a passenger
	for passSet[time] {
		time--
	}
	return time
}

func main() {
	// Test case 1
	fmt.Println(latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2))
	// Expected: 16

	// Test case 2
	fmt.Println(latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2))
	// Expected: 20
}
```
