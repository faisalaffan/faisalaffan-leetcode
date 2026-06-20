# 0729 — My Calendar I

## Deskripsi

**Soal:** [0729. My Calendar I](https://leetcode.com/problems/my-calendar-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n) per booking  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #729: My Calendar I
// https://leetcode.com/problems/my-calendar-i/
// Difficulty: Medium
// Time: O(log n) per booking
// Space: O(n)

import "fmt"

func main() {
	cal := ConstructorCalendar()
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(15, 25))
	fmt.Println(cal.Book(20, 30))
}

type MyCalendar struct {
	books [][2]int
}

func ConstructorCalendar() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	for _, b := range c.books {
		if max(b[0], start) < min(b[1], end) {
			return false
		}
	}
	c.books = append(c.books, [2]int{start, end})
	return true
}
```
