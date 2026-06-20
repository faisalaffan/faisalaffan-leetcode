# 2793 — Status Of Flight Tickets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func statusOfFlightTickets(flights [][]int, passengers [][]int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2793: Status of Flight Tickets
// https://leetcode.com/problems/status-of-flight-tickets/
// Difficulty: Hard [Paid]
//
// Given flights (flight_id, capacity) and passengers (passenger_id, flight_id, booking_time),
// determine each passenger's ticket status. Within each flight, passengers are sorted by
// booking_time; the first `capacity` are "Confirmed", the rest are "Waitlist".
// O(N log N + F log F) time, O(N + F) space.

import (
	"fmt"
	"sort"
)

func statusOfFlightTickets(flights [][]int, passengers [][]int) []string {
  // Membuat map (HashMap) — pencarian O(1)
	capMap := make(map[int]int)
	for _, f := range flights {
		capMap[f[0]] = f[1]
	}

  // Membuat map (HashMap) — pencarian O(1)
	byFlight := make(map[int][]int)
	for i, p := range passengers {
		byFlight[p[1]] = append(byFlight[p[1]], i)
	}

	res := make([]string, len(passengers))
	for fid, indices := range byFlight {
  // Custom sort dengan comparator
		sort.Slice(indices, func(i, j int) bool {
			return passengers[indices[i]][2] < passengers[indices[j]][2]
		})
		cap := capMap[fid]
		for i, idx := range indices {
			if i < cap {
				res[idx] = "Confirmed"
			} else {
				res[idx] = "Waitlist"
			}
		}
	}
	return res
}

func main() {
	// Problem example
	flights := [][]int{{1, 2}, {2, 2}, {3, 1}}
	passengers := [][]int{
		{101, 1, 202307101630},
		{102, 1, 202307101745},
		{103, 1, 202307101200},
		{104, 2, 202307051323},
		{105, 2, 202307050900},
		{106, 3, 202307081110},
		{107, 3, 202307080910},
	}
	for _, s := range statusOfFlightTickets(flights, passengers) {
		fmt.Println(s)
	}

	// Single flight, single passenger - confirmed
	fmt.Println(statusOfFlightTickets(
		[][]int{{1, 1}},
		[][]int{{1, 1, 100}}))

	// Single flight, over capacity - first confirmed, rest waitlist
	fmt.Println(statusOfFlightTickets(
		[][]int{{1, 1}},
		[][]int{{1, 1, 300}, {2, 1, 200}, {3, 1, 100}}))

	// All confirmed - capacity >= passenger count
	fmt.Println(statusOfFlightTickets(
		[][]int{{1, 5}},
		[][]int{{1, 1, 300}, {2, 1, 200}, {3, 1, 100}}))

	// Multiple flights with varying capacities
	fmt.Println(statusOfFlightTickets(
		[][]int{{1, 1}, {2, 2}},
		[][]int{{1, 2, 100}, {2, 2, 200}, {3, 1, 300}, {4, 2, 50}}))

	// Empty flights (should not happen per constraints)
	fmt.Println(statusOfFlightTickets([][]int{}, [][]int{}))

	// Flight with zero capacity - all waitlist
	fmt.Println(statusOfFlightTickets(
		[][]int{{1, 0}},
		[][]int{{1, 1, 100}, {2, 1, 200}}))
}
```
