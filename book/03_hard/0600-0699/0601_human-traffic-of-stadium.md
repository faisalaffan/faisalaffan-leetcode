# 0601 — Human Traffic Of Stadium

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func humanTrafficOfStadium(records []StadiumRecord) []StadiumRecord`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(N log N) for sorting, Space: O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #601: Human Traffic of Stadium
// https://leetcode.com/problems/human-traffic-of-stadium/
// Difficulty: Hard
//
// Find all rows where 3 or more consecutive rows have >= 100 people.
// Return them ordered by visit_date (ascending).

// StadiumRecord represents a row in the Stadium table.
type StadiumRecord struct {
	ID        int
	VisitDate string // "YYYY-MM-DD"
	People    int
}

// humanTrafficOfStadium finds all records belonging to consecutive groups of >= 3
// with people >= 100.
// Time: O(N log N) for sorting, Space: O(N)
func humanTrafficOfStadium(records []StadiumRecord) []StadiumRecord {
	if len(records) == 0 {
		return nil
	}

	// Sort by ID (which correlates with visit_date).
  // Custom sort
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// Mark qualifying rows (people >= 100).
	n := len(records)
	qualifies := make([]bool, n)
	for i, r := range records {
		qualifies[i] = r.People >= 100
	}

	// Find runs of length >= 3.
	canInclude := make([]bool, n)
	i := 0
	for i < n {
		if !qualifies[i] {
			i++
			continue
		}
		// Start of a qualifying run.
		start := i
		for i < n && qualifies[i] {
			i++
		}
		runLen := i - start
		if runLen >= 3 {
			for j := start; j < i; j++ {
				canInclude[j] = true
			}
		}
	}

	var result []StadiumRecord
	for j, include := range canInclude {
		if include {
			result = append(result, records[j])
		}
	}
	return result
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0601 Human Traffic of Stadium ===")

	records := []StadiumRecord{
		{ID: 1, VisitDate: "2017-01-01", People: 10},
		{ID: 2, VisitDate: "2017-01-02", People: 109},
		{ID: 3, VisitDate: "2017-01-03", People: 150},
		{ID: 4, VisitDate: "2017-01-04", People: 99},
		{ID: 5, VisitDate: "2017-01-05", People: 145},
		{ID: 6, VisitDate: "2017-01-06", People: 1455},
		{ID: 7, VisitDate: "2017-01-07", People: 199},
		{ID: 8, VisitDate: "2017-01-08", People: 188},
	}

	fmt.Println("Stadium records:")
	for _, r := range records {
		fmt.Printf("  ID %d, %s, %d people\n", r.ID, r.VisitDate, r.People)
	}

	results := humanTrafficOfStadium(records)
	fmt.Println("\nRecords with >=3 consecutive days with >=100 people:")
	for _, r := range results {
		fmt.Printf("  ID %d, %s, %d people\n", r.ID, r.VisitDate, r.People)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Short run (length 2 only).
	recs2 := []StadiumRecord{
		{ID: 1, VisitDate: "2020-01-01", People: 150},
		{ID: 2, VisitDate: "2020-01-02", People: 200},
	}
	r2 := humanTrafficOfStadium(recs2)
	fmt.Println("Run of 2:", len(r2), "results (expected 0)")

	// All >= 100, run of 4.
	recs3 := []StadiumRecord{
		{ID: 1, VisitDate: "2020-01-01", People: 100},
		{ID: 2, VisitDate: "2020-01-02", People: 100},
		{ID: 3, VisitDate: "2020-01-03", People: 100},
		{ID: 4, VisitDate: "2020-01-04", People: 100},
	}
	fmt.Println("All >= 100, run of 4:")
	for _, r := range humanTrafficOfStadium(recs3) {
		fmt.Printf("  ID %d, %s\n", r.ID, r.VisitDate)
	}

	// Two separate runs >= 3.
	recs4 := []StadiumRecord{
		{ID: 1, VisitDate: "2020-01-01", People: 100},
		{ID: 2, VisitDate: "2020-01-02", People: 100},
		{ID: 3, VisitDate: "2020-01-03", People: 100},
		{ID: 4, VisitDate: "2020-01-04", People: 50},
		{ID: 5, VisitDate: "2020-01-05", People: 100},
		{ID: 6, VisitDate: "2020-01-06", People: 100},
		{ID: 7, VisitDate: "2020-01-07", People: 100},
	}
	fmt.Println("Two separate runs:")
	for _, r := range humanTrafficOfStadium(recs4) {
		fmt.Printf("  ID %d, %s\n", r.ID, r.VisitDate)
	}

	// Empty.
	fmt.Println("Empty:", len(humanTrafficOfStadium(nil)))
}
```
