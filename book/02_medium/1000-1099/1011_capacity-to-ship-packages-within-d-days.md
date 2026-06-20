# 1011 — Capacity To Ship Packages Within D Days

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func shipWithinDays(weights []int, days int) int
```

> **💡 Hint:** Binary search on capacity

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search

**Kompleksitas Waktu:** O(n * log(sum(weights)))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1011: Capacity To Ship Packages Within D Days
// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
// Difficulty: Medium
//
// Approach: Binary search on capacity
// Time: O(n * log(sum(weights)))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)) // 15
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))             // 6
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))                // 3
}

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if canShip(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canShip(weights []int, days int, capacity int) bool {
	dayCount := 1
	current := 0
	for _, w := range weights {
		if current+w > capacity {
			dayCount++
			current = w
			if dayCount > days {
				return false
			}
		} else {
			current += w
		}
	}
	return true
}
```
