# 3214 — Year On Year Growth Rate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func yearOnYearGrowthRate(records []Record) []GrowthResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3214: Year on Year Growth Rate
// https://leetcode.com/problems/year-on-year-growth-rate/
// Difficulty: Hard [Paid]
//
// Given a list of (year, month, value) records, compute the Year-over-Year
// growth rate for each period as a percentage:
//
//   YoY = ((current - prev_year_same_period) / prev_year_same_period) * 100
//
// If no record exists for the same period in the previous year, omit that row.
//
// Output: sorted by year, month with growth rates rounded to 2 decimal places.

import (
	"fmt"
	"math"
	"sort"
)

type Record struct {
	Year  int
	Month int
	Value float64
}

type GrowthResult struct {
	Year       int
	Month      int
	GrowthRate float64 // percentage, e.g. 10.5 means 10.5% growth
}

func yearOnYearGrowthRate(records []Record) []GrowthResult {
	// Build lookup: (year, month) -> value
  // Membuat map (HashMap) — pencarian O(1)
	lookup := make(map[[2]int]float64)
  // Alokasi slice integer
	yearMonthSet := make([][2]int, 0)
	for _, r := range records {
		key := [2]int{r.Year, r.Month}
		if _, exists := lookup[key]; !exists {
			yearMonthSet = append(yearMonthSet, key)
		}
		lookup[key] = r.Value
	}

  // Custom sort dengan comparator
	sort.Slice(yearMonthSet, func(i, j int) bool {
		if yearMonthSet[i][0] != yearMonthSet[j][0] {
			return yearMonthSet[i][0] < yearMonthSet[j][0]
		}
		return yearMonthSet[i][1] < yearMonthSet[j][1]
	})

	ans := make([]GrowthResult, 0)
	for _, ym := range yearMonthSet {
		year, month := ym[0], ym[1]
		prevKey := [2]int{year - 1, month}
		currVal, currOK := lookup[ym]
		prevVal, prevOK := lookup[prevKey]
		if currOK && prevOK && prevVal != 0 {
			growth := ((currVal - prevVal) / prevVal) * 100
			// Round to 2 decimal places.
			growth = math.Round(growth*100) / 100
			ans = append(ans, GrowthResult{Year: year, Month: month, GrowthRate: growth})
		}
	}
	return ans
}

func main() {
	records := []Record{
		{2020, 1, 100},
		{2020, 2, 120},
		{2021, 1, 110},
		{2021, 2, 130},
	}
	res := yearOnYearGrowthRate(records)
	for _, r := range res {
		fmt.Printf("%d-%d: %.2f%%\n", r.Year, r.Month, r.GrowthRate)
	}
}
```
