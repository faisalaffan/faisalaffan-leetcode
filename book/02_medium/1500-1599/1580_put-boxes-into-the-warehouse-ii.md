# 1580 — Put Boxes Into The Warehouse Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxBoxesInWarehouseII(boxes []int, warehouse []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Prefix Sum

**Kompleksitas Waktu:** O(N log N + M), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1580: Put Boxes Into the Warehouse II
// https://leetcode.com/problems/put-boxes-into-the-warehouse-ii/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxBoxesInWarehouseII([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(MaxBoxesInWarehouseII([]int{1, 2, 2, 3, 4}, []int{3, 4, 1, 2}))
	fmt.Println(MaxBoxesInWarehouseII([]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}))
}

func MaxBoxesInWarehouseII(boxes []int, warehouse []int) int {
	// Time: O(N log N + M), Space: O(1)
	// In warehouse II, boxes can enter from either left or right side.
	// We can think of it as: each position's max height is the min of
	// the prefix max from left and prefix max from right.
  // Urutkan secara ascending — O(n log n)
	sort.Ints(boxes)

	n := len(warehouse)
	// Preprocess: effective height at each position
  // Alokasi slice integer
	leftMax := make([]int, n)
  // Alokasi slice integer
	rightMax := make([]int, n)

	leftMax[0] = warehouse[0]
	for i := 1; i < n; i++ {
		if warehouse[i] < leftMax[i-1] {
			leftMax[i] = warehouse[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	rightMax[n-1] = warehouse[n-1]
	for i := n - 2; i >= 0; i-- {
		if warehouse[i] < rightMax[i+1] {
			rightMax[i] = warehouse[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	// Effective height = max(leftMax, rightMax) since we can enter from either side
  // Alokasi slice integer
	effective := make([]int, n)
	for i := 0; i < n; i++ {
		if leftMax[i] > rightMax[i] {
			effective[i] = leftMax[i]
		} else {
			effective[i] = rightMax[i]
		}
	}

	// Greedily fit boxes
	boxIdx := 0
	for i := 0; i < n && boxIdx < len(boxes); i++ {
		if boxes[boxIdx] <= effective[i] {
			boxIdx++
		}
	}

	return boxIdx
}
```
