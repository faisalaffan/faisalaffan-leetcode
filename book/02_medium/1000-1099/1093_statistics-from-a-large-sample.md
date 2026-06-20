# 1093 — Statistics From A Large Sample

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sampleStats(count []int) []float64
```

> **💡 Hint:** Single pass to compute min, max, sum, mode.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = len(count) = 256  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1093: Statistics from a Large Sample
// https://leetcode.com/problems/statistics-from-a-large-sample/
// Difficulty: Medium
//
// Approach: Single pass to compute min, max, sum, mode.
//           Two-pointer for median.
// Time: O(n) where n = len(count) = 256
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(sampleStats([]int{0, 1, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	// Expected: [1.00000,3.00000,2.37500,2.50000,3.00000]
}

func sampleStats(count []int) []float64 {
	n := 0
	sum := 0
	minVal := -1
	maxVal := 0
	modeVal := 0
	modeCount := 0

	for i, c := range count {
		if c > 0 {
			n += c
			sum += i * c
			if minVal == -1 {
				minVal = i
			}
			maxVal = i
			if c > modeCount {
				modeCount = c
				modeVal = i
			}
		}
	}

	mean := float64(sum) / float64(n)

	// Median
	median := 0.0
	left := n / 2
	right := left + 1
	if n%2 == 1 {
		right = left
	}

	leftVal, rightVal := 0, 0
	acc := 0
	for i, c := range count {
		if c > 0 {
			if acc < left && acc+c >= left {
				leftVal = i
			}
			if acc < right && acc+c >= right {
				rightVal = i
			}
			acc += c
		}
	}
	median = float64(leftVal+rightVal) / 2.0

	return []float64{float64(minVal), float64(maxVal), mean, median, float64(modeVal)}
}
```
