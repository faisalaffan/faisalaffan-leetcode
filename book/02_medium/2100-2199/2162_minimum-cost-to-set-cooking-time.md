# 2162 — Minimum Cost To Set Cooking Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2162: Minimum Cost to Set Cooking Time
// https://leetcode.com/problems/minimum-cost-to-set-cooking-time/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
)

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	// Try all valid representations: minutes:seconds
	cost := int(1e9)

	// mm:ss where 0 <= mm <= 99, 0 <= ss <= 99
	for m := 0; m <= 99; m++ {
		for s := 0; s <= 99; s++ {
			if m*60+s == targetSeconds {
				digits := []int{}
				if m >= 10 {
					digits = append(digits, m/10)
				}
				digits = append(digits, m%10)
				if s < 10 && len(digits) > 0 {
					// If we have minutes, seconds always need 2 digits
					digits = append(digits, s/10)
				} else if s >= 10 {
					digits = append(digits, s/10)
				}
				digits = append(digits, s%10)

				cur := startAt
				total := 0
				for _, d := range digits {
					if cur != d {
						total += moveCost
						cur = d
					}
					total += pushCost
				}
				if total < cost {
					cost = total
				}
			}
		}
	}

	return cost
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCostSetTime(1, 2, 1, 600))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", minCostSetTime(0, 1, 2, 76))
	// Expected: 6

	// Test case 3
	fmt.Println("Test 3:", minCostSetTime(9, 100, 1, 600))
	// Expected: 8
}
```
