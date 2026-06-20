# 2651 — Calculate Delayed Arrival Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CalculateDelayedArrivalTime(arrivalTime int, delayedTime int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2651: Calculate Delayed Arrival Time
// https://leetcode.com/problems/calculate-delayed-arrival-time/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CalculateDelayedArrivalTime(15, 5))
	fmt.Println(CalculateDelayedArrivalTime(23, 10))
}

func CalculateDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
```
