# 1154 — Day Of The Year

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func dayOfYear(date string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1154: Day of the Year
// https://leetcode.com/problems/day-of-the-year/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09")) // 9
	fmt.Println(dayOfYear("2019-02-10")) // 41
	fmt.Println(dayOfYear("2000-03-01")) // 61
}

// LeetCode submission: dayOfYear
func dayOfYear(date string) int {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leap {
		days[1] = 29
	}

	ans := 0
	for i := 0; i < month-1; i++ {
		ans += days[i]
	}
	return ans + day
}
```
