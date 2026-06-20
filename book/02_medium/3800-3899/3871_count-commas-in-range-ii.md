# 3871 — Count Commas In Range Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountCommasInRangeIi(n int64) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log N)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3871: Count Commas in Range II
// https://leetcode.com/problems/count-commas-in-range-ii/
// Difficulty: Medium
// Time: O(log N) | Space: O(1)
// Approach: Iterate powers of 1000 starting from 1000. For each power x <= n,
// add n - x + 1 (numbers that gain a comma at this magnitude).

import "fmt"

func CountCommasInRangeIi(n int64) int64 {
	var ans int64 = 0
	for x := int64(1000); x <= n; x *= 1000 {
		ans += n - x + 1
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountCommasInRangeIi(1002)) // Expected: 3

	// Example 2
	fmt.Println(CountCommasInRangeIi(998)) // Expected: 0

	// Extra
	fmt.Println(CountCommasInRangeIi(1000000)) // Expected: 999001 + 1 = 999002
}
```
