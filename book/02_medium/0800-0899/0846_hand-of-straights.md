# 0846 — Hand Of Straights

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func HandOfStraights(hand []int, groupSize int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #846: Hand of Straights
// https://leetcode.com/problems/hand-of-straights/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(HandOfStraights([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(HandOfStraights([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(HandOfStraights([]int{2, 1}, 2))
}

// Time: O(n log n) | Space: O(n)
func HandOfStraights(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	for _, card := range hand {
		count[card]++
	}

  // Alokasi slice integer
	unique := make([]int, 0, len(count))
	for card := range count {
		unique = append(unique, card)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(unique)

	for _, card := range unique {
		if count[card] > 0 {
			freq := count[card]
			for i := 0; i < groupSize; i++ {
				if count[card+i] < freq {
					return false
				}
				count[card+i] -= freq
			}
		}
	}

	return true
}
```
