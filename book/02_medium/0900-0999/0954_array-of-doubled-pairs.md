# 0954 — Array Of Doubled Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canReorderDoubled(arr []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #954: Array of Doubled Pairs
// https://leetcode.com/problems/array-of-doubled-pairs/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(n)
func canReorderDoubled(arr []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

  // Alokasi slice integer
	keys := make([]int, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
  // Custom sort dengan comparator
	sort.Slice(keys, func(i, j int) bool {
		return abs(keys[i]) < abs(keys[j])
	})

	for _, v := range keys {
		if freq[v] > freq[2*v] {
			return false
		}
		freq[2*v] -= freq[v]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{1, 2, 4, 16, 8, 4}))
}
```
