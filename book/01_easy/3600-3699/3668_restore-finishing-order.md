# 3668 — Restore Finishing Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RestoreFinishingOrder(order []int, friends []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3668: Restore Finishing Order
// https://leetcode.com/problems/restore-finishing-order/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RestoreFinishingOrder([]int{3, 1, 2, 4, 5}, []int{2, 4}))
	fmt.Println(RestoreFinishingOrder([]int{1, 2, 3, 4, 5}, []int{1, 3, 5}))
}

// Time: O(n)
// Space: O(n)
func RestoreFinishingOrder(order []int, friends []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	friendSet := make(map[int]bool)
	for _, f := range friends {
		friendSet[f] = true
	}

  // Alokasi slice integer
	res := make([]int, 0, len(friends))
	for _, id := range order {
		if friendSet[id] {
			res = append(res, id)
		}
	}
	return res
}

func init() {
	_ = sort.Ints
}
```
