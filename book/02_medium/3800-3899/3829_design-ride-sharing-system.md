# 3829 — Design Ride Sharing System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructorRideSharing() RideSharingSystem
```

> **💡 Hint:** Use two FIFO queues for riders and drivers with timestamp ordering.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, BFS

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
