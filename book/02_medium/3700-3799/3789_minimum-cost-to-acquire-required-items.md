# 3789 — Minimum Cost To Acquire Required Items

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minimumCostToAcquireRequiredItems(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3789: Minimum Cost to Acquire Required Items
// https://leetcode.com/problems/minimum-cost-to-acquire-required-items/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumCostToAcquireRequiredItems(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	a := int64(need1)*int64(cost1) + int64(need2)*int64(cost2)
	b := int64(costBoth) * int64(max(need1, need2))
	mn := min(need1, need2)
	c := int64(costBoth)*int64(mn) + int64(need1-mn)*int64(cost1) + int64(need2-mn)*int64(cost2)

	ans := a
	if b < ans {
		ans = b
	}
	if c < ans {
		ans = c
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumCostToAcquireRequiredItems(3, 2, 1, 3, 2))
	fmt.Println(minimumCostToAcquireRequiredItems(5, 4, 15, 2, 3))
	fmt.Println(minimumCostToAcquireRequiredItems(10, 10, 5, 5, 5))
}
```
