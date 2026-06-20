# 3308 — Find Top Performing Driver

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func topPerformingDriver(drivers []Driver, vehicles []Vehicle, trips []Trip) []FuelRank`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(d + v + t) Space: O(d + v)  |  **Ruang:** O(d + v)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3308: Find Top Performing Driver
// https://leetcode.com/problems/find-top-performing-driver/
// Difficulty: Medium
// Time: O(d + v + t) Space: O(d + v)

import (
	"fmt"
	"sort"
)

func main() {
	drivers := []Driver{{1, "Gasoline"}, {2, "Gasoline"}, {3, "Diesel"}}
	vehicles := []Vehicle{{1, 1}, {2, 2}, {3, 1}}
	trips := []Trip{{1, 1, 4.5, 100}, {2, 1, 4.0, 200}, {3, 2, 5.0, 150}, {4, 3, 3.5, 50}}
	fmt.Println(topPerformingDriver(drivers, vehicles, trips))

	drivers2 := []Driver{{1, "Electric"}, {2, "Electric"}}
	vehicles2 := []Vehicle{{1, 1}, {2, 2}}
	trips2 := []Trip{{1, 1, 5.0, 100}, {2, 2, 5.0, 80}}
	fmt.Println(topPerformingDriver(drivers2, vehicles2, trips2))
}

type Driver struct {
	ID       int
	FuelType string
}

type Vehicle struct {
	ID       int
	DriverID int
}

type Trip struct {
	ID       int
	VehicleID int
	Rating   float64
	Distance int
}

type FuelRank struct {
	FuelType   string
	DriverID   int
	Rating     float64
	Distance   int
	Accidents  int
}

func topPerformingDriver(drivers []Driver, vehicles []Vehicle, trips []Trip) []FuelRank {
	// Build driver -> fuel type map
  // HashMap: O(1) lookup
	driverFuel := make(map[int]string)
	for _, d := range drivers {
		driverFuel[d.ID] = d.FuelType
	}

	// Build vehicle -> driver map
  // HashMap: O(1) lookup
	vehicleDriver := make(map[int]int)
  // HashMap: O(1) lookup
	driverVehicles := make(map[int][]int)
	for _, v := range vehicles {
		vehicleDriver[v.ID] = v.DriverID
		driverVehicles[v.DriverID] = append(driverVehicles[v.DriverID], v.ID)
	}

	// Group trips by vehicle, then by driver
	type stats struct {
		sumRating float64
		count     int
		distance  int
	}
  // HashMap: O(1) lookup
	driverStats := make(map[int]*stats)
	for _, t := range trips {
		dID := vehicleDriver[t.VehicleID]
		if driverStats[dID] == nil {
			driverStats[dID] = &stats{}
		}
		driverStats[dID].sumRating += t.Rating
		driverStats[dID].count++
		driverStats[dID].distance += t.Distance
	}

	type candidate struct {
		fuelType  string
		driverID  int
		rating    float64
		distance  int
		accidents int
	}

  // HashMap: O(1) lookup
	fuelCands := make(map[string][]candidate)
	for dID, s := range driverStats {
		avgRating := s.sumRating / float64(s.count)
		avgRating = float64(int(avgRating*100)) / 100 // round to 2 decimals
		ft := driverFuel[dID]
		fuelCands[ft] = append(fuelCands[ft], candidate{
			fuelType: ft, driverID: dID,
			rating: avgRating, distance: s.distance,
		})
	}

	var result []FuelRank
	for ft, cands := range fuelCands {
  // Custom sort
		sort.Slice(cands, func(i, j int) bool {
			if cands[i].rating != cands[j].rating {
				return cands[i].rating > cands[j].rating
			}
			if cands[i].distance != cands[j].distance {
				return cands[i].distance > cands[j].distance
			}
			return cands[i].driverID < cands[j].driverID
		})
		best := cands[0]
		result = append(result, FuelRank{
			FuelType: ft, DriverID: best.driverID,
			Rating: best.rating, Distance: best.distance,
		})
	}

  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		return result[i].FuelType < result[j].FuelType
	})
	return result
}
```
