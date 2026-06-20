# 1518 — Water Bottles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numWaterBottles(numBottles int, numExchange int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1518: Water Bottles
// https://leetcode.com/problems/water-bottles/
// Difficulty: Easy
//
// LeetCode submission: func numWaterBottles(numBottles int, numExchange int) int

import "fmt"

func main() {
	fmt.Println(WaterBottles(9, 3))  // 13
	fmt.Println(WaterBottles(15, 4)) // 19
	fmt.Println(WaterBottles(5, 5))  // 6
}

// Time: O(log n), Space: O(1)
func WaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles
	for empty >= numExchange {
		newBottles := empty / numExchange
		total += newBottles
		empty = newBottles + empty%numExchange
	}
	return total
}
```
