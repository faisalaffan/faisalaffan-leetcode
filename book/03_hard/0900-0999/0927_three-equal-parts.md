# 0927 — Three Equal Parts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func threeEqualParts(arr []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #927: Three Equal Parts
// https://leetcode.com/problems/three-equal-parts/
// Difficulty: Hard
// Partition binary array into 3 parts with equal binary value.
// Leading zeros allowed. Parts must be non-empty.

import "fmt"

func threeEqualParts(arr []int) []int {
	totalOnes := 0
	for _, v := range arr {
		if v == 1 {
			totalOnes++
		}
	}

	if totalOnes%3 != 0 {
		return []int{-1, -1}
	}

	n := len(arr)
	if totalOnes == 0 {
		return []int{0, n - 1}
	}

	onesPerPart := totalOnes / 3

	// Find positions of the first, second, and third part's first 1
	first1 := -1
	second1 := -1
	third1 := -1

	count := 0
	for i, v := range arr {
		if v == 1 {
			count++
			if count == 1 {
				first1 = i
			}
			if count == onesPerPart+1 {
				second1 = i
			}
			if count == 2*onesPerPart+1 {
				third1 = i
			}
		}
	}

	// Now compare each part: arr[first1..second1-1], arr[second1..third1-1], arr[third1..]
	a, b, c := first1, second1, third1
	for c < n {
		if arr[a] != arr[b] || arr[b] != arr[c] {
			return []int{-1, -1}
		}
		a++
		b++
		c++
	}

	return []int{a - 1, b - 1}
}

func main() {
	fmt.Println(threeEqualParts([]int{1,0,1,0,1})) // Expected: [0,3]
	fmt.Println(threeEqualParts([]int{1,1,0,0,1})) // Expected: [0,2]
	fmt.Println(threeEqualParts([]int{1,1,0,1,1})) // Expected: [-1,-1]
	fmt.Println(threeEqualParts([]int{0,0,0,0,0})) // Expected: [0,4]
	fmt.Println(threeEqualParts([]int{1,0,1,0,1,0})) // Expected: [1,4]
}
```
