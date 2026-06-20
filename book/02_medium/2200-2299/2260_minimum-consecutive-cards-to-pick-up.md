# 2260 — Minimum Consecutive Cards To Pick Up

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCardPickup(cards []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2260: Minimum Consecutive Cards to Pick Up
// https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumCardPickup(cards []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	last := make(map[int]int)
	minLen := len(cards) + 1

	for i, c := range cards {
		if prev, ok := last[c]; ok {
			dist := i - prev + 1
			if dist < minLen {
				minLen = dist
			}
		}
		last[c] = i
	}

	if minLen > len(cards) {
		return -1
	}
	return minLen
}

func main() {
	// Test case 1
	fmt.Println(minimumCardPickup([]int{3, 4, 2, 3, 4, 7}))
	// Expected: 4

	// Test case 2
	fmt.Println(minimumCardPickup([]int{1, 0, 5, 3}))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumCardPickup([]int{1, 2, 3, 4, 5, 1}))
	// Expected: 6
}
```
