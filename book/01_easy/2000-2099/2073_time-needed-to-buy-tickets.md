# 2073 — Time Needed To Buy Tickets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func TimeNeededToBuyTickets(tickets []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2073: Time Needed to Buy Tickets
// https://leetcode.com/problems/time-needed-to-buy-tickets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TimeNeededToBuyTickets([]int{2, 3, 2}, 2)) // 6
	fmt.Println(TimeNeededToBuyTickets([]int{5, 1, 1, 1}, 0)) // 8
}

// Time: O(n), Space: O(1)
func TimeNeededToBuyTickets(tickets []int, k int) int {
	time := 0
	for i, t := range tickets {
		if i <= k {
			if t <= tickets[k] {
				time += t
			} else {
				time += tickets[k]
			}
		} else {
			if t < tickets[k] {
				time += t
			} else {
				time += tickets[k] - 1
			}
		}
	}
	return time
}
```
