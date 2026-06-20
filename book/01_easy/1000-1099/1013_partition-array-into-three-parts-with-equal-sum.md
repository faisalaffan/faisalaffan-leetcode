# 1013 — Partition Array Into Three Parts With Equal Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canThreePartsEqualSum(arr []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1013: Partition Array Into Three Parts With Equal Sum
// https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1})) // true
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1})) // false
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))    // true
}

// LeetCode submission: canThreePartsEqualSum
func canThreePartsEqualSum(arr []int) bool {
	total := 0
	for _, v := range arr {
		total += v
	}
	if total%3 != 0 {
		return false
	}
	target := total / 3
	sum, count := 0, 0
	for _, v := range arr {
		sum += v
		if sum == target {
			count++
			sum = 0
		}
	}
	return count >= 3
}

func partitionArrayIntoThreePartsWithEqualSum(arr []int) bool {
	return canThreePartsEqualSum(arr)
}
```
