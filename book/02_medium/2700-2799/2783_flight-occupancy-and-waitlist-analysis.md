# 2783 — Flight Occupancy And Waitlist Analysis

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FlightOccupancyAndWaitlistAnalysis(flights []FlightStatus) []FlightStatus
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
