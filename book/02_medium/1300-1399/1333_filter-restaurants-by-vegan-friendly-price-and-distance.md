# 1333 — Filter Restaurants By Vegan Friendly Price And Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n) for storing filtered results

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1333: Filter Restaurants by Vegan-Friendly, Price and Distance
// https://leetcode.com/problems/filter-restaurants-by-vegan-friendly-price-and-distance/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}, 1, 50, 10))
	// [3,1,5]

	// Test case 2
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}, 0, 50, 10))
	// [4,3,2,1,5]

	// Test case 3
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
	}, 1, 30, 10))
	// []
}

type restaurant struct {
	id, rating, veganFriendly, price, distance int
}

// Time: O(n log n) for sorting
// Space: O(n) for storing filtered results
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var filtered []restaurant
	for _, r := range restaurants {
		if veganFriendly == 1 && r[2] != 1 {
			continue
		}
		if r[3] > maxPrice {
			continue
		}
		if r[4] > maxDistance {
			continue
		}
		filtered = append(filtered, restaurant{r[0], r[1], r[2], r[3], r[4]})
	}

  // Custom sort dengan comparator
	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].rating != filtered[j].rating {
			return filtered[i].rating > filtered[j].rating
		}
		return filtered[i].id > filtered[j].id
	})

  // Alokasi slice integer
	result := make([]int, len(filtered))
	for i, r := range filtered {
		result[i] = r.id
	}
	return result
}
```
