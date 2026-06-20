# 1225 — Report Contiguous Dates

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func getContiguousPeriods(data []DailyStatus) []DateRange
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1225: Report Contiguous Dates
// https://leetcode.com/problems/report-contiguous-dates/
// Difficulty: Hard [Paid]
//
// Given a table of dates with a "state" (succeeded or failed), report
// contiguous date ranges (period_state, start_date, end_date) for each
// state, grouped by consecutive runs.

import (
	"fmt"
	"sort"
)

// DailyStatus represents one day's status.
type DailyStatus struct {
	Date  string // "YYYY-MM-DD"
	State string // "succeeded" or "failed"
}

// DateRange represents a contiguous period of the same state.
type DateRange struct {
	State     string // "succeeded" or "failed"
	StartDate string
	EndDate   string
}

func main() {
	data := []DailyStatus{
		{"2019-01-01", "succeeded"},
		{"2019-01-02", "succeeded"},
		{"2019-01-03", "succeeded"},
		{"2019-01-04", "failed"},
		{"2019-01-05", "failed"},
		{"2019-01-06", "succeeded"},
	}

	periods := getContiguousPeriods(data)
	for _, p := range periods {
		fmt.Printf("%s | %s | %s\n", p.State, p.StartDate, p.EndDate)
	}
}

// getContiguousPeriods groups consecutive dates of the same state into ranges.
func getContiguousPeriods(data []DailyStatus) []DateRange {
	if len(data) == 0 {
		return nil
	}

	// Sort by date
  // Custom sort dengan comparator
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date < data[j].Date
	})

	var periods []DateRange
	start := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(data); i++ {
		// Check if the next day exists and has a different state
		if i+1 < len(data) && data[i+1].State == data[i].State {
			continue
		}
		// End of a contiguous block
		periods = append(periods, DateRange{
			State:     data[i].State,
			StartDate: data[start].Date,
			EndDate:   data[i].Date,
		})
		start = i + 1
	}

	return periods
}
```
