# 0881 — Boats To Save People

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BoatsToSavePeople(people []int, limit int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #881: Boats to Save People
// https://leetcode.com/problems/boats-to-save-people/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BoatsToSavePeople([]int{1, 2}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 2, 2, 1}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 5, 3, 4}, 5))
}

// Time: O(n log n) | Space: O(log n)
func BoatsToSavePeople(people []int, limit int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(people)
	left, right := 0, len(people)-1
	ans := 0

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		ans++
	}

	return ans
}
```
