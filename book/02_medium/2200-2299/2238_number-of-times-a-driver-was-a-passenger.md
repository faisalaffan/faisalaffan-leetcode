# 2238 — Number Of Times A Driver Was A Passenger

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPassengers(rides [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2238: Number of Times a Driver Was a Passenger
// https://leetcode.com/problems/number-of-times-a-driver-was-a-passenger/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func countPassengers(rides [][]int) []int {
	// Each ride: [driver_id, passenger_id]
  // Membuat map (HashMap) — pencarian O(1)
	passengerCount := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	drivers := make(map[int]bool)

	for _, ride := range rides {
		driver, passenger := ride[0], ride[1]
		drivers[driver] = true
		passengerCount[passenger]++
	}

	maxID := 0
	for id := range drivers {
		if id > maxID {
			maxID = id
		}
	}

  // Alokasi slice integer
	result := make([]int, maxID+1)
	for id := 1; id <= maxID; id++ {
		if drivers[id] {
			result[id] = passengerCount[id]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(countPassengers([][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Expected: [0, 1, 1, 1]

	// Test case 2
	fmt.Println(countPassengers([][]int{{1, 2}, {1, 3}, {1, 4}}))
	// Expected: [0, 0]
}
```
