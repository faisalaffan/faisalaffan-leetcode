# 1321 — Restaurant Growth

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func restaurantGrowth(customers []struct { visitedOn string amount int }) []avgResult`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sliding Window, Sorting

**Waktu:** O(n log n) due to sorting  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1321: Restaurant Growth
// https://leetcode.com/problems/restaurant-growth/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	customers := []struct {
		visitedOn string
		amount    int
	}{
		{"2020-01-01", 10},
		{"2020-01-02", 10},
		{"2020-01-03", 10},
		{"2020-01-04", 10},
		{"2020-01-05", 10},
		{"2020-01-06", 10},
		{"2020-01-07", 10},
		{"2020-01-08", 20},
		{"2020-01-09", 20},
		{"2020-01-10", 20},
	}

	result := restaurantGrowth(customers)
	for _, r := range result {
		fmt.Printf("%s %.2f\n", r.date, r.avg)
	}
}

type avgResult struct {
	date string
	avg  float64
}

// Time: O(n log n) due to sorting
// Space: O(n)
func restaurantGrowth(customers []struct {
	visitedOn string
	amount    int
}) []avgResult {
	// Group by date and sum amounts
	type daySum struct {
		date string
		total int
		count int
	}

  // HashMap: O(1) lookup
	dateMap := make(map[string]*daySum)
	for _, c := range customers {
		if _, ok := dateMap[c.visitedOn]; !ok {
			dateMap[c.visitedOn] = &daySum{date: c.visitedOn}
		}
		dateMap[c.visitedOn].total += c.amount
		dateMap[c.visitedOn].count++
	}

	dates := make([]string, 0, len(dateMap))
	for d := range dateMap {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	// Sliding window of 7 days
	var result []avgResult
  // Alokasi slice
	window := make([]int, 0, 7)

	for _, d := range dates {
		ds := dateMap[d]
		window = append(window, ds.total)
		if len(window) > 7 {
			window = window[1:]
		}
		if len(window) == 7 {
			sum := 0
			for _, v := range window {
				sum += v
			}
			result = append(result, avgResult{d, float64(sum) / 7.0})
		}
	}

	return result
}
```
