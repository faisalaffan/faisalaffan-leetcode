# 3829 — Design Ride Sharing System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func ConstructorRideSharing() RideSharingSystem`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** O(1) per operation  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3829: Design Ride Sharing System
// https://leetcode.com/problems/design-ride-sharing-system/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(N)
// Approach: Use two FIFO queues for riders and drivers with timestamp ordering.

import "fmt"

type RideSharingSystem struct {
	t       int
	riders  [][2]int // (timestamp, riderId)
	drivers [][2]int // (timestamp, driverId)
	// rider timestamp lookup for cancel
	riderTS map[int]int
}

func ConstructorRideSharing() RideSharingSystem {
	return RideSharingSystem{
		riderTS: make(map[int]int),
	}
}

func (rs *RideSharingSystem) AddRider(riderId int) {
	rs.riderTS[riderId] = rs.t
	rs.riders = append(rs.riders, [2]int{rs.t, riderId})
	rs.t++
}

func (rs *RideSharingSystem) AddDriver(driverId int) {
	rs.drivers = append(rs.drivers, [2]int{rs.t, driverId})
	rs.t++
}

func (rs *RideSharingSystem) MatchDriverWithRider() [2]int {
	if len(rs.riders) == 0 || len(rs.drivers) == 0 {
		return [2]int{-1, -1}
	}
	driver := rs.drivers[0]
	rs.drivers = rs.drivers[1:]
	rider := rs.riders[0]
	rs.riders = rs.riders[1:]
	return [2]int{driver[1], rider[1]}
}

func (rs *RideSharingSystem) CancelRider(riderId int) {
	ts, ok := rs.riderTS[riderId]
	if !ok {
		return
	}
	for i, r := range rs.riders {
		if r[0] == ts && r[1] == riderId {
			rs.riders = append(rs.riders[:i], rs.riders[i+1:]...)
			break
		}
	}
	delete(rs.riderTS, riderId)
}

func main() {
	rs := ConstructorRideSharing()
	rs.AddRider(1)
	rs.AddDriver(10)
	rs.AddRider(2)
	rs.AddDriver(20)
	rs.AddRider(3)
	fmt.Println(rs.MatchDriverWithRider()) // Expected: [10 1]
	fmt.Println(rs.MatchDriverWithRider()) // Expected: [20 2]
	rs.CancelRider(3)
	fmt.Println(rs.MatchDriverWithRider()) // Expected: [-1 -1]
}
```
