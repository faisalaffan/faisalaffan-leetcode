# 2739 — Total Distance Traveled

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func TotalDistanceTraveled(mainTank int, additionalTank int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2739: Total Distance Traveled
// https://leetcode.com/problems/total-distance-traveled/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalDistanceTraveled(5, 10))
	fmt.Println(TotalDistanceTraveled(1, 2))
}

func TotalDistanceTraveled(mainTank int, additionalTank int) int {
	total := 0
	for mainTank > 0 {
		if mainTank >= 5 {
			mainTank -= 5
			total += 50
			if additionalTank > 0 {
				additionalTank--
				mainTank++
			}
		} else {
			total += mainTank * 10
			mainTank = 0
		}
	}
	return total
}
```
