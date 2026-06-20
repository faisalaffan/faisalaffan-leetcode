# 2805 — Custom Interval

## Deskripsi

**Soal:** [2805. Custom Interval](https://leetcode.com/problems/custom-interval/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func CustomInterval(fn func() , interval time.Duration, times int)`

## Solusi Go

```go
package main

// LeetCode #2805: Custom Interval
// https://leetcode.com/problems/custom-interval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"time"
)

func CustomInterval(fn func(), interval time.Duration, times int) {
	for i := 0; i < times; i++ {
		fn()
		time.Sleep(interval)
	}
}

func main() {
	count := 0
	fn := func() {
		count++
		fmt.Println("Executed:", count)
	}
	CustomInterval(fn, 10*time.Millisecond, 3)
	fmt.Println("Done")
}
```
