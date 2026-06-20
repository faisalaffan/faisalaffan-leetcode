# 3570 — Find Books With No Available Copies

## Deskripsi

**Soal:** [3570. Find Books With No Available Copies](https://leetcode.com/problems/find-books-with-no-available-copies/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3570: Find Books with No Available Copies
// https://leetcode.com/problems/find-books-with-no-available-copies/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	books := []LibBook{
		{BookID: 1, Title: "Book A", AvailableCopies: 0},
		{BookID: 2, Title: "Book B", AvailableCopies: 3},
		{BookID: 3, Title: "Book C", AvailableCopies: 0},
	}
	result := FindBooksWithNoAvailableCopies(books)
	for _, b := range result {
		fmt.Printf("%d: %s\n", b.BookID, b.Title)
	}
}

// LibBook represents a library book.
type LibBook struct {
	BookID          int
	Title           string
	AvailableCopies int
}

// FindBooksWithNoAvailableCopies returns books with zero available copies, sorted by book_id.
// Time: O(n log n). Space: O(n).
func FindBooksWithNoAvailableCopies(books []LibBook) []LibBook {
	result := []LibBook{}
	for _, b := range books {
		if b.AvailableCopies == 0 {
			result = append(result, b)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].BookID < result[j].BookID
	})
	return result
}
```
