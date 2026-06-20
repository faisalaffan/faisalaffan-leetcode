# 2854 — Rolling Average Steps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RollingAverageSteps(steps []int, k int) []float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2854: Rolling Average Steps
// https://leetcode.com/problems/rolling-average-steps/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(k)

import "fmt"

func RollingAverageSteps(steps []int, k int) []float64 {
	n := len(steps)
	if n < k {
		return []float64{}
	}

	result := make([]float64, n-k+1)
	var sum int
	for i := 0; i < k; i++ {
		sum += steps[i]
	}
	result[0] = float64(sum) / float64(k)

	for i := k; i < n; i++ {
		sum += steps[i] - steps[i-k]
		result[i-k+1] = float64(sum) / float64(k)
	}

	return result
}

func main() {
	fmt.Println(RollingAverageSteps([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(RollingAverageSteps([]int{10, 20}, 2))
}
```
