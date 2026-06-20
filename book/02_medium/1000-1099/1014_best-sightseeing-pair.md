# 1014 — Best Sightseeing Pair

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScoreSightseeingPair(values []int) int
```

> **💡 Hint:** Track max value of (values[i] + i) seen so far

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1014: Best Sightseeing Pair
// https://leetcode.com/problems/best-sightseeing-pair/
// Difficulty: Medium
//
// Approach: Track max value of (values[i] + i) seen so far
// Score = values[i] + values[j] + i - j = (values[i] + i) + (values[j] - j)
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6})) // 11
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))           // 2
}

func maxScoreSightseeingPair(values []int) int {
	maxI := values[0]
	result := 0

	for j := 1; j < len(values); j++ {
		if maxI+values[j]-j > result {
			result = maxI + values[j] - j
		}
		if values[j]+j > maxI {
			maxI = values[j] + j
		}
	}

	return result
}
```
