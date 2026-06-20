# 2228 — Users With Two Purchases Within Seven Days

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findUsers(purchases [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2228: Users With Two Purchases Within Seven Days
// https://leetcode.com/problems/users-with-two-purchases-within-seven-days/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findUsers(purchases [][]int) []int {
	// Group purchases by user
  // Membuat map (HashMap) — pencarian O(1)
	userPurchases := make(map[int][]int)
	for _, p := range purchases {
		userID, date := p[0], p[1]
		userPurchases[userID] = append(userPurchases[userID], date)
	}

	result := []int{}
	for userID, dates := range userPurchases {
		if len(dates) < 2 {
			continue
		}
  // Urutkan secara ascending — O(n log n)
		sort.Ints(dates)
		for i := 1; i < len(dates); i++ {
			if dates[i]-dates[i-1] <= 7 {
				result = append(result, userID)
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
	fmt.Println(findUsers([][]int{{1, 1}, {2, 2}, {1, 7}, {1, 15}, {2, 3}}))
	// Expected: [1]

	// Test case 2
	fmt.Println(findUsers([][]int{{1, 1}, {2, 1}, {3, 1}}))
	// Expected: []
}
```
