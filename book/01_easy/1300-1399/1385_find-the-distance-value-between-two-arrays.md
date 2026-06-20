# 1385 — Find The Distance Value Between Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int

import (
	"fmt"
	"sort"
)

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log m + m log m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1385: Find the Distance Value Between Two Arrays
// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/
// Difficulty: Easy
//
// LeetCode submission: func findTheDistanceValue(arr1 []int, arr2 []int, d int) int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTheDistanceValueBetweenTwoArrays([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)) // 2
	fmt.Println(FindTheDistanceValueBetweenTwoArrays([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3)) // 2
}

// Time: O(n log m + m log m), Space: O(1)
func FindTheDistanceValueBetweenTwoArrays(arr1 []int, arr2 []int, d int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr2)
	count := 0
	for _, v := range arr1 {
		if isFar(v, arr2, d) {
			count++
		}
	}
	return count
}

func isFar(v int, arr []int, d int) bool {
	idx := sort.SearchInts(arr, v)
	if idx < len(arr) && abs(arr[idx]-v) <= d {
		return false
	}
	if idx > 0 && abs(arr[idx-1]-v) <= d {
		return false
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
