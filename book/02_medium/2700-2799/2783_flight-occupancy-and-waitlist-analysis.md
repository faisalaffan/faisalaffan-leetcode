# 2783 — Flight Occupancy And Waitlist Analysis

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FlightOccupancyAndWaitlistAnalysis(flights []FlightStatus) []FlightStatus`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2783: Flight Occupancy and Waitlist Analysis
// https://leetcode.com/problems/flight-occupancy-and-waitlist-analysis/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type FlightStatus struct {
	Capacity      int
	Booked        int
	Waitlisted    int
}

func FlightOccupancyAndWaitlistAnalysis(flights []FlightStatus) []FlightStatus {
	results := make([]FlightStatus, len(flights))
	for i, f := range flights {
		available := f.Capacity - f.Booked
		if available < 0 {
			available = 0
		}
		// Waitlisted passengers fill available spots
		canBoard := f.Waitlisted
		if canBoard > available {
			canBoard = available
		}
		results[i] = FlightStatus{
			Capacity:      f.Capacity,
			Booked:        f.Booked + canBoard,
			Waitlisted:    f.Waitlisted - canBoard,
		}
	}
	return results
}

func main() {
	flights := []FlightStatus{
		{Capacity: 100, Booked: 95, Waitlisted: 10},
		{Capacity: 50, Booked: 50, Waitlisted: 5},
	}
	fmt.Println(FlightOccupancyAndWaitlistAnalysis(flights))
}
```
