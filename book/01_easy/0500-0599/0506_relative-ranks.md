# 0506 — Relative Ranks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RelativeRanks(score []int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #506: Relative Ranks
// https://leetcode.com/problems/relative-ranks/
// Difficulty: Easy

import (
	"fmt"
	"sort"
	"strconv"
)

// Time: O(n log n), Space: O(n)
func RelativeRanks(score []int) []string {
  // Alokasi slice integer
	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
  // Membuat map (HashMap) — pencarian O(1)
	rank := make(map[int]string)
	for i, s := range sorted {
		switch i {
		case 0:
			rank[s] = "Gold Medal"
		case 1:
			rank[s] = "Silver Medal"
		case 2:
			rank[s] = "Bronze Medal"
		default:
			rank[s] = strconv.Itoa(i + 1)
		}
	}
	result := make([]string, len(score))
	for i, s := range score {
		result[i] = rank[s]
	}
	return result
}

func main() {
	fmt.Println(RelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(RelativeRanks([]int{10, 3, 8, 9, 4}))
}
```
