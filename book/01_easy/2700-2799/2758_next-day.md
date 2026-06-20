# 2758 — Next Day

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NextDay(date string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2758: Next Day
// https://leetcode.com/problems/next-day/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JS Date problem adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(NextDay("2024-03-01"))
	fmt.Println(NextDay("2024-12-31"))
}

func NextDay(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	return t.AddDate(0, 0, 1).Format("2006-01-02")
}
```
