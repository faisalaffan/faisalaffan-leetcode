# 3100 — Water Bottles Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxBottlesDrunk(numBottles int, numExchange int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3100: Water Bottles II
// https://leetcode.com/problems/water-bottles-ii/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles

	for empty >= numExchange {
		empty -= numExchange
		numExchange++
		total++
		empty++
	}

	return total
}

func main() {
	fmt.Println(maxBottlesDrunk(13, 6)) // Expected: 15
	fmt.Println(maxBottlesDrunk(10, 3)) // Expected: 13
}
```
