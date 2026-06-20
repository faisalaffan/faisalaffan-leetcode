# 1360 — Number Of Days Between Two Dates

## Deskripsi

**Soal:** [1360. Number Of Days Between Two Dates](https://leetcode.com/problems/number-of-days-between-two-dates/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func daysBetweenDates(date1 string, date2 string) int`

## Solusi Go

```go
package main

// LeetCode #1360: Number of Days Between Two Dates
// https://leetcode.com/problems/number-of-days-between-two-dates/
// Difficulty: Easy
//
// LeetCode submission: func daysBetweenDates(date1 string, date2 string) int

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(NumberOfDaysBetweenTwoDates("2019-06-29", "2019-06-30")) // 1
	fmt.Println(NumberOfDaysBetweenTwoDates("2020-01-15", "2019-12-31")) // 15
}

// Time: O(1), Space: O(1)
func NumberOfDaysBetweenTwoDates(date1 string, date2 string) int {
	format := "2006-01-02"
	d1, _ := time.Parse(format, date1)
	d2, _ := time.Parse(format, date2)
	diff := d2.Sub(d1)
	if diff < 0 {
		diff = -diff
	}
	return int(diff.Hours() / 24)
}
```
