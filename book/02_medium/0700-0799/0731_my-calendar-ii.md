# 0731 — My Calendar Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConstructorCalendar2() MyCalendarTwo`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2) per booking  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #731: My Calendar II
// https://leetcode.com/problems/my-calendar-ii/
// Difficulty: Medium
// Time: O(n^2) per booking
// Space: O(n)

import "fmt"

func main() {
	cal := ConstructorCalendar2()
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(50, 60))
	fmt.Println(cal.Book(10, 40))
	fmt.Println(cal.Book(5, 15))
	fmt.Println(cal.Book(5, 10))
	fmt.Println(cal.Book(25, 55))
}

type MyCalendarTwo struct {
	books    [][2]int
	overlaps [][2]int
}

func ConstructorCalendar2() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start int, end int) bool {
	for _, o := range c.overlaps {
		if max(o[0], start) < min(o[1], end) {
			return false
		}
	}
	for _, b := range c.books {
		if max(b[0], start) < min(b[1], end) {
			c.overlaps = append(c.overlaps, [2]int{max(b[0], start), min(b[1], end)})
		}
	}
	c.books = append(c.books, [2]int{start, end})
	return true
}
```
