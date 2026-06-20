# 1094 — Car Pooling

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func carPooling(trips [][]int, capacity int) bool
```

> **💡 Hint:** Difference array (prefix sum)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n + maxLocation)  
**Kompleksitas Ruang:** O(maxLocation)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1094: Car Pooling
// https://leetcode.com/problems/car-pooling/
// Difficulty: Medium
//
// Approach: Difference array (prefix sum)
// Time: O(n + maxLocation)
// Space: O(maxLocation)

import "fmt"

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4)) // false
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5)) // true
}

func carPooling(trips [][]int, capacity int) bool {
	maxLoc := 0
	for _, t := range trips {
		if t[2] > maxLoc {
			maxLoc = t[2]
		}
	}

  // Alokasi slice integer
	diff := make([]int, maxLoc+2)
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
	}

	current := 0
	for _, d := range diff {
		current += d
		if current > capacity {
			return false
		}
	}

	return true
}
```
