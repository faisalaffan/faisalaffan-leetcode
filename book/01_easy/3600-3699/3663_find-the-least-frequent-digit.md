# 3663 — Find The Least Frequent Digit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindTheLeastFrequentDigit(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n) - number of digits  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3663: Find The Least Frequent Digit
// https://leetcode.com/problems/find-the-least-frequent-digit/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindTheLeastFrequentDigit(1553322))
	fmt.Println(FindTheLeastFrequentDigit(723344511))
}

// Time: O(log n) - number of digits
// Space: O(1)
func FindTheLeastFrequentDigit(n int) int {
	cnt := [10]int{}
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}

	minCnt := math.MaxInt
	ans := 0
	for d := 0; d <= 9; d++ {
		if cnt[d] > 0 && (cnt[d] < minCnt || (cnt[d] == minCnt && d < ans)) {
			minCnt = cnt[d]
			ans = d
		}
	}
	return ans
}
```
