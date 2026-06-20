# 1298 — Maximum Candies You Can Get From Boxes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxCandies(status []int, candies []int, keys [][]int,
	containedBoxes [][]int, initialBoxes []int) int
```

> **💡 Hint:** Iterative box opening.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1298: Maximum Candies You Can Get from Boxes
// https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/
// Difficulty: Hard
//
// Approach: Iterative box opening.
// Track which boxes we have, which keys we have, and which boxes are opened.
// Repeatedly scan all boxes: if we have the box AND we have the key AND it is
// not yet opened, open it, collect candies, add contained boxes and keys.
// Continue until no more progress.

import "fmt"

func maxCandies(status []int, candies []int, keys [][]int,
	containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	hasBox := make([]bool, n)
	hasKey := make([]bool, n)
	for _, b := range initialBoxes {
		hasBox[b] = true
	}
	for i := 0; i < n; i++ {
		if status[i] == 1 {
			hasKey[i] = true
		}
	}

	opened := make([]bool, n)
	total := 0
	for {
		progress := false
		for i := 0; i < n; i++ {
			if hasBox[i] && hasKey[i] && !opened[i] {
				opened[i] = true
				progress = true
				total += candies[i]
				for _, k := range keys[i] {
					hasKey[k] = true
				}
				for _, b := range containedBoxes[i] {
					hasBox[b] = true
				}
			}
		}
		if !progress {
			break
		}
	}
	return total
}

func main() {
	fmt.Println(maxCandies(
		[]int{1, 0, 1, 0},
		[]int{7, 5, 4, 100},
		[][]int{{}, {}, {1}, {}},
		[][]int{{1, 2}, {3}, {}, {}},
		[]int{0},
	)) // 16

	fmt.Println(maxCandies(
		[]int{1, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 1, 1},
		[][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
		[][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
		[]int{0},
	)) // 6
}
```
