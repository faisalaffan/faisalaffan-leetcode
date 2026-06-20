# 3642 — Find Books With Polarized Opinions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findBooksWithPolarizedOpinions(books []book, sessions []readingSession) []bookResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3642: Find Books with Polarized Opinions
// https://leetcode.com/problems/find-books-with-polarized-opinions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type book struct {
	id     int
	title  string
	author string
	genre  string
	pages  int
}

type readingSession struct {
	bookID int
	rating int
}

type bookResult struct {
	bookID            int
	title             string
	author            string
	genre             string
	pages             int
	ratingSpread      int
	polarizationScore float64
}

func findBooksWithPolarizedOpinions(books []book, sessions []readingSession) []bookResult {
	// Group ratings by book
  // Membuat map (HashMap) — pencarian O(1)
	ratingsByBook := make(map[int][]int)
	for _, s := range sessions {
		ratingsByBook[s.bookID] = append(ratingsByBook[s.bookID], s.rating)
	}

  // Membuat map (HashMap) — pencarian O(1)
	bookMap := make(map[int]book)
	for _, b := range books {
		bookMap[b.id] = b
	}

	var results []bookResult

	for bookID, ratings := range ratingsByBook {
		if len(ratings) < 5 {
			continue
		}

		maxRating := ratings[0]
		minRating := ratings[0]
		highCount := 0
		lowCount := 0

		for _, r := range ratings {
			if r > maxRating {
				maxRating = r
			}
			if r < minRating {
				minRating = r
			}
			if r >= 4 {
				highCount++
			}
			if r <= 2 {
				lowCount++
			}
		}

		if highCount == 0 || lowCount == 0 {
			continue
		}

		extremeCount := highCount + lowCount
		polarizationScore := float64(extremeCount) / float64(len(ratings))

		if polarizationScore < 0.6 {
			continue
		}

		b, ok := bookMap[bookID]
		if !ok {
			continue
		}

		pScore := float64(int(polarizationScore*100+0.5)) / 100
		results = append(results, bookResult{
			bookID:            b.id,
			title:             b.title,
			author:            b.author,
			genre:             b.genre,
			pages:             b.pages,
			ratingSpread:      maxRating - minRating,
			polarizationScore: pScore,
		})
	}

  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		if results[i].polarizationScore != results[j].polarizationScore {
			return results[i].polarizationScore > results[j].polarizationScore
		}
		return results[i].title > results[j].title
	})

	return results
}

func main() {
	books := []book{
		{1, "Book A", "Author A", "Fiction", 300},
		{2, "Book B", "Author B", "Non-Fiction", 250},
		{3, "Book C", "Author C", "Sci-Fi", 400},
	}
	sessions := []readingSession{
		{1, 5}, {1, 5}, {1, 1}, {1, 5}, {1, 1}, {1, 4},
		{2, 3}, {2, 4}, {2, 3}, {2, 4}, {2, 3},
		{3, 1}, {3, 5}, {3, 1}, {3, 5}, {3, 1}, {3, 5},
	}
	results := findBooksWithPolarizedOpinions(books, sessions)
	for _, r := range results {
		fmt.Printf("Book %d (%s): spread=%d, score=%.2f\n", r.bookID, r.title, r.ratingSpread, r.polarizationScore)
	}
	if len(results) == 0 {
		fmt.Println("[]")
	}
}
```
