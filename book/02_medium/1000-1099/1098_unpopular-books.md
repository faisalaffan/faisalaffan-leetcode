# 1098 — Unpopular Books

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func unpopularBooks(books []string, orders []struct {
	name     string
	quantity int
	daysAgo  int
}) []string
```

> **💡 Hint:** Filter books ordered less than 10 times in the last year

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m) where n = books, m = orders  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1098: Unpopular Books
// https://leetcode.com/problems/unpopular-books/
// Difficulty: Medium
//
// Approach: Filter books ordered less than 10 times in the last year
// Time: O(n + m) where n = books, m = orders
// Space: O(n)

import "fmt"

func main() {
	// Books: (book_id, name)
	books := []string{"Book A", "Book B", "Book C", "Book D"}
	// Orders: (book_name, quantity, days_ago)
	orders := []struct {
		name     string
		quantity int
		daysAgo  int
	}{
		{"Book A", 5, 30},
		{"Book B", 15, 10},
		{"Book C", 8, 20},
	}
	fmt.Println(unpopularBooks(books, orders))
}

func unpopularBooks(books []string, orders []struct {
	name     string
	quantity int
	daysAgo  int
}) []string {
  // Membuat map (HashMap) — pencarian O(1)
	orderCount := make(map[string]int)
	for _, o := range orders {
		if o.daysAgo <= 365 {
			orderCount[o.name] += o.quantity
		}
	}

	result := make([]string, 0)
	for _, b := range books {
		count, exists := orderCount[b]
		if !exists || count < 10 {
			result = append(result, b)
		}
	}

	return result
}
```
