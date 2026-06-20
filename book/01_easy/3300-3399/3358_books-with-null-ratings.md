# 3358 — Books With Null Ratings

## Deskripsi

**Soal:** [3358. Books With Null Ratings](https://leetcode.com/problems/books-with-null-ratings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #3358: Books with NULL Ratings
// https://leetcode.com/problems/books-with-null-ratings/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	books := []Book{
		{BookID: 1, Title: "Book A", Author: "Author X", PublishedYear: 2020, Rating: nil},
		{BookID: 2, Title: "Book B", Author: "Author Y", PublishedYear: 2021, Rating: ptr(4)},
		{BookID: 3, Title: "Book C", Author: "Author Z", PublishedYear: 2019, Rating: nil},
	}
	result := BooksWithNullRatings(books)
	for _, b := range result {
		fmt.Println(b)
	}
}

func ptr(i int) *int { return &i }

// Book represents a book with optional rating.
type Book struct {
	BookID        int
	Title         string
	Author        string
	PublishedYear int
	Rating        *int
}

// BooksWithNullRatings returns books that have NULL ratings, sorted by book_id.
// Time: O(n log n). Space: O(n).
func BooksWithNullRatings(books []Book) []Book {
	result := []Book{}
	for _, b := range books {
		if b.Rating == nil {
			result = append(result, b)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].BookID < result[j].BookID
	})
	return result
}
```
