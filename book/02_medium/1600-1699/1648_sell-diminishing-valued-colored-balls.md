# 1648 — Sell Diminishing Valued Colored Balls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxProfit(inventory []int, orders int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1648: Sell Diminishing-Valued Colored Balls
// https://leetcode.com/problems/sell-diminishing-valued-colored-balls/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxProfit([]int{2, 5}, 4))
	fmt.Println(MaxProfit([]int{3, 5}, 6))
	fmt.Println(MaxProfit([]int{2, 8, 4, 10, 6}, 20))
}

func MaxProfit(inventory []int, orders int) int {
	// Time: O(N log N), Space: O(N)
	const mod = 1_000_000_007

	// Sort descending
  // Custom sort dengan comparator
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] > inventory[j]
	})

	// Append 0 for convenience
	inventory = append(inventory, 0)
	n := len(inventory)

	profit := 0
	count := 0

	for i := 0; i < n-1; i++ {
		if inventory[i] == inventory[i+1] {
			continue
		}

		// Height difference between current and next level
		height := inventory[i] - inventory[i+1]
		width := i + 1 // number of types with this count
		total := height * width

		if count+total <= orders {
			// Take all balls at this level
			// Sum for this level: width * sum of (inventory[i], inventory[i]-1, ..., inventory[i+1]+1)
			top := inventory[i]
			bottom := inventory[i+1] + 1
			sumLevel := (top + bottom) * height / 2
			profit = (profit + sumLevel*width) % mod
			count += total
		} else {
			// Take only some
			remaining := orders - count
			fullRows := remaining / width
			extra := remaining % width

			top := inventory[i]
			bottom := top - fullRows + 1
			sumFull := (top + bottom) * fullRows / 2
			profit = (profit + sumFull*width) % mod
			profit = (profit + bottom*extra) % mod

			count = orders
			break
		}
	}

	return profit
}
```
