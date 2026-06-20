# 1635 — Hopper Company Queries I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func hopperQueriesI(drivers []Driver, rides []Ride, accepted []AcceptedRide) [][3]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1635: Hopper Company Queries I
// https://leetcode.com/problems/hopper-company-queries-i/
// Difficulty: Hard [Paid] (SQL)
//
// For each month of 2020, report the number of active drivers by
// month-end and the number of accepted rides that month.
//
// Approach: Process drivers and rides data in Go to simulate the
// SQL query results.

import (
	"fmt"
	"time"
)

// Driver represents a driver
type Driver struct {
	DriverID int
	JoinDate time.Time
}

// Ride represents a ride request
type Ride struct {
	RideID      int
	UserID      int
	RequestedAt time.Time
}

// AcceptedRide represents an accepted ride
type AcceptedRide struct {
	RideID        int
	DriverID      int
	RideDistance  int
	RideDuration  int
}

func main() {
	// Example
	drivers := []Driver{
		{10, parseDate("2019-12-10")},
		{8, parseDate("2020-1-13")},
		{5, parseDate("2020-2-16")},
		{7, parseDate("2020-3-8")},
		{4, parseDate("2020-5-17")},
		{1, parseDate("2020-10-24")},
		{6, parseDate("2021-1-5")},
	}
	rides := []Ride{
		{6, 82, parseDate("2019-12-9")},
		{1, 17, parseDate("2020-1-2")},
		{10, 36, parseDate("2020-1-11")},
		{11, 98, parseDate("2020-1-19")},
		{12, 44, parseDate("2020-1-29")},
		{3, 1, parseDate("2020-1-17")},
		{9, 90, parseDate("2020-2-13")},
		{2, 50, parseDate("2020-2-22")},
	}
	accepted := []AcceptedRide{
		{2, 10, 63, 38},
		{13, 10, 58, 0},
		{7, 8, 51, 23},
		{3, 5, 50, 36},
		{4, 5, 58, 19},
		{11, 7, 53, 26},
		{1, 10, 50, 42},
		{5, 7, 55, 42},
		{12, 8, 51, 23},
		{6, 7, 56, 41},
	}

	fmt.Println(hopperQueriesI(drivers, rides, accepted))
}

func parseDate(s string) time.Time {
	t, err := time.Parse("2006-1-2", s)
	if err != nil {
		t, err = time.Parse("2006-01-02", s)
		if err != nil {
			panic(err)
		}
	}
	return t
}

func hopperQueriesI(drivers []Driver, rides []Ride, accepted []AcceptedRide) [][3]int {
	// Count drivers active by end of each month in 2020
  // HashMap: O(1) lookup
	acceptedRideSet := make(map[int]bool)
	for _, ar := range accepted {
		acceptedRideSet[ar.RideID] = true
	}

  // HashMap: O(1) lookup
	rideMonth := make(map[int]int) // rideID -> month
	for _, r := range rides {
		if r.RequestedAt.Year() == 2020 {
			rideMonth[r.RideID] = int(r.RequestedAt.Month())
		}
	}

  // Alokasi slice
	activeDrivers := make([]int, 13)
	for _, d := range drivers {
		joinYear, joinMonth := d.JoinDate.Year(), d.JoinDate.Month()
		if joinYear < 2020 {
			for m := 1; m <= 12; m++ {
				activeDrivers[m]++
			}
		} else if joinYear == 2020 {
			for m := int(joinMonth); m <= 12; m++ {
				activeDrivers[m]++
			}
		}
	}

  // Alokasi slice
	acceptedRides := make([]int, 13)
	for rideID, month := range rideMonth {
		if acceptedRideSet[rideID] {
			acceptedRides[month]++
		}
	}

  // Alokasi slice
	result := make([][3]int, 12)
	for m := 1; m <= 12; m++ {
		result[m-1] = [3]int{m, activeDrivers[m], acceptedRides[m]}
	}
	return result
}
```
