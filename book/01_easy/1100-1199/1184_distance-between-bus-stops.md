# 1184 — Distance Between Bus Stops

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func distanceBetweenBusStops(distance []int, start, destination int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1184: Distance Between Bus Stops
// https://leetcode.com/problems/distance-between-bus-stops/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1)) // 1
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2)) // 3
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3)) // 4
}

// LeetCode submission: distanceBetweenBusStops
func distanceBetweenBusStops(distance []int, start, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	forward := 0
	for i := start; i < destination; i++ {
		forward += distance[i]
	}
	total := 0
	for _, d := range distance {
		total += d
	}
	backward := total - forward
	if backward < forward {
		return backward
	}
	return forward
}
```
