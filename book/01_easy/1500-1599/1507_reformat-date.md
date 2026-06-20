# 1507 — Reformat Date

## Deskripsi

**Soal:** [1507. Reformat Date](https://leetcode.com/problems/reformat-date/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func reformatDate(date string) string`

## Solusi Go

```go
package main

// LeetCode #1507: Reformat Date
// https://leetcode.com/problems/reformat-date/
// Difficulty: Easy
//
// LeetCode submission: func reformatDate(date string) string

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReformatDate("20th Oct 2052")) // "2052-10-20"
	fmt.Println(ReformatDate("6th Jun 1933"))  // "1933-06-06"
	fmt.Println(ReformatDate("26th May 1960")) // "1960-05-26"
}

// Time: O(1), Space: O(1)
func ReformatDate(date string) string {
	months := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04",
		"May": "05", "Jun": "06", "Jul": "07", "Aug": "08",
		"Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
	}
	parts := strings.Split(date, " ")
	day := parts[0][:len(parts[0])-2]
	if len(day) == 1 {
		day = "0" + day
	}
	return parts[2] + "-" + months[parts[1]] + "-" + day
}
```
