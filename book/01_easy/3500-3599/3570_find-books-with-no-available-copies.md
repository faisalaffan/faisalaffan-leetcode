# 3570 — Find Books With No Available Copies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindBooksWithNoAvailableCopies(books []LibBook) []LibBook
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i].BookID < result[j].BookID
	})
	return result
}
```
