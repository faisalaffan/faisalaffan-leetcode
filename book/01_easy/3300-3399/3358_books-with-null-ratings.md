# 3358 — Books With Null Ratings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ptr(i int) *int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i].BookID < result[j].BookID
	})
	return result
}
```
