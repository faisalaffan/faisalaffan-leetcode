# 2145 — Count The Hidden Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfArrays(differences []int, lower int, upper int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2145: Count the Hidden Sequences
// https://leetcode.com/problems/count-the-hidden-sequences/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfArrays(differences []int, lower int, upper int) int {
	cur := int64(0)
	minVal := int64(0)
	maxVal := int64(0)

	for _, d := range differences {
		cur += int64(d)
		if cur < minVal {
			minVal = cur
		}
		if cur > maxVal {
			maxVal = cur
		}
	}

	// Our sequence starts at some value x in [lower, upper]
	// All elements: x + prefix[i] must be in [lower, upper]
	// So: lower <= x + minVal AND x + maxVal <= upper
	// => x >= lower - minVal AND x <= upper - maxVal
	// => valid x in [max(lower, lower-minVal), min(upper, upper-maxVal)]

	low := int64(lower)
	high := int64(upper)
	minStart := low - minVal
	maxStart := high - maxVal

	if minStart > maxStart {
		return 0
	}

	start := minStart
	if start < low {
		start = low
	}
	end := maxStart
	if end > high {
		end = high
	}

	if start > end {
		return 0
	}
	return int(end - start + 1)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfArrays([]int{1, -3, 4}, 1, 6))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", numberOfArrays([]int{4, -7, 2}, 3, 6))
	// Expected: 0
}
```
