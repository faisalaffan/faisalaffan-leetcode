# 1282 — Group The People Given The Group Size They Belong To

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func groupThePeople(groupSizes []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1282: Group the People Given the Group Size They Belong To
// https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
// Difficulty: Medium

// Group people by desired group size, then partition each group.

// Time: O(n)
// Space: O(n)

func groupThePeople(groupSizes []int) [][]int {
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[int][]int)
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)

	for person, size := range groupSizes {
		groups[size] = append(groups[size], person)
		if len(groups[size]) == size {
			result = append(result, groups[size])
			delete(groups, size)
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Printf("%v\n", groupThePeople([]int{2, 1, 3, 3, 3, 2}))
}
```
