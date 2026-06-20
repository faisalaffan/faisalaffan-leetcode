# 1264 — Page Recommendations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func pageRecommendations(userID int, likes []like, friendships []friend) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1264: Page Recommendations
// https://leetcode.com/problems/page-recommendations/
// Difficulty: Medium [Paid]

// Recommend pages liked by friends of user1 but not liked by user1.

// Time: O(n log n)
// Space: O(n)

type like struct {
	userID int
	pageID int
}

type friend struct {
	user1 int
	user2 int
}

func pageRecommendations(userID int, likes []like, friendships []friend) []int {
  // Membuat map (HashMap) — pencarian O(1)
	likedByUser := make(map[int]bool)
	for _, l := range likes {
		if l.userID == userID {
			likedByUser[l.pageID] = true
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	friends := make(map[int]bool)
	for _, f := range friendships {
		if f.user1 == userID {
			friends[f.user2] = true
		}
		if f.user2 == userID {
			friends[f.user1] = true
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	recommend := make(map[int]bool)
	for _, l := range likes {
		if friends[l.userID] && !likedByUser[l.pageID] {
			recommend[l.pageID] = true
		}
	}

  // Alokasi slice integer
	result := make([]int, 0, len(recommend))
	for p := range recommend {
		result = append(result, p)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	likes := []like{
		{1, 101}, {1, 102}, {2, 101}, {2, 103}, {3, 102},
	}
	friendships := []friend{
		{1, 2}, {1, 3},
	}
	fmt.Printf("%v (expected: [103] or [102 103])\n",
		pageRecommendations(1, likes, friendships))

	fmt.Printf("%v (expected: [])\n",
		pageRecommendations(2, likes, friendships))
}
```
