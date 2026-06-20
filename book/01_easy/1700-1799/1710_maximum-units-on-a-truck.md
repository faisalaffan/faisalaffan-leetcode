# 1710 — Maximum Units On A Truck

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumUnits(boxTypes [][]int, truckSize int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(log n) (for sorting)  
**Kompleksitas Ruang:** O(log n) (for sorting)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1710: Maximum Units on a Truck
// https://leetcode.com/problems/maximum-units-on-a-truck/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func MaximumUnits(boxTypes [][]int, truckSize int) int {
  // Custom sort dengan comparator
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	totalUnits := 0
	for _, box := range boxTypes {
		take := box[0]
		if truckSize < take {
			take = truckSize
		}
		totalUnits += take * box[1]
		truckSize -= take
		if truckSize == 0 {
			break
		}
	}
	return totalUnits
}

func main() {
	fmt.Println(MaximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(MaximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}
```
