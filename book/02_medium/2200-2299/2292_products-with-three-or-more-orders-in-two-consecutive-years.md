# 2292 — Products With Three Or More Orders In Two Consecutive Years

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findProducts(orders [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2292: Products With Three or More Orders in Two Consecutive Years
// https://leetcode.com/problems/products-with-three-or-more-orders-in-two-consecutive-years/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Order struct {
	ProductID int
	Year      int
}

func findProducts(orders [][]int) []int {
	// Group orders by product and year
  // Membuat map (HashMap) — pencarian O(1)
	productYears := make(map[int]map[int]int)
	for _, o := range orders {
		productID, year := o[0], o[1]
		if productYears[productID] == nil {
			productYears[productID] = make(map[int]int)
		}
		productYears[productID][year]++
	}

	result := []int{}
	for pid, years := range productYears {
		yearList := []int{}
		for y := range years {
			yearList = append(yearList, y)
		}
  // Urutkan secara ascending — O(n log n)
		sort.Ints(yearList)

		for i := 1; i < len(yearList); i++ {
			if yearList[i]-yearList[i-1] == 1 &&
				years[yearList[i]] >= 3 &&
				years[yearList[i-1]] >= 3 {
				result = append(result, pid)
				break
			}
		}
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	// Test case 1
	fmt.Println(findProducts([][]int{{1, 2020}, {1, 2020}, {1, 2020}, {1, 2021}, {1, 2021}, {1, 2021}, {2, 2020}}))
	// Expected: [1]

	// Test case 2
	fmt.Println(findProducts([][]int{{1, 2020}, {2, 2020}}))
	// Expected: []
}
```
