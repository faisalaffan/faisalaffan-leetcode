# 0624 — Maximum Distance In Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxDistance(arrays [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = number of arrays  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #624: Maximum Distance in Arrays
// https://leetcode.com/problems/maximum-distance-in-arrays/
// Difficulty: Medium
// Time: O(n) where n = number of arrays
// Space: O(1)

import (
	"fmt"
)

func main() {
	arrays := [][]int{
		{1, 2, 3},
		{4, 5},
		{1, 2, 3},
	}
	fmt.Println(MaxDistance(arrays))
}

func MaxDistance(arrays [][]int) int {
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]
	maxDist := 0

	for i := 1; i < len(arrays); i++ {
		arr := arrays[i]
		dist1 := abs(arr[len(arr)-1] - minVal)
		dist2 := abs(maxVal - arr[0])
		if dist1 > maxDist {
			maxDist = dist1
		}
		if dist2 > maxDist {
			maxDist = dist2
		}
		if arr[0] < minVal {
			minVal = arr[0]
		}
		if arr[len(arr)-1] > maxVal {
			maxVal = arr[len(arr)-1]
		}
	}

	return maxDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
