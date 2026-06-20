# 1445 — Apples Oranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func computeDifference(sales []Sale) map[string]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1445: Apples & Oranges
// https://leetcode.com/problems/apples-oranges/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Given sales of apples and oranges by date, compute the difference
// between apples sold and oranges sold for each date.

import "fmt"

// Sale represents a sale record.
type Sale struct {
	Date    string
	Fruit   string // "apples" or "oranges"
	SoldNum int
}

// computeDifference returns the difference (apples - oranges) for each date.
func computeDifference(sales []Sale) map[string]int {
  // Membuat map (HashMap) — pencarian O(1)
	diff := make(map[string]int)

	for _, s := range sales {
		if s.Fruit == "apples" {
			diff[s.Date] += s.SoldNum
		} else if s.Fruit == "oranges" {
			diff[s.Date] -= s.SoldNum
		}
	}

	return diff
}

func main() {
	sales := []Sale{
		{"2023-01-01", "apples", 10},
		{"2023-01-01", "oranges", 8},
		{"2023-01-02", "apples", 15},
		{"2023-01-02", "oranges", 20},
		{"2023-01-03", "apples", 5},
		{"2023-01-04", "oranges", 3},
	}

	diffs := computeDifference(sales)
	fmt.Println("Difference (apples - oranges) by date:")
	for date, diff := range diffs {
		fmt.Printf("  %s: %d\n", date, diff)
	}

	// Test 2: More sales, check zero diff
	sales2 := []Sale{
		{"2024-06-01", "apples", 100},
		{"2024-06-01", "oranges", 100},
		{"2024-06-02", "apples", 50},
	}
	diffs2 := computeDifference(sales2)
	fmt.Println("\nDifference 2:")
	for date, diff := range diffs2 {
		fmt.Printf("  %s: %d\n", date, diff)
	}

	// Test 3: Empty input
	diffs3 := computeDifference(nil)
	fmt.Printf("\nEmpty input diff count: %d\n", len(diffs3))

	// Test 4: Single fruit
	sales4 := []Sale{
		{"2024-07-01", "apples", 30},
	}
	diffs4 := computeDifference(sales4)
	fmt.Println("\nDifference 4 (single fruit):")
	for date, diff := range diffs4 {
		fmt.Printf("  %s: %d\n", date, diff)
	}
}
```
