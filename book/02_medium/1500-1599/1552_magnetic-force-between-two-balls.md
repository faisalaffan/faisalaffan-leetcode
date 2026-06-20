# 1552 — Magnetic Force Between Two Balls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxDistance(position []int, m int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N log N + N log MAX), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1552: Magnetic Force Between Two Balls
// https://leetcode.com/problems/magnetic-force-between-two-balls/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxDistance([]int{1, 2, 3, 4, 7}, 3))
	fmt.Println(MaxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2))
	fmt.Println(MaxDistance([]int{1, 2, 3, 4, 5, 6}, 3))
}

func MaxDistance(position []int, m int) int {
	// Time: O(N log N + N log MAX), Space: O(1)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(position)

	low, high := 1, position[len(position)-1]-position[0]
	result := 0

	for low <= high {
		mid := low + (high-low)/2
		if canPlace(position, m, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func canPlace(position []int, m int, minDist int) bool {
	count := 1
	lastPos := position[0]

	for i := 1; i < len(position); i++ {
		if position[i]-lastPos >= minDist {
			count++
			lastPos = position[i]
			if count >= m {
				return true
			}
		}
	}

	return false
}
```
