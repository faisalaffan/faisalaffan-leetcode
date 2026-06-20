# 1058 — Minimize Rounding Error To Meet Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimizeRoundingErrorToMeetTarget(prices []string, target int) string
```

> **💡 Hint:** For each price, floor and ceil. Compute min total error via DP.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * target) effectively O(n) after sorting diffs  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1058: Minimize Rounding Error to Meet Target
// https://leetcode.com/problems/minimize-rounding-error-to-meet-target/
// Difficulty: Medium
//
// Approach: For each price, floor and ceil. Compute min total error via DP.
// Time: O(n * target) effectively O(n) after sorting diffs
// Space: O(n)

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(minimizeRoundingErrorToMeetTarget([]string{"0.700", "2.800", "4.900"}, 8)) // "1.000"
	fmt.Println(minimizeRoundingErrorToMeetTarget([]string{"1.500", "2.500", "3.500"}, 10)) // "-1"
}

func minimizeRoundingErrorToMeetTarget(prices []string, target int) string {
	n := len(prices)
  // Alokasi slice integer
	floors := make([]int, n)
	diffs := make([]float64, n)
	floorSum := 0

	for i, p := range prices {
		f, _ := strconv.ParseFloat(p, 64)
		floor := int(math.Floor(f))
		floors[i] = floor
		floorSum += floor
		diffs[i] = f - float64(floor)
	}

	// We need to ceil (target - floorSum) items
	ceilCount := target - floorSum
	if ceilCount < 0 || ceilCount > n {
		return "-1"
	}

	// Sort diffs descending to ceil those with smallest rounding error
  // Custom sort dengan comparator
	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i] > diffs[j]
	})

	totalErr := 0.0
	for i := 0; i < ceilCount; i++ {
		totalErr += 1.0 - diffs[i]
	}
	for i := ceilCount; i < n; i++ {
		totalErr += diffs[i]
	}

	return fmt.Sprintf("%.3f", math.Round(totalErr*1000)/1000)
}
```
